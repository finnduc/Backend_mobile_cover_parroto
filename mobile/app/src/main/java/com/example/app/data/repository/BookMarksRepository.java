package com.example.app.data.repository;

import android.content.Context;

import com.example.app.data.remote.RetrofitClient;
import com.example.app.data.remote.api.BookMarksApi;
import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.bookmarks.BookmarksModel;
import com.example.app.data.remote.model.response.bookmarks.BookmarksResponse;
import com.example.app.utils.ApiCallWrapper;
import com.example.app.utils.BaseCallback;

import java.util.List;

public class BookMarksRepository {
    private final BookMarksApi bookMarksApi;

    public BookMarksRepository(Context context) {
        this.bookMarksApi = RetrofitClient.getInstance(context).getBookMarksApi();
    }

    public void getBookmarks(int page, int limit, BaseCallback<ApiResponse<List<BookmarksModel>>> callback) {
        bookMarksApi.getBookmarks(page, limit).enqueue(new ApiCallWrapper<>(callback));
    }

    public void toggleBookmark(int lessonId, BaseCallback<ApiResponse<BookmarksResponse>> callback) {
        bookMarksApi.toggleBookmark(lessonId).enqueue(new ApiCallWrapper<>(callback));
    }

    public void deleteBookmark(int id, BaseCallback<ApiResponse<BookmarksResponse>> callback) {
        bookMarksApi.deleteBookmark(id).enqueue(new ApiCallWrapper<>(callback));
    }
}
