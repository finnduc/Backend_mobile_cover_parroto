package com.example.app.data.remote.model.request;

import com.google.gson.annotations.SerializedName;

public class SyncRequest {


    @SerializedName("firebase_token")
    private String fireBaseToken;
    @SerializedName("name")
    private String name;

    public SyncRequest(String fireBaseToken, String name) {
        this.fireBaseToken = fireBaseToken;
        this.name = name;
    }

}
