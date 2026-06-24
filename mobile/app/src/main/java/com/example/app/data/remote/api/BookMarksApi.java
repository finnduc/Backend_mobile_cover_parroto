package com.example.app.data.remote.api;

import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.bookmarks.BookmarksModel;
import com.example.app.data.remote.model.response.bookmarks.BookmarksResponse;

import java.util.List;

import retrofit2.Call;
import retrofit2.http.DELETE;
import retrofit2.http.GET;
import retrofit2.http.POST;
import retrofit2.http.Path;
import retrofit2.http.Query;

public interface BookMarksApi {
    @GET("lesson-bookmarks")
    Call<ApiResponse<List<BookmarksModel>>> getBookmarks(
            @Query("page") int page,
            @Query("limit") int limit
    );

    @POST("lesson-bookmarks/{lessonId}/toggle")
    Call<ApiResponse<BookmarksResponse>> toggleBookmark(
            @Path("lessonId") int lessonId
    );

    @DELETE("lesson-bookmarks/{id}")
    Call<ApiResponse<BookmarksResponse>> deleteBookmark(
            @Path("id") int id
    );
}
