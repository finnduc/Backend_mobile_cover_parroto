package com.example.app.data.remote.model.response.pronunciation;

import com.google.gson.annotations.SerializedName;

public class PronunciationProgressResponse {
    @SerializedName("user_id")
    private String userId;

    @SerializedName("lesson_id")
    private int lessonId;

    @SerializedName("transcript_id")
    private int transcriptId;

    @SerializedName("best_score")
    private Double bestScore;

    @SerializedName("completed_at")
    private String completedAt;

    private String feedback;
    public String getFeedback() {
        return feedback;
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

    public Double getBestScore() {
        return bestScore;
    }

    public String getCompletedAt() {
        return completedAt;
    }

    public void setBestScore(Double bestScore) {
        this.bestScore = bestScore;
    }

    public void setFeedback(String feedback) {
        this.feedback = feedback;
    }
}
