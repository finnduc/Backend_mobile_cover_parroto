package com.example.app.data.remote

import com.clerk.api.Clerk
import com.clerk.api.auth.types.VerificationType
import com.clerk.api.network.serialization.ClerkResult
import com.clerk.api.network.serialization.errorMessage
import com.clerk.api.signup.SignUp
import com.clerk.api.signup.sendCode
import com.clerk.api.signup.verifyCode
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.SupervisorJob
import kotlinx.coroutines.delay
import kotlinx.coroutines.launch
import kotlinx.coroutines.runBlocking
import kotlinx.coroutines.withContext

object ClerkAuthBridge {
    interface AuthCallback {
        fun onSuccess(token: String?)
        fun onNeedsVerification()
        fun onError(message: String)
    }

    interface SimpleCallback {
        fun onSuccess()
        fun onError(message: String)
    }

    private val scope = CoroutineScope(SupervisorJob() + Dispatchers.IO)

    @JvmStatic
    fun signIn(email: String, password: String, callback: AuthCallback) {
        scope.launch {
            try {
                when (val result = Clerk.auth.signInWithPassword {
                    identifier = email
                    this.password = password
                }) {
                    is ClerkResult.Success -> {
                        val signIn = result.value
                        android.util.Log.d("ClerkAuth", "signIn success, status=${signIn.status}, sessionId=${signIn.createdSessionId}")
                        if (signIn.status.name.equals("COMPLETE", ignoreCase = true)) {
                            activateSessionAndDispatch(signIn.createdSessionId, callback)
                        } else {
                            val statusName = signIn.status.name
                            val userMsg = when (statusName) {
                                "NEEDS_SECOND_FACTOR" -> "Tài khoản của bạn yêu cầu xác thực 2 yếu tố (2FA/MFA). Hãy tắt 2FA trong Clerk Dashboard hoặc cấu hình phù hợp."
                                "NEEDS_NEW_PASSWORD" -> "Bạn cần đặt lại mật khẩu mới để tiếp tục."
                                else -> "Đăng nhập chưa hoàn tất (Trạng thái: $statusName)."
                            }
                            dispatchError(callback, userMsg)
                        }
                    }
                    is ClerkResult.Failure -> dispatchError(callback, result.errorMessage)
                }
            } catch (e: Exception) {
                dispatchError(callback, e.message ?: "Sign in failed")
            }
        }
    }

    @JvmStatic
    fun signUp(email: String, password: String, fullName: String, callback: AuthCallback) {
        scope.launch {
            try {
                val names = splitName(fullName)
                when (val result = Clerk.auth.signUp {
                    this.email = email
                    this.password = password
                    firstName = names.first
                    lastName = names.second
                }) {
                    is ClerkResult.Success -> {
                        val signUp = result.value
                        if (isComplete(signUp)) {
                            activateSessionAndDispatch(signUp.createdSessionId, callback)
                        } else {
                            when (val sendResult = signUp.sendCode { this.email = signUp.emailAddress }) {
                                is ClerkResult.Success -> withContext(Dispatchers.Main) {
                                    callback.onNeedsVerification()
                                }
                                is ClerkResult.Failure -> dispatchError(callback, sendResult.errorMessage)
                            }
                        }
                    }
                    is ClerkResult.Failure -> dispatchError(callback, result.errorMessage)
                }
            } catch (e: Exception) {
                dispatchError(callback, e.message ?: "Sign up failed")
            }
        }
    }

    @JvmStatic
    fun verifySignUp(code: String, callback: AuthCallback) {
        scope.launch {
            try {
                val signUp = Clerk.auth.currentSignUp
                if (signUp == null) {
                    dispatchError(callback, "No sign up is waiting for verification")
                    return@launch
                }
                when (val result = signUp.verifyCode(code, VerificationType.EMAIL)) {
                    is ClerkResult.Success -> {
                        // After verification, the signUp object should have a createdSessionId
                        val updatedSignUp = Clerk.auth.currentSignUp ?: result.value
                        if (isComplete(updatedSignUp)) {
                            activateSessionAndDispatch(updatedSignUp.createdSessionId, callback)
                        } else {
                            dispatchError(callback, "Đăng ký chưa hoàn tất (Trạng thái: ${updatedSignUp.status.name})")
                        }
                    }
                    is ClerkResult.Failure -> dispatchError(callback, result.errorMessage)
                }
            } catch (e: Exception) {
                dispatchError(callback, e.message ?: "Verification failed")
            }
        }
    }

    @JvmStatic
    fun signOut(callback: SimpleCallback) {
        scope.launch {
            try {
                when (val result = Clerk.auth.signOut()) {
                    is ClerkResult.Success -> withContext(Dispatchers.Main) { callback.onSuccess() }
                    is ClerkResult.Failure -> withContext(Dispatchers.Main) {
                        callback.onError(result.errorMessage)
                    }
                }
            } catch (e: Exception) {
                withContext(Dispatchers.Main) {
                    callback.onError(e.message ?: "Sign out failed")
                }
            }
        }
    }

    @JvmStatic
    fun getTokenBlocking(): String? = runBlocking(Dispatchers.IO) {
        getTokenWithRetry("getToken")
    }

    private suspend fun getTokenWithRetry(tag: String, maxAttempts: Int = 3): String? {
        var lastError: String? = null
        for (attempt in 1..maxAttempts) {
            try {
                when (val result = Clerk.auth.getToken()) {
                    is ClerkResult.Success -> {
                        val token = result.value
                        if (token == null || token.isEmpty()) {
                            android.util.Log.e("ClerkAuth", "$tag returned empty token (attempt $attempt/$maxAttempts)")
                        } else {
                            return token
                        }
                    }
                    is ClerkResult.Failure -> {
                        lastError = result.errorMessage
                        android.util.Log.w("ClerkAuth", "$tag failed (attempt $attempt/$maxAttempts): $lastError")
                    }
                }
            } catch (e: Exception) {
                lastError = e.message
                android.util.Log.w("ClerkAuth", "$tag exception (attempt $attempt/$maxAttempts)", e)
            }
            if (attempt < maxAttempts) {
                delay(500L * attempt)
            }
        }
        android.util.Log.e("ClerkAuth", "$tag failed after $maxAttempts attempts: $lastError")
        return null
    }

    private suspend fun activateSessionAndDispatch(sessionId: String?, callback: AuthCallback) {
        if (!sessionId.isNullOrEmpty()) {
            try {
                when (val activeResult = Clerk.auth.setActive(sessionId)) {
                    is ClerkResult.Success -> {
                        android.util.Log.d("ClerkAuth", "setActive success for session: $sessionId")
                    }
                    is ClerkResult.Failure -> {
                        android.util.Log.e("ClerkAuth", "setActive failed: ${activeResult.errorMessage}")
                    }
                }
            } catch (e: Exception) {
                android.util.Log.e("ClerkAuth", "setActive exception", e)
            }
        } else {
            android.util.Log.w("ClerkAuth", "No sessionId available for setActive")
        }
        dispatchSuccess(callback)
    }

    private suspend fun dispatchSuccess(callback: AuthCallback) {
        val token = getTokenWithRetry("dispatchSuccess")
        withContext(Dispatchers.Main) {
            callback.onSuccess(token)
        }
    }

    private suspend fun dispatchError(callback: AuthCallback, message: String) {
        withContext(Dispatchers.Main) {
            callback.onError(message)
        }
    }

    private fun splitName(fullName: String): Pair<String, String> {
        val parts = fullName.trim().split(Regex("\\s+")).filter { it.isNotBlank() }
        if (parts.isEmpty()) return "" to ""
        if (parts.size == 1) return parts[0] to ""
        return parts[0] to parts.drop(1).joinToString(" ")
    }

    private fun isComplete(signUp: SignUp): Boolean {
        return signUp.status.name.equals("COMPLETE", ignoreCase = true)
    }
}
