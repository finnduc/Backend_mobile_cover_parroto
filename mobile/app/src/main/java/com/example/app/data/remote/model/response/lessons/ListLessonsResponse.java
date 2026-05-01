package com.example.app.data.remote.model.response.lessons;

import com.example.app.data.remote.model.response.ApiResponse;
import com.google.gson.annotations.SerializedName;

import java.util.List;

public class ListLessonsResponse<T> {
    @SerializedName("data")
    private List<T> data;
    @SerializedName("meta")
    private ApiResponse.metaresponse meta ;


    public List<T> getData() {
        return data;
    }

    public ApiResponse.metaresponse getMeta() {
        return meta;
    }
}
