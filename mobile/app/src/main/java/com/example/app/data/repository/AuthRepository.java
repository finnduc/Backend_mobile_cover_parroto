package com.example.app.data.repository;

import android.content.Context;

import com.example.app.data.local.TokenManager;
import com.example.app.data.remote.RetrofitClient;
import com.example.app.data.remote.api.AuthApi;
import com.example.app.data.remote.api.UserApi;
import com.example.app.data.remote.model.request.auth.LoginRequest;
import com.example.app.data.remote.model.request.auth.RegisterRequest;
import com.example.app.data.remote.model.request.auth.UpdateProfileRequest;
import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.auth.FirebaseSignUpResponse;
import com.example.app.data.remote.model.response.user.UserResponse;
import com.example.app.utils.Constants;

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

    public interface authCallBack<T>{
        void onSuccess(T data);
        void onError(String message);
    }

    public void Register(String email, String password,String name, authCallBack<String> callback) {
        RegisterRequest registerRequest = new RegisterRequest(email, password, true);
        authApi.register(registerRequest, Constants.FIREBASE_API_KEY).enqueue(new Callback<FirebaseSignUpResponse>(){
            @Override
            public void onResponse(Call<FirebaseSignUpResponse> call, Response<FirebaseSignUpResponse> response) {
                if (response.isSuccessful() && response.body() != null) {
                    String Token = response.body().getIdToken();
                    String RefreshToken = response.body().getRefreshToken();
                    tokenManager.saveToken(Token, RefreshToken);
                    UpdateProfileRequest updateProfileRequest = new UpdateProfileRequest(Token, name, true);
                    authApi.updateProfile(updateProfileRequest,Constants.FIREBASE_API_KEY).enqueue(new Callback<FirebaseSignUpResponse>(){
                        @Override
                        public void onResponse(Call<FirebaseSignUpResponse> call, Response<FirebaseSignUpResponse> response) {
                            if (response.isSuccessful() && response.body() != null) {
                                authApi.auth().enqueue(new Callback<ApiResponse<String>>() {
                                    @Override
                                    public void onResponse(Call<ApiResponse<String>> call, Response<ApiResponse<String>> response) {
                                        if(response.isSuccessful() && response.body() != null){
                                            callback.onSuccess(response.body().getData());
                                        }
                                        else {
                                            callback.onError("Lỗi đăng ký local");
                                        }
                                    }

                                    @Override
                                    public void onFailure(Call<ApiResponse<String>> call, Throwable t) {
                                        callback.onError(t.getMessage());
                                    }
                                }
);
                            }
                            else {
                                callback.onError("Lỗi phần update profile");
                            }
                        }

                        @Override
                        public void onFailure(Call<FirebaseSignUpResponse> call, Throwable t) {
                            callback.onError(t.getMessage());
                        }
                    });
                }
                else {
                    callback.onError("Lỗi đăng ký");
                }
            }

            @Override
            public void onFailure(Call<FirebaseSignUpResponse> call, Throwable t) {
                callback.onError(t.getMessage());
            }
        });
    };

    public void login(String email, String password, authCallBack<String> callback) {
        LoginRequest loginRequest = new LoginRequest(email, password);
        authApi.login(loginRequest, Constants.FIREBASE_API_KEY).enqueue(new Callback<FirebaseSignUpResponse>() {
            @Override
            public void onResponse(Call<FirebaseSignUpResponse> call, Response<FirebaseSignUpResponse> response) {
                if (response.isSuccessful() && response.body() != null) {
                    tokenManager.saveToken(response.body().getIdToken(), response.body().getRefreshToken());
                    userApi.getProfile().enqueue(new Callback<ApiResponse<UserResponse>>() {
                        @Override
                        public void onResponse(Call<ApiResponse<UserResponse>> call, Response<ApiResponse<UserResponse>> response) {
                            if(response.isSuccessful() && response.body() != null){
                                UserResponse userResponse = response.body().getData();
                                tokenManager.saveUserInfo(userResponse.getId(),userResponse.getEmail(),
                                        userResponse.getName(),userResponse.getAvatar_url());
                                callback.onSuccess("Thành công");
                            }
                            else {
                                callback.onError("Lỗi fetch user");
                            }
                        }

                        @Override
                        public void onFailure(Call<ApiResponse<UserResponse>> call, Throwable t) {
                            callback.onError(t.getMessage());
                        }
                    });

                }
                else {
                    callback.onError("Sai tai khoan mat khau");
                }
            }

            @Override
            public void onFailure(Call<FirebaseSignUpResponse> call, Throwable t) {
                callback.onError(t.getMessage());
            }
        });

    }



}