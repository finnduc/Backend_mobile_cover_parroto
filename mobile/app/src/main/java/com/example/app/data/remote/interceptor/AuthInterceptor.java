package com.example.app.data.remote.interceptor;

import androidx.annotation.NonNull;

import com.example.app.data.local.TokenManager;
import com.example.app.data.remote.ClerkAuthBridge;

import java.io.IOException;

import okhttp3.Interceptor;
import okhttp3.Request;
import okhttp3.Response;

public class AuthInterceptor implements Interceptor {

    private TokenManager tokenManager;
    public AuthInterceptor(TokenManager tokenManager) {
        this.tokenManager = tokenManager;
    }


    @Override
    @NonNull
    public Response intercept(@NonNull Chain chain) throws IOException {
        Request originalRequest = chain.request();
        String token = getValidToken();
        if (token == null) {
            return chain.proceed(originalRequest);
        }
        Request newRequest = originalRequest.newBuilder()
                .header("Authorization", "Bearer " + token)
                .build();
        return chain.proceed(newRequest);
    }
    private String getValidToken() {
        String freshToken = ClerkAuthBridge.getTokenBlocking();
        if (freshToken != null && !freshToken.isEmpty()) {
            tokenManager.saveToken(freshToken, "");
            return freshToken;
        }
        return tokenManager.getIdToken();
    }

}

