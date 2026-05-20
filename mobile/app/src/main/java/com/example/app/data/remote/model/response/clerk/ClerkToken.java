package com.example.app.data.remote.model.response.clerk;

import com.google.gson.annotations.SerializedName;

public class ClerkToken {
    @SerializedName("jwt")
    private String jwt;

    public String getJwt() { return jwt; }
}
