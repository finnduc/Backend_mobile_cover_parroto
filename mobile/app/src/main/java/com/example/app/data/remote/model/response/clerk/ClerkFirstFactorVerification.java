package com.example.app.data.remote.model.response.clerk;

import com.google.gson.annotations.SerializedName;

public class ClerkFirstFactorVerification {
    @SerializedName("strategy")
    private String strategy;

    @SerializedName("status")
    private String status;

    @SerializedName("external_verification_redirect_url")
    private String authorizationUrl;

    public String getAuthorizationUrl() { return authorizationUrl; }
    public String getStrategy() { return strategy; }
    public String getStatus() { return status; }
}
