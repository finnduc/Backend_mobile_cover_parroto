package com.example.app.data.repository;

import android.content.Context;

import com.example.app.data.local.TokenManager;
import com.example.app.data.remote.RetrofitClient;
import com.example.app.data.remote.api.AuthApi;
import com.example.app.data.remote.model.request.SyncRequest;
import com.example.app.data.remote.model.request.TokenRequest;
import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.SyncResponse;
import com.example.app.data.remote.model.response.TokenResponse;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class AuthRepository {

    private final AuthApi authApi;
    private final TokenManager tokenManager;

    public AuthRepository(Context context) {
        this.authApi = RetrofitClient.getInstance(context).getAuthApi();
        this.tokenManager = TokenManager.getInstance(context);
    }

    // ====== CALLBACK INTERFACE ======

    public interface AuthCallback<T> {
        void onSuccess(T data);
        void onError(String message);
    }

    // ====== LOGIN ======

    public void login(String email, String password, AuthCallback<SyncResponse> callback) {
        TokenRequest request = new TokenRequest(email, password);

        authApi.getToken(request).enqueue(new Callback<ApiResponse<TokenResponse>>() {
            @Override
            public void onResponse(Call<ApiResponse<TokenResponse>> call,
                                   Response<ApiResponse<TokenResponse>> response) {
                if (response.isSuccessful()
                        && response.body() != null
                        && response.body().issuccess()) {

                    TokenResponse tokenData = response.body().getData();
                    tokenManager.saveToken(tokenData.getIdToken(), tokenData.getRefreshToken());

                    // Bước 2: sync user
                    syncUser(tokenData.getIdToken(), callback);

                } else {
                    callback.onError("Email hoặc mật khẩu không đúng");
                }
            }

            @Override
            public void onFailure(Call<ApiResponse<TokenResponse>> call, Throwable t) {
                callback.onError("Lỗi kết nối: " + t.getMessage());
            }
        });
    }

    // ====== SYNC USER ======

    public void syncUser(String idToken, AuthCallback<SyncResponse> callback) {
        SyncRequest request = new SyncRequest(idToken);

        authApi.syncUser(request).enqueue(new Callback<ApiResponse<SyncResponse>>() {
            @Override
            public void onResponse(Call<ApiResponse<SyncResponse>> call,
                                   Response<ApiResponse<SyncResponse>> response) {
                if (response.isSuccessful()
                        && response.body() != null
                        && response.body().issuccess()) {

                    SyncResponse user = response.body().getData();
                    tokenManager.saveUserInfo(
                            user.getId(),
                            user.getEmail(),
                            user.getName(),
                            user.getAvatarURL()
                    );

                    callback.onSuccess(user);

                } else {
                    callback.onError("Không thể đồng bộ tài khoản");
                }
            }

            @Override
            public void onFailure(Call<ApiResponse<SyncResponse>> call, Throwable t) {
                callback.onError("Lỗi kết nối: " + t.getMessage());
            }
        });
    }

    // ====== REGISTER ======

    public void register(String email, String password, AuthCallback<SyncResponse> callback) {
        // TODO: Gọi Firebase REST API để tạo tài khoản
        // Sau đó gọi syncUser() giống login
    }

    // ====== LOGOUT ======

    public void logout() {
        tokenManager.clear();
    }

    // ====== CHECK LOGIN ======

    public boolean isLoggedIn() {
        return tokenManager.hasToken();
    }
}