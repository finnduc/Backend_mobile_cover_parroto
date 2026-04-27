package com.example.app.data.remote.model.response;
import com.google.gson.annotations.SerializedName;


// giống bên auth/token , bên đây respone của auth/sync


public class SyncResponse {

    private int id;
    private String email;
    private String name;
    @SerializedName("avatar_url")
    private String avatarURL;

    public int getId() {
        return id;
    }

    public String getEmail() {
        return email;
    }

    public String getName() {
        return name;
    }

    public String getAvatarURL() {
        return avatarURL;
    }
}
