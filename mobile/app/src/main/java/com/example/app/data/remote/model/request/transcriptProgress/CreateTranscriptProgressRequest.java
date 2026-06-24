package com.example.app.data.remote.model.request.transcriptProgress;

import com.google.gson.annotations.SerializedName;

public class CreateTranscriptProgressRequest {
    @SerializedName("transcript_id")
    private int transcriptId;

    @SerializedName("lesson_id")
    private int lessonId;

    public CreateTranscriptProgressRequest(int transcriptId, int lessonId) {
        this.transcriptId = transcriptId;
        this.lessonId = lessonId;
    }
}
