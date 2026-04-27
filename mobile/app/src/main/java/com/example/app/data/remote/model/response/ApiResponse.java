package com.example.app.data.remote.model.response;

public class ApiResponse<T> {

    private T data;
    private ApiError error;

    public T getData(){
        return data;
    }

    public ApiError getError(){
        return error;
    }

    public boolean issuccess(){
        return data != null && error == null;
    }

    public static class ApiError {
        private String message;
        private int code;
        public String getMessage() {
            return message;
        }
        public int getCode() {
            return code;
        }
    }

}
