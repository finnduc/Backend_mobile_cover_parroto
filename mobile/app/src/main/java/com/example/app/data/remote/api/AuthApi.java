package com.example.app.data.remote.api;

import com.example.app.data.remote.model.request.auth.LoginRequest;
import com.example.app.data.remote.model.request.auth.RegisterRequest;
import com.example.app.data.remote.model.request.auth.UpdateProfileRequest;
import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.auth.FirebaseSignUpResponse;


import retrofit2.Call;
import retrofit2.http.Body;
import retrofit2.http.POST;
import retrofit2.http.Query;

public interface AuthApi {
    @POST("https://identitytoolkit.googleapis.com/v1/accounts:signUp")
    Call<FirebaseSignUpResponse> register(@Body RegisterRequest request, @Query("key") String FIREBASE_API_KEY);

    @POST("https://identitytoolkit.googleapis.com/v1/accounts:update")
    Call<FirebaseSignUpResponse> updateProfile(@Body UpdateProfileRequest updateProfileRequest, @Query("key") String FIREBASE_API_KEY);

    @POST("auth/complete-signup")
    Call<ApiResponse<String>> auth();

    @POST("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword")
    Call<FirebaseSignUpResponse> login(@Body LoginRequest request, @Query("key") String FIREBASE_API_KEY);

}





