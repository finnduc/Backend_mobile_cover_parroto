package com.example.app.data.remote.model.response;
import com.google.gson.annotations.SerializedName;


// Định nghĩa các biến trong json response của auth/token , sau cnay thì gửi lên auth/sync

public class TokenResponse {
    @SerializedName("id_token")
    private String idToken;

    @SerializedName("refresh_token")
    private String refreshToken;

    @SerializedName("expires_in")
    private String expiresIn;

    private String email;
    public String getIdToken() {
        return idToken;
    }

    public String getRefreshToken() {
        return refreshToken;
    }

    public String getExpiresIn() {
        return expiresIn;
    }

    public String getEmail() {
        return email;
    }
}
