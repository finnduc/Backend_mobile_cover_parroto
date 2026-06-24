package com.example.app.data.remote.model.response.bookmarks;

import com.google.gson.annotations.SerializedName;

public class BookmarksResponse {
    @SerializedName("id")
    private int id;

    @SerializedName("user_id")
    private String userId;

    @SerializedName("lesson_id")
    private int lessonId;

    @SerializedName("transcript_id")
    private int transcriptId;

    @SerializedName("note")
    private String note;

    @SerializedName("created_at")
    private String createdAt;

    public BookmarksResponse(
            int id,
            String userId,
            int lessonId,
            int transcriptId,
            String note,
            String createdAt
    ) {
        this.id = id;
        this.userId = userId;
        this.lessonId = lessonId;
        this.transcriptId = transcriptId;
        this.note = note;
        this.createdAt = createdAt;
    }

    public int getId() {
        return id;
    }

    public String getUserId() {
        return userId;
    }

    public int getLessonId() {
        return lessonId;
    }

    public int getTranscriptId() {
        return transcriptId;
    }

    public String getNote() {
        return note;
    }

    public String getCreatedAt() {
        return createdAt;
    }
}
