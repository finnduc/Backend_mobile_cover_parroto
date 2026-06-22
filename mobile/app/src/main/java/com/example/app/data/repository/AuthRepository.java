package com.example.app.data.repository;

import android.content.Context;

import com.example.app.data.local.TokenManager;
import com.example.app.data.remote.ClerkAuthBridge;
import com.example.app.data.remote.RetrofitClient;
import com.example.app.data.remote.api.AuthApi;
import com.example.app.data.remote.api.UserApi;
import com.example.app.data.remote.model.request.auth.UpdateProfileRequest;
import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.user.UserResponse;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class AuthRepository {

    private final AuthApi authApi;
    private final UserApi userApi;
    private final TokenManager tokenManager;

    public AuthRepository(Context context) {
        this.authApi = RetrofitClient.getInstance(context).getAuthApi();
        this.userApi = RetrofitClient.getInstance(context).getUserApi();
        this.tokenManager = TokenManager.getInstance(context);
    }

    public interface authCallBack<T> {
        void onSuccess(T data);

        void onError(String message);

        default void onNeedsVerification() {
            onError("Vui lòng xác minh email để hoàn tất đăng ký");
        }
    }

    public void Register(String email, String password, String name, authCallBack<String> callback) {
        ClerkAuthBridge.signUp(email, password, name, new ClerkAuthBridge.AuthCallback() {
            @Override
            public void onSuccess(String token) {
                saveToken(token);
                completeSignupAndSync(new authCallBack<UserResponse>() {
                    @Override
                    public void onSuccess(UserResponse data) {
                        callback.onSuccess("Đăng ký thành công");
                    }

                    @Override
                    public void onError(String message) {
                        clearAuthState();
                        callback.onError(message);
                    }
                });
            }

            @Override
            public void onNeedsVerification() {
                callback.onNeedsVerification();
            }

            @Override
            public void onError(String message) {
                clearAuthState();
                callback.onError(message);
            }
        });
    }

    public void verifyRegistration(String code, authCallBack<String> callback) {
        ClerkAuthBridge.verifySignUp(code, new ClerkAuthBridge.AuthCallback() {
            @Override
            public void onSuccess(String token) {
                saveToken(token);
                completeSignupAndSync(new authCallBack<UserResponse>() {
                    @Override
                    public void onSuccess(UserResponse data) {
                        callback.onSuccess("Đăng ký thành công");
                    }

                    @Override
                    public void onError(String message) {
                        clearAuthState();
                        callback.onError(message);
                    }
                });
            }

            @Override
            public void onNeedsVerification() {
                callback.onNeedsVerification();
            }

            @Override
            public void onError(String message) {
                callback.onError(message);
            }
        });
    }

    public void login(String email, String password, authCallBack<UserResponse> callback) {
        ClerkAuthBridge.signIn(email, password, new ClerkAuthBridge.AuthCallback() {
            @Override
            public void onSuccess(String token) {
                saveToken(token);
                syncAuthenticatedUser(new authCallBack<UserResponse>() {
                    @Override
                    public void onSuccess(UserResponse data) {
                        callback.onSuccess(data);
                    }

                    @Override
                    public void onError(String message) {
                        clearAuthState();
                        callback.onError(message);
                    }
                });
            }

            @Override
            public void onNeedsVerification() {
                callback.onNeedsVerification();
            }

            @Override
            public void onError(String message) {
                callback.onError(message);
            }
        });
    }

    public void updateProfile(UpdateProfileRequest request, authCallBack<ApiResponse<UserResponse>> callback) {
        userApi.updateProfile(request).enqueue(new Callback<ApiResponse<UserResponse>>() {
            @Override
            public void onResponse(Call<ApiResponse<UserResponse>> call, Response<ApiResponse<UserResponse>> response) {
                if (response.isSuccessful() && response.body() != null && response.body().getData() != null) {
                    callback.onSuccess(response.body());
                } else {
                    callback.onError(getErrorMessage(response));
                }
            }

            @Override
            public void onFailure(Call<ApiResponse<UserResponse>> call, Throwable t) {
                callback.onError("Lỗi: " + t.getMessage());
            }
        });
    }

    private void completeSignupAndSync(authCallBack<UserResponse> callback) {
        authApi.completeSignup().enqueue(new Callback<ApiResponse<String>>() {
            @Override
            public void onResponse(Call<ApiResponse<String>> call, Response<ApiResponse<String>> response) {
                if (response.isSuccessful()) {
                    syncAuthenticatedUser(callback);
                } else {
                    callback.onError("Lỗi hoàn tất đăng ký server: " + getErrorMessage(response));
                }
            }

            @Override
            public void onFailure(Call<ApiResponse<String>> call, Throwable t) {
                callback.onError("Lỗi hoàn tất đăng ký server: " + t.getMessage());
            }
        });
    }

    private void syncAuthenticatedUser(authCallBack<UserResponse> callback) {
        authApi.auth().enqueue(new Callback<ApiResponse<UserResponse>>() {
            @Override
            public void onResponse(Call<ApiResponse<UserResponse>> call, Response<ApiResponse<UserResponse>> response) {
                if (response.isSuccessful() && response.body() != null && response.body().getData() != null) {
                    UserResponse userResponse = response.body().getData();
                    saveUserInfo(userResponse);
                    callback.onSuccess(userResponse);
                } else {
                    callback.onError("Lỗi đồng bộ server: " + getErrorMessage(response));
                }
            }

            @Override
            public void onFailure(Call<ApiResponse<UserResponse>> call, Throwable t) {
                callback.onError("Lỗi đồng bộ server: " + t.getMessage());
            }
        });
    }

    private void saveToken(String token) {
        if (token != null && !token.isEmpty()) {
            tokenManager.saveToken(token, "");
        }
    }

    private void saveUserInfo(UserResponse userResponse) {
        tokenManager.saveUserInfo(
                userResponse.getUid(),
                userResponse.getEmail(),
                userResponse.getName(),
                userResponse.getAvatarUrl(),
                userResponse.getUserRole(),
                userResponse.getPhone()
        );
    }

    private void clearAuthState() {
        tokenManager.clear();
    }

    private String getErrorMessage(Response<?> response) {
        try {
            if (response.errorBody() != null) {
                return response.errorBody().string();
            }
        } catch (Exception e) {
            return e.getMessage();
        }
        return "Code " + response.code();
    }
}
