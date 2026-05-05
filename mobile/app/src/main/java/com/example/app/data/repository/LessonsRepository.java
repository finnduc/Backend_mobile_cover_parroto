package com.example.app.data.repository;

import android.content.Context;

import com.example.app.data.remote.RetrofitClient;
import com.example.app.data.remote.api.LessonsApi;
import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.lessons.LessonsResponse;
import com.example.app.data.remote.model.response.lessons.ListLessonsResponse;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class LessonsRepository {

    private LessonsApi lessonsApi;

    public LessonsRepository(Context context) {
        this.lessonsApi = RetrofitClient.getInstance(context).getLessonsApi();
    }

    public interface lessonsCallback<T>{
        void onSuccess(T data);
        void onError(String message);
    }
    public void getLessons(int limit,int page, int categoryId, String level, lessonsCallback<ListLessonsResponse<LessonsResponse>> callback){
        lessonsApi.getLessons(limit,page,categoryId,level).enqueue(new Callback<ApiResponse<ListLessonsResponse<LessonsResponse>>>() {
            @Override
            public void onResponse(Call<ApiResponse<ListLessonsResponse<LessonsResponse>>> call, Response<ApiResponse<ListLessonsResponse<LessonsResponse>>> response) {
                if(response.isSuccessful() && response.body() != null && response.body().issuccess()){
                    ListLessonsResponse<LessonsResponse> lessonsData = response.body().getData();
                    callback.onSuccess(lessonsData);
                }
                else {
                    try {
                        String errorDetail = response.errorBody() != null ? response.errorBody().string() : "Lỗi không xác định";
                        callback.onError(errorDetail);
                     } catch (Exception e){
                        callback.onError(e.getMessage());
                    }
                }
            }

            @Override
            public void onFailure(Call<ApiResponse<ListLessonsResponse<LessonsResponse>>> call, Throwable t) {
                callback.onError(t.getMessage());
            }
        });
    }

    public void getLessonsDetail(int lessonId, lessonsCallback<LessonsResponse> callback){
        lessonsApi.getLessonsDetail(lessonId).enqueue(new Callback<ApiResponse<LessonsResponse>>(){
            @Override
            public void onResponse(Call<ApiResponse<LessonsResponse>> call, Response<ApiResponse<LessonsResponse>> response) {
                if (response.isSuccessful() && response.body() != null && response.body().issuccess()){
                    LessonsResponse lessonsResponseData = response.body().getData();
                    callback.onSuccess(lessonsResponseData);
                }
                else {
                    try {
                        String errorDetail = response.errorBody() != null ? response.errorBody().string() : "Lỗi không xác định";
                        callback.onError(errorDetail);
                    } catch (Exception e){
                        callback.onError(e.getMessage());
                    }

                }
            }

            @Override
            public void onFailure(Call<ApiResponse<LessonsResponse>> call, Throwable t) {
                callback.onError(t.getMessage());
            }
        });
    }


}
