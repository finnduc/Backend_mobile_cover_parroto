package com.example.app.data.remote.api;

import com.example.app.data.remote.model.request.transcriptProgress.CreateTranscriptProgressRequest;
import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.pronunciation.PronunciationProgressResponse;
import com.example.app.data.remote.model.response.pronunciation.PronunciationResponse;

import java.util.List;

import okhttp3.MultipartBody;
import okhttp3.RequestBody;
import retrofit2.Call;
import retrofit2.http.GET;
import retrofit2.http.Multipart;
import retrofit2.http.POST;
import retrofit2.http.Part;
import retrofit2.http.Body;
import retrofit2.http.Query;

public interface PronunciationApi {
    @Multipart
    @POST("shadowing-status/transcribe")
    Call<ApiResponse<PronunciationResponse>> assessPronunciation(
            @Part MultipartBody.Part audio,
            @Part("referenceText") RequestBody referenceText,
            @Part("lessonId") RequestBody lessonId,
            @Part("transcriptId") RequestBody transcriptId
    );

    @GET("shadowing-status")
    Call<ApiResponse<List<PronunciationProgressResponse>>> getPronunciationProgress(
            @Query("lesson_id") int lessonId
    );

   @POST("shadowing-status")
   Call<ApiResponse<PronunciationProgressResponse>> updatePronunciationProgress(
           @Body CreateTranscriptProgressRequest request
   );

}
