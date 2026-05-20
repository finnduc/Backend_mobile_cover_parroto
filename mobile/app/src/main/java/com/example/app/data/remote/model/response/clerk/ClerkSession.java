package com.example.app.data.remote.model.response.clerk;

import com.google.gson.annotations.SerializedName;

public class ClerkSession {
    @SerializedName("id")
    private String id;

    @SerializedName("last_active_token")
    private ClerkToken lastActiveToken;

    public String getId() { return id; }
    public ClerkToken getLastActiveToken() { return lastActiveToken; }
}
