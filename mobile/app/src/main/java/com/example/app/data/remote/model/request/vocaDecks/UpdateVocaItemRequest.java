package com.example.app.data.remote.model.request.vocaDecks;

import com.google.gson.annotations.SerializedName;

public class UpdateVocaItemRequest {
    @SerializedName("phrase")
    private String phrase;

    @SerializedName("normalized_phrase")
    private String normalizedPhrase;

    @SerializedName("meaning")
    private String meaning;

    @SerializedName("example_sentence")
    private String exampleSentence;

    @SerializedName("note")
    private String note;

    public UpdateVocaItemRequest(Integer lessonId,
                                  Integer transcriptId,
                                  String phrase,
                                  String normalizedPhrase,
                                  String meaning,
                                  String exampleSentence,
                                  String note) {
        this.phrase = phrase;
        this.normalizedPhrase = normalizedPhrase;
        this.meaning = meaning;
        this.exampleSentence = exampleSentence;
        this.note = note;
    }
}
