package com.example.app.data.repository;

import android.content.Context;

import com.example.app.data.remote.RetrofitClient;
import com.example.app.data.remote.api.ChatApi;
import com.example.app.data.remote.model.request.chat.SendMessageRequest;
import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.chat.ChatHistoryResponse;
import com.example.app.data.remote.model.response.chat.ChatMessageResponse;
import com.example.app.utils.ApiCallWrapper;
import com.example.app.utils.BaseCallback;

public class ChatRepository {
    private final ChatApi chatApi;

    public ChatRepository(Context context) {
        this.chatApi = RetrofitClient.getInstance(context).getChatApi();
    }

    public void getChatHistory(Integer beforeId, Integer limit, BaseCallback<ApiResponse<ChatHistoryResponse>> callback) {
        chatApi.getChatHistory(beforeId, limit).enqueue(new ApiCallWrapper<>(callback));
    }

    public void sendMessage(String content, BaseCallback<ApiResponse<ChatMessageResponse>> callback) {
        chatApi.sendMessage(new SendMessageRequest(content)).enqueue(new ApiCallWrapper<>(callback));
    }
}
