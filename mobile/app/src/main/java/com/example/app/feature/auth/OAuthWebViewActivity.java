package com.example.app.feature.auth;

import android.annotation.SuppressLint;
import android.content.Intent;
import android.net.Uri;
import android.os.Bundle;
import android.webkit.WebResourceRequest;
import android.webkit.WebView;
import android.webkit.WebViewClient;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;

import com.example.app.MainActivity;
import com.example.app.data.local.TokenManager;
import com.example.app.data.remote.ClerkRetrofitClient;
import com.example.app.data.remote.RetrofitClient;
import com.example.app.data.remote.api.UserApi;
import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.clerk.ClerkClient;
import com.example.app.data.remote.model.response.user.UserResponse;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class OAuthWebViewActivity extends AppCompatActivity {

    public static final String EXTRA_AUTH_URL = "auth_url";

    @SuppressLint("SetJavaScriptEnabled")
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        String authUrl = getIntent().getStringExtra(EXTRA_AUTH_URL);
        if (authUrl == null) {
            finish();
            return;
        }

        WebView webView = new WebView(this);
        setContentView(webView);

        webView.getSettings().setJavaScriptEnabled(true);
        webView.getSettings().setDomStorageEnabled(true);

        webView.setWebViewClient(new WebViewClient() {
            @Override
            public boolean shouldOverrideUrlLoading(WebView view, WebResourceRequest request) {
                String url = request.getUrl().toString();
                if (url.startsWith("engflix://oauth")) {
                    handleOAuthCallback(Uri.parse(url));
                    return true;
                }
                return false;
            }
        });

        webView.loadUrl(authUrl);
    }

    private void handleOAuthCallback(Uri uri) {
        String nonce = uri.getQueryParameter("rotating_token_nonce");
        if (nonce == null || nonce.isEmpty()) {
            Toast.makeText(this, "Lỗi OAuth: thiếu token", Toast.LENGTH_SHORT).show();
            finish();
            return;
        }

        ClerkRetrofitClient.getInstance().getClerkApi()
                .getClientByNonce(nonce)
                .enqueue(new Callback<ClerkClient>() {
                    @Override
                    public void onResponse(Call<ClerkClient> call, Response<ClerkClient> response) {
                        if (!response.isSuccessful() || response.body() == null) {
                            Toast.makeText(OAuthWebViewActivity.this, "Xác thực thất bại", Toast.LENGTH_SHORT).show();
                            finish();
                            return;
                        }

                        ClerkClient client = response.body();
                        String jwt = client.getJwt();
                        String sessionId = client.getSessionId();

                        if (jwt == null) {
                            Toast.makeText(OAuthWebViewActivity.this, "Không lấy được token", Toast.LENGTH_SHORT).show();
                            finish();
                            return;
                        }

                        TokenManager tokenManager = TokenManager.getInstance(OAuthWebViewActivity.this);
                        tokenManager.saveToken(jwt, sessionId != null ? sessionId : "");

                        UserApi userApi = RetrofitClient.getInstance(OAuthWebViewActivity.this).getUserApi();
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
                        Toast.makeText(OAuthWebViewActivity.this, "Lỗi: " + t.getMessage(), Toast.LENGTH_SHORT).show();
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
