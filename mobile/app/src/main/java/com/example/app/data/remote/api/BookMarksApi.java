package com.example.app.data.remote.api;

import com.example.app.data.remote.model.request.bookmarks.CreateBookMarksRequest;
import com.example.app.data.remote.model.request.note.UpdateNoteRequest;
import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.bookmarks.BookmarksModel;
import com.example.app.data.remote.model.response.bookmarks.BookmarksResponse;

import java.util.List;

import retrofit2.Call;
import retrofit2.http.Body;
import retrofit2.http.DELETE;
import retrofit2.http.GET;
import retrofit2.http.POST;
import retrofit2.http.Path;
import retrofit2.http.Query;
import retrofit2.http.PUT;

public interface BookMarksApi {
    @GET("transcript-bookmarks")
    Call<ApiResponse<List<BookmarksModel>>>
    getBookmarks(
            @Query("lesson_id") Integer lessonId,
                 @Query("limit") String limit,
                 @Query("page") String page);

    @GET("transcript-bookmarks/{lessonId}")
    Call<ApiResponse<BookmarksModel>>
    getBookmarkByLessonId(@Path("lessonId") int lessonId, @Query("page") int page, @Query("limit") int limit);

    @POST("transcript-bookmarks")
    Call<ApiResponse<BookmarksResponse>>
    createBookmark(@Body CreateBookMarksRequest request);
    @PUT("transcript-bookmarks/{transcriptId}")
    Call<ApiResponse<BookmarksResponse>>
    updateBookmark(@Path("transcriptId") int transcriptId,
                   @Body UpdateNoteRequest request);
    @DELETE("transcript-bookmarks/{transcriptId}")
    Call<ApiResponse<BookmarksResponse>>
    deleteBookmark(@Path("transcriptId") int transcriptId);

}
