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
                    is ClerkResult.Success -> dispatchSuccess(callback)
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
                            dispatchSuccess(callback)
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
                    is ClerkResult.Success -> dispatchSuccess(callback)
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
        try {
            when (val result = Clerk.auth.getToken()) {
                is ClerkResult.Success -> result.value
                is ClerkResult.Failure -> null
            }
        } catch (_: Exception) {
            null
        }
    }

    private suspend fun dispatchSuccess(callback: AuthCallback) {
        val token = getTokenBlocking()
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
