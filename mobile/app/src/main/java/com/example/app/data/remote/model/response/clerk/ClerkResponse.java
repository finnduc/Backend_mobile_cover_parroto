package com.example.app.data.remote.model.response.clerk;

import com.google.gson.annotations.SerializedName;

public class ClerkResponse<T> {
    @SerializedName("response")
    private T response;

    @SerializedName("client")
    private ClerkClient client;

    public T getResponse() { return response; }
    public ClerkClient getClient() { return client; }
}
