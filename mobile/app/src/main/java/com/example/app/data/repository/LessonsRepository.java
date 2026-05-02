package com.example.app.data.repository;

import android.content.Context;

import com.example.app.data.remote.RetrofitClient;
import com.example.app.data.remote.api.CategoryApi;
import com.example.app.data.remote.api.LessonsApi;
import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.lessons.Lessons;
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
    public void getLessons(int limit,int page, int categoryId, String level, lessonsCallback<ListLessonsResponse<Lessons>> callback){
        lessonsApi.getLessons(limit,page,categoryId,level).enqueue(new Callback<ApiResponse<ListLessonsResponse<Lessons>>>() {
            @Override
            public void onResponse(Call<ApiResponse<ListLessonsResponse<Lessons>>> call, Response<ApiResponse<ListLessonsResponse<Lessons>>> response) {
                if(response.isSuccessful() && response.body() != null && response.body().issuccess()){
                    ListLessonsResponse<Lessons> lessonsData = response.body().getData();
                    callback.onSuccess(lessonsData);
                }
                else {
                    try {
                        String errorDetail = response.errorBody() != null ? response.errorBody().string() : "Lỗi không xác định";
                        callback.onError(errorDetail);
                     } catch (Exception e){
                        callback.onError("error");
                    }
                }
            }

            @Override
            public void onFailure(Call<ApiResponse<ListLessonsResponse<Lessons>>> call, Throwable t) {
                callback.onError(t.getMessage());
            }
        });
    }

    public void getLessonsDetail(int lessonId, lessonsCallback<Lessons> callback){
        lessonsApi.getLessonsDetail(lessonId).enqueue(new Callback<ApiResponse<Lessons>>(){
            @Override
            public void onResponse(Call<ApiResponse<Lessons>> call, Response<ApiResponse<Lessons>> response) {
                if (response.isSuccessful() && response.body() != null && response.body().issuccess()){
                    Lessons lessonsData = response.body().getData();
                    callback.onSuccess(lessonsData);
                }
                else {
                    try {
                        String errorDetail = response.errorBody() != null ? response.errorBody().string() : "Lỗi không xác định";
                        callback.onError(errorDetail);
                    } catch (Exception e){
                        callback.onError("error");
                    }

                }
            }

            @Override
            public void onFailure(Call<ApiResponse<Lessons>> call, Throwable t) {
                callback.onError(t.getMessage());
            }
        });
    }


}
