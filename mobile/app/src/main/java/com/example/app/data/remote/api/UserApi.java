package com.example.app.data.remote.api;

import com.example.app.data.remote.model.request.auth.UpdateProfileRequest;
import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.user.UserResponse;

import retrofit2.Call;
import retrofit2.http.Body;
import retrofit2.http.GET;
import retrofit2.http.PUT;

public interface UserApi {

    @GET("user/profile")
    Call<ApiResponse<UserResponse>> getProfile();
    @PUT("user/profile")
    Call<ApiResponse<UserResponse>> updateProfile(@Body UpdateProfileRequest request);

}
