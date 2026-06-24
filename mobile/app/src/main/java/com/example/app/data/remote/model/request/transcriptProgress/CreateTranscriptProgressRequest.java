package com.example.app.data.remote.model.request.transcriptProgress;

import com.google.gson.annotations.SerializedName;

public class CreateTranscriptProgressRequest {
    @SerializedName("transcript_id")
    private int transcriptId;

    @SerializedName("lesson_id")
    private int lessonId;

    @SerializedName("best_score")
    private Double bestScore;

    @SerializedName("feedback")
    private String feedback;

    public CreateTranscriptProgressRequest(int transcriptId, int lessonId) {
        this.transcriptId = transcriptId;
        this.lessonId = lessonId;
    }

    public CreateTranscriptProgressRequest(int transcriptId, int lessonId, Double bestScore, String feedback) {
        this.transcriptId = transcriptId;
        this.lessonId = lessonId;
        this.bestScore = bestScore;
        this.feedback = feedback;
    }
}
