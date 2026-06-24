package com.example.app.data.repository;

import android.content.Context;

import com.example.app.data.remote.RetrofitClient;
import com.example.app.data.remote.api.BookMarksApi;
import com.example.app.data.remote.model.request.bookmarks.CreateBookMarksRequest;
import com.example.app.data.remote.model.request.note.UpdateNoteRequest;
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

    public void getTranscriptBookmarks(int page, int limit, BaseCallback<ApiResponse<List<BookmarksModel>>> callback) {
        bookMarksApi.getTranscriptBookmarks(page, limit).enqueue(new ApiCallWrapper<>(callback));
    }

    public void createBookmark(
            int lessonId,
            CreateBookMarksRequest request,
            BaseCallback<ApiResponse<BookmarksResponse>> callback
    ) {
        bookMarksApi.createTranscriptBookmark(request).enqueue(new ApiCallWrapper<>(callback));
    }

    public void updateBookmark(
            int transcriptId,
            UpdateNoteRequest request,
            BaseCallback<ApiResponse<BookmarksResponse>> callback
    ) {
        bookMarksApi.updateTranscriptBookmark(transcriptId, request).enqueue(new ApiCallWrapper<>(callback));
    }

    public void deleteBookmark(int id, BaseCallback<ApiResponse<BookmarksResponse>> callback) {
        bookMarksApi.deleteBookmark(id).enqueue(new ApiCallWrapper<>(callback));
    }

    public void deleteTranscriptBookmark(int transcriptId, BaseCallback<ApiResponse<BookmarksResponse>> callback) {
        bookMarksApi.deleteTranscriptBookmark(transcriptId).enqueue(new ApiCallWrapper<>(callback));
    }
}
