package com.example.app.data.remote.model.response.clerk;

import com.google.gson.annotations.SerializedName;

public class ClerkSignInAttempt {
    @SerializedName("id")
    private String id;

    @SerializedName("status")
    private String status;

    @SerializedName("first_factor_verification")
    private ClerkFirstFactorVerification firstFactorVerification;

    public String getId() { return id; }
    public String getStatus() { return status; }
    public ClerkFirstFactorVerification getFirstFactorVerification() { return firstFactorVerification; }
}
