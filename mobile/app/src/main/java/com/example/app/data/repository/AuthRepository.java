package com.example.app.data.repository;

import android.content.Context;

import com.example.app.BuildConfig;
import com.example.app.data.local.TokenManager;
import com.example.app.data.remote.RetrofitClient;
import com.example.app.data.remote.api.AuthApi;
import com.example.app.data.remote.model.request.auth.RegisterRequest;
import com.example.app.data.remote.model.request.auth.SyncRequest;
import com.example.app.data.remote.model.request.auth.TokenRequest;
import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.auth.FirebaseSignUpResponse;
import com.example.app.data.remote.model.response.auth.SyncResponse;
import com.example.app.data.remote.model.response.auth.TokenResponse;

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


    public interface AuthCallback<T> {
        void onSuccess(T data);
        void onError(String message);
    }

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

                    syncUser(tokenData.getIdToken(),"", callback);

                } else {
                    try {
                        String errorDetail = response.errorBody() != null ? response.errorBody().string() : "Lỗi không xác định";
                        callback.onError("Lỗi từ đăng nhập: " + errorDetail);
                    } catch (Exception e) {
                        callback.onError("Tài khoản mật khẩu không đúng");
                    }
                }
            }

            @Override
            public void onFailure(Call<ApiResponse<TokenResponse>> call, Throwable t) {
                callback.onError("Lỗi kết nối: " + t.getMessage());
            }
        });
    }


    public void syncUser(String idToken,String name, AuthCallback<SyncResponse> callback) {
        SyncRequest request = new SyncRequest(idToken,name);

        authApi.synUser(request).enqueue(new Callback<ApiResponse<SyncResponse>>() {
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

    public void register(String email, String password, String name, AuthCallback<SyncResponse> callback) {
        RegisterRequest request = new RegisterRequest(email, password, true);

        authApi.getRegister(BuildConfig.FIREBASE_API_KEY, request).enqueue(new Callback<FirebaseSignUpResponse>() {
            @Override
            public void onResponse(Call<FirebaseSignUpResponse> call, Response<FirebaseSignUpResponse> response) {
                if (response.isSuccessful() && response.body() != null) {
                    FirebaseSignUpResponse user = response.body();
                    tokenManager.saveToken(user.getIdToken(), user.getRefreshToken());
                    String idToken = user.getIdToken();
                    syncUser(idToken, name, callback);

                } else {
                    try {
                        String errorDetail = response.errorBody() != null ? response.errorBody().string() : "Lỗi không xác định";
                        callback.onError("Lỗi từ Firebase: " + errorDetail);
                    } catch (Exception e) {
                        callback.onError("Không thể đọc lỗi từ Firebase");
                    }
                }
            }

            @Override
            public void onFailure(Call<FirebaseSignUpResponse> call, Throwable t) {
                callback.onError("Lỗi kết nối: " + t.getMessage());
            }
        });
    }


    public void logout() {
        tokenManager.clear();
    }

    public boolean isLoggedIn() {
        return tokenManager.hasToken();
    }
}