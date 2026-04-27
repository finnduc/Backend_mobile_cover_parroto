package com.example.app.data.remote.api;

import com.example.app.data.remote.model.request.SyncRequest;
import com.example.app.data.remote.model.request.TokenRequest;
import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.SyncResponse;
import com.example.app.data.remote.model.response.TokenResponse;

import retrofit2.Call;
import retrofit2.http.Body;
import retrofit2.http.POST;

public interface AuthApi {

    @POST("auth/token")
    Call<ApiResponse<TokenResponse>> getToken(@Body TokenRequest request);

    @POST("auth/sync")
    Call<ApiResponse<SyncResponse>> syncUser(@Body SyncRequest request);
}