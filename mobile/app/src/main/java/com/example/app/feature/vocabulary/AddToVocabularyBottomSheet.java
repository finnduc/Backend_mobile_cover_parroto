package com.example.app.feature.vocabulary;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ArrayAdapter;
import android.widget.Button;
import android.widget.EditText;
import android.widget.Spinner;
import android.widget.TextView;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;

import com.example.app.R;
import com.example.app.data.remote.model.request.vocaDecks.AddItemsToDeckRequest;
import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.vocabulary.VocaDecksResponse;
import com.example.app.data.remote.model.response.vocabulary.VocaItemsResponse;
import com.example.app.data.repository.VocabularyRepository;
import com.example.app.utils.BaseCallback;
import com.google.android.flexbox.FlexboxLayout;
import com.google.android.material.bottomsheet.BottomSheetDialogFragment;

import java.util.ArrayList;
import java.util.List;

public class AddToVocabularyBottomSheet extends BottomSheetDialogFragment {

    private static final String ARG_TRANSCRIPT_ID = "transcript_id";
    private static final String ARG_LESSON_ID = "lesson_id";
    private static final String ARG_TRANSCRIPT_TEXT = "transcript_text";

    private int transcriptId;
    private int lessonId;
    private String transcriptText;

    private FlexboxLayout flexboxWords;
    private Spinner spinnerDecks;
    private EditText etMeaning;
    private EditText etExampleSentence;
    private Button btnSubmit;

    private VocabularyRepository vocabularyRepository;
    private List<VocaDecksResponse> decks = new ArrayList<>();
    private List<String> selectedWords = new ArrayList<>();
    private List<String> allWords = new ArrayList<>();

    public static AddToVocabularyBottomSheet newInstance(int transcriptId, int lessonId, String transcriptText) {
        AddToVocabularyBottomSheet fragment = new AddToVocabularyBottomSheet();
        Bundle args = new Bundle();
        args.putInt(ARG_TRANSCRIPT_ID, transcriptId);
        args.putInt(ARG_LESSON_ID, lessonId);
        args.putString(ARG_TRANSCRIPT_TEXT, transcriptText);
        fragment.setArguments(args);
        return fragment;
    }

    @Override
    public void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        if (getArguments() != null) {
            transcriptId = getArguments().getInt(ARG_TRANSCRIPT_ID);
            lessonId = getArguments().getInt(ARG_LESSON_ID);
            transcriptText = getArguments().getString(ARG_TRANSCRIPT_TEXT);
        }
    }

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.bottom_sheet_add_to_vocabulary, container, false);

        vocabularyRepository = new VocabularyRepository(requireContext());

        flexboxWords = view.findViewById(R.id.flexboxWords);
        spinnerDecks = view.findViewById(R.id.spinnerDecks);
        etMeaning = view.findViewById(R.id.etMeaning);
        etExampleSentence = view.findViewById(R.id.etExampleSentence);
        btnSubmit = view.findViewById(R.id.btnSubmit);

        etExampleSentence.setText(transcriptText);

        setupWordChips();
        loadDecks();

        btnSubmit.setOnClickListener(v -> submit());

        return view;
    }

    private void setupWordChips() {
        String[] words = transcriptText.split("\\s+");
        allWords.clear();
        selectedWords.clear();

        for (String word : words) {
            allWords.add(word);
            selectedWords.add(word);
        }

        refreshWordChips();
    }

    private void refreshWordChips() {
        flexboxWords.removeAllViews();

        for (int i = 0; i < allWords.size(); i++) {
            String word = allWords.get(i);
            boolean isSelected = selectedWords.contains(word);

            TextView chip = new TextView(requireContext());
            chip.setText(word);
            chip.setPadding(24, 12, 24, 12);
            chip.setBackgroundResource(isSelected ? R.drawable.bg_chip_selected : R.drawable.bg_chip_unselected);
            chip.setTextColor(isSelected ? 0xFFFFFFFF : 0xFF333333);

            final int index = i;
            chip.setOnClickListener(v -> {
                if (selectedWords.contains(word)) {
                    selectedWords.remove(word);
                } else {
                    selectedWords.add(word);
                }
                refreshWordChips();
            });

            FlexboxLayout.LayoutParams params = new FlexboxLayout.LayoutParams(
                    ViewGroup.LayoutParams.WRAP_CONTENT,
                    ViewGroup.LayoutParams.WRAP_CONTENT
            );
            params.setMargins(8, 8, 8, 8);
            chip.setLayoutParams(params);

            flexboxWords.addView(chip);
        }
    }

    private void loadDecks() {
        vocabularyRepository.getMyVocabularyDecks(new BaseCallback<ApiResponse<List<VocaDecksResponse>>>() {
            @Override
            public void onSuccess(ApiResponse<List<VocaDecksResponse>> data) {
                if (!isAdded() || data == null || data.getData() == null) return;

                decks.clear();
                decks.addAll(data.getData());

                List<String> deckNames = new ArrayList<>();
                for (VocaDecksResponse deck : decks) {
                    deckNames.add(deck.getName());
                }

                ArrayAdapter<String> adapter = new ArrayAdapter<>(
                        requireContext(),
                        android.R.layout.simple_spinner_item,
                        deckNames
                );
                adapter.setDropDownViewResource(android.R.layout.simple_spinner_dropdown_item);
                spinnerDecks.setAdapter(adapter);
            }

            @Override
            public void onError(String message) {
                if (isAdded()) {
                    Toast.makeText(requireContext(), "Lỗi tải bộ từ vựng: " + message, Toast.LENGTH_SHORT).show();
                }
            }
        });
    }

    private void submit() {
        if (decks.isEmpty()) {
            Toast.makeText(requireContext(), "Vui lòng chọn bộ từ vựng", Toast.LENGTH_SHORT).show();
            return;
        }

        if (selectedWords.isEmpty()) {
            Toast.makeText(requireContext(), "Vui lòng chọn từ", Toast.LENGTH_SHORT).show();
            return;
        }

        int deckPosition = spinnerDecks.getSelectedItemPosition();
        if (deckPosition < 0 || deckPosition >= decks.size()) {
            Toast.makeText(requireContext(), "Vui lòng chọn bộ từ vựng", Toast.LENGTH_SHORT).show();
            return;
        }

        int deckId = decks.get(deckPosition).getId();
        String phrase = String.join(" ", selectedWords);
        String normalizedPhrase = phrase.toLowerCase().trim();
        String meaning = etMeaning.getText().toString().trim();
        String exampleSentence = etExampleSentence.getText().toString().trim();

        AddItemsToDeckRequest request = new AddItemsToDeckRequest(
                lessonId,
                transcriptId,
                phrase,
                normalizedPhrase,
                meaning,
                exampleSentence,
                ""
        );

        btnSubmit.setEnabled(false);

        vocabularyRepository.addVocaItemToDeck(deckId, request, new BaseCallback<ApiResponse<VocaItemsResponse>>() {
            @Override
            public void onSuccess(ApiResponse<VocaItemsResponse> data) {
                if (isAdded()) {
                    Toast.makeText(requireContext(), "Đã thêm vào bộ từ vựng", Toast.LENGTH_SHORT).show();
                    dismiss();
                }
            }

            @Override
            public void onError(String message) {
                if (isAdded()) {
                    btnSubmit.setEnabled(true);
                    Toast.makeText(requireContext(), "Lỗi: " + message, Toast.LENGTH_SHORT).show();
                }
            }
        });
    }
}
