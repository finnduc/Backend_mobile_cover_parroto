package com.example.app.data.remote.api;

import com.example.app.data.remote.model.request.chat.SendMessageRequest;
import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.chat.ChatHistoryResponse;
import com.example.app.data.remote.model.response.chat.ChatMessageResponse;

import retrofit2.Call;
import retrofit2.http.Body;
import retrofit2.http.GET;
import retrofit2.http.POST;
import retrofit2.http.Query;

public interface ChatApi {

    @GET("chat/messages")
    Call<ApiResponse<ChatHistoryResponse>> getChatHistory(
            @Query("before_id") Integer beforeId,
            @Query("limit") Integer limit
    );

    @POST("chat/messages")
    Call<ApiResponse<ChatMessageResponse>> sendMessage(
            @Body SendMessageRequest request
    );
}
