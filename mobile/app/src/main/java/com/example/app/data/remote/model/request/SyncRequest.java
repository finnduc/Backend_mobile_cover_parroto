package com.example.app.data.remote.model.request;

import com.google.gson.annotations.SerializedName;

public class SyncRequest {


    @SerializedName("firebase_token")
    private String fireBaseToken;

    public SyncRequest(String fireBaseToken) {
        this.fireBaseToken = fireBaseToken;
    }

}
