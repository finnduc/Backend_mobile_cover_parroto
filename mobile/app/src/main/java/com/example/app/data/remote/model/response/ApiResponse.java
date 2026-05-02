package com.example.app.data.remote.model.response;

import com.google.gson.annotations.SerializedName;

public class ApiResponse<T> {
    @SerializedName("data")
    private T data;
    @SerializedName("error")
    private String error;
    @SerializedName("meta")
    private metaresponse meta;
    public metaresponse getMeta(){
        return meta;
    }


    public boolean issuccess(){
        return data != null && error == null;
    }


    public T getData(){
        return data;
    }

    public String getError(){
        return error;
    }
    public static class metaresponse{
        private int limit;
        private int page;
        private int total;
        @SerializedName("total_pages")
        private int totalPages;

        public int getLimit() {
            return limit;
        }

        public int getPage() {
            return page;
        }

        public int getTotal_pages() {
            return totalPages;
        }

        public int getTotal() {
            return total;
        }
    }

}
