package com.example.app.utils;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class ApiCallWrapper<T> implements Callback<T> {
    private final BaseCallback<T> callback;

    public ApiCallWrapper(BaseCallback<T> callback) {
        this.callback = callback;
    }

    @Override
    public void onResponse(Call<T> call, Response<T> response) {
        if (response.isSuccessful() && response.body() != null) {
            callback.onSuccess(response.body());
        } else {
            try {
                String errorDetail = response.errorBody() != null
                        ? response.errorBody().string()
                        : "Lỗi không xác định (Code: " + response.code() + ")";
                callback.onError(errorDetail);
            } catch (Exception e) {
                String message = e.getMessage();
                callback.onError("Lỗi đọc dữ liệu: " + (message != null ? message : "Không rõ nguyên nhân"));
            }
        }
    }

    @Override
    public void onFailure(Call<T> call, Throwable t) {
        String message = t.getMessage();
        callback.onError("Lỗi kết nối: " + (message != null ? message : "Không thể kết nối đến máy chủ"));
    }
}
