package com.example.app.data.remote.model.response.chat;

import com.google.gson.annotations.SerializedName;
import java.util.List;

public class ChatHistoryResponse {
    @SerializedName("messages")
    private List<ChatMessageResponse> messages;

    @SerializedName("has_more")
    private boolean hasMore;

    @SerializedName("next_id")
    private Integer nextId;

    public List<ChatMessageResponse> getMessages() {
        return messages;
    }

    public void setMessages(List<ChatMessageResponse> messages) {
        this.messages = messages;
    }

    public boolean isHasMore() {
        return hasMore;
    }

    public void setHasMore(boolean hasMore) {
        this.hasMore = hasMore;
    }

    public Integer getNextId() {
        return nextId;
    }

    public void setNextId(Integer nextId) {
        this.nextId = nextId;
    }
}
