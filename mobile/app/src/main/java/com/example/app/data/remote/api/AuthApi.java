package com.example.app.data.remote.api;


import com.example.app.data.remote.model.request.auth.RegisterRequest;
import com.example.app.data.remote.model.request.auth.SyncRequest;
import com.example.app.data.remote.model.request.auth.TokenRequest;
import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.auth.FirebaseSignUpResponse;
import com.example.app.data.remote.model.response.auth.SyncResponse;
import com.example.app.data.remote.model.response.auth.TokenResponse;

import retrofit2.Call;
import retrofit2.http.Body;
import retrofit2.http.POST;
import retrofit2.http.Query;

public interface AuthApi {

    @POST("auth/token")
    Call<ApiResponse<TokenResponse>> getToken(@Body TokenRequest request);

    @POST("auth/sync")
    Call<ApiResponse<SyncResponse>> synUser(@Body SyncRequest request);

    @POST("https://identitytoolkit.googleapis.com/v1/accounts:signUp")
    Call<FirebaseSignUpResponse> getRegister(@Query("key") String apiKey, @Body RegisterRequest request);
}