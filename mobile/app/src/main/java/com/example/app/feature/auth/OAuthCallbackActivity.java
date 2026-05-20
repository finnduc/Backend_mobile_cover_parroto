package com.example.app.feature.auth;

import android.content.Intent;
import android.net.Uri;
import android.os.Bundle;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;

import com.example.app.MainActivity;
import com.example.app.data.local.TokenManager;
import com.example.app.data.remote.ClerkRetrofitClient;
import com.example.app.data.remote.RetrofitClient;
import com.example.app.data.remote.api.ClerkApi;
import com.example.app.data.remote.api.UserApi;
import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.clerk.ClerkClient;
import com.example.app.data.remote.model.response.user.UserResponse;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class OAuthCallbackActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        handleIntent(getIntent());
    }

    @Override
    protected void onNewIntent(Intent intent) {
        super.onNewIntent(intent);
        handleIntent(intent);
    }

    private void handleIntent(Intent intent) {
        Uri data = intent.getData();
        if (data == null) {
            finish();
            return;
        }

        String nonce = data.getQueryParameter("rotating_token_nonce");
        if (nonce == null || nonce.isEmpty()) {
            Toast.makeText(this, "Lỗi OAuth: không có nonce", Toast.LENGTH_SHORT).show();
            finish();
            return;
        }

        ClerkApi clerkApi = ClerkRetrofitClient.getInstance().getClerkApi();
        clerkApi.getClientByNonce(nonce).enqueue(new Callback<ClerkClient>() {
            @Override
            public void onResponse(Call<ClerkClient> call, Response<ClerkClient> response) {
                if (!response.isSuccessful() || response.body() == null) {
                    Toast.makeText(OAuthCallbackActivity.this, "Xác thực thất bại", Toast.LENGTH_SHORT).show();
                    finish();
                    return;
                }

                ClerkClient client = response.body();
                String jwt = client.getJwt();
                String sessionId = client.getSessionId();

                if (jwt == null) {
                    Toast.makeText(OAuthCallbackActivity.this, "Không lấy được token", Toast.LENGTH_SHORT).show();
                    finish();
                    return;
                }

                TokenManager tokenManager = TokenManager.getInstance(OAuthCallbackActivity.this);
                tokenManager.saveToken(jwt, sessionId != null ? sessionId : "");

                UserApi userApi = RetrofitClient.getInstance(OAuthCallbackActivity.this).getUserApi();
                userApi.getProfile().enqueue(new Callback<ApiResponse<UserResponse>>() {
                    @Override
                    public void onResponse(Call<ApiResponse<UserResponse>> call2, Response<ApiResponse<UserResponse>> res) {
                        if (res.isSuccessful() && res.body() != null && res.body().getData() != null) {
                            UserResponse user = res.body().getData();
                            tokenManager.saveUserInfo(user.getId(), user.getEmail(), user.getName(), user.getAvatar_url());
                        }
                        goToMain();
                    }

                    @Override
                    public void onFailure(Call<ApiResponse<UserResponse>> call2, Throwable t) {
                        goToMain();
                    }
                });
            }

            @Override
            public void onFailure(Call<ClerkClient> call, Throwable t) {
                Toast.makeText(OAuthCallbackActivity.this, "Lỗi kết nối: " + t.getMessage(), Toast.LENGTH_SHORT).show();
                finish();
            }
        });
    }

    private void goToMain() {
        Intent intent = new Intent(this, MainActivity.class);
        intent.setFlags(Intent.FLAG_ACTIVITY_NEW_TASK | Intent.FLAG_ACTIVITY_CLEAR_TASK);
        startActivity(intent);
        finish();
    }
}
