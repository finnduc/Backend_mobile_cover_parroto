package com.example.app.data.remote.api;

import com.example.app.data.remote.model.response.ApiResponse;
import retrofit2.Call;
import retrofit2.http.POST;

public interface AuthApi {
    // API đồng bộ user từ Firebase xuống Database Backend
    @POST("auth/complete-signup")
    Call<ApiResponse<String>> auth();
}