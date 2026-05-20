package com.example.app.data.remote.api;

import com.example.app.data.remote.model.response.clerk.ClerkClient;
import com.example.app.data.remote.model.response.clerk.ClerkResponse;
import com.example.app.data.remote.model.response.clerk.ClerkSignInAttempt;

import retrofit2.Call;
import retrofit2.http.Field;
import retrofit2.http.FormUrlEncoded;
import retrofit2.http.GET;
import retrofit2.http.POST;
import retrofit2.http.Path;
import retrofit2.http.Query;

public interface ClerkApi {

    @FormUrlEncoded
    @POST("v1/client/sign_ins")
    Call<ClerkResponse<ClerkSignInAttempt>> createSignIn(
            @Field("identifier") String identifier
    );

    @FormUrlEncoded
    @POST("v1/client/sign_ins/{id}/attempt_first_factor")
    Call<ClerkResponse<ClerkSignInAttempt>> attemptPassword(
            @Path("id") String signInId,
            @Field("strategy") String strategy,
            @Field("password") String password
    );

    @FormUrlEncoded
    @POST("v1/client/sign_ups")
    Call<ClerkResponse<ClerkSignInAttempt>> createSignUp(
            @Field("email_address") String email,
            @Field("password") String password,
            @Field("first_name") String firstName
    );

    @FormUrlEncoded
    @POST("v1/client/sign_ins")
    Call<ClerkResponse<ClerkSignInAttempt>> createOAuthSignIn(
            @Field("strategy") String strategy,
            @Field("redirect_url") String redirectUrl,
            @Field("action_complete_redirect_url") String actionCompleteRedirectUrl
    );

    @GET("v1/client")
    Call<ClerkClient> getClientByNonce(
            @Query("rotating_token_nonce") String nonce
    );
}
