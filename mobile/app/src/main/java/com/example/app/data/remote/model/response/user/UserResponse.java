package com.example.app.data.remote.model.response.user;

import com.google.gson.annotations.SerializedName;

public class UserResponse {
    @SerializedName("avatar_url")
    private String avatarUrl;
    private String email;
    private String id;
    private String name;

    public String getAvatar_url() {
        return avatarUrl;
    }

    public String getEmail() {
        return email;
    }

    public String getId() {
        return id;
    }

    public String getName() {
        return name;
    }
}
