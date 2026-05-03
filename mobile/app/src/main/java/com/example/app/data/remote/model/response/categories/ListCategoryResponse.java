package com.example.app.data.remote.model.response.categories;

import com.example.app.data.remote.model.response.ApiResponse;

import java.util.List;

public class ListCategoryResponse<T> {
    private List<T> data;
    private ApiResponse.metaresponse meta;

    public List<T> getData() {
        return data;
    }

    public ApiResponse.metaresponse getMeta() {
        return meta;
    }
}
