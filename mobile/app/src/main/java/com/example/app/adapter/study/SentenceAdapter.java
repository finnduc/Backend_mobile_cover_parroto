package com.example.app.adapter.study;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.example.app.R;
import com.example.app.data.remote.model.response.transcripts.TranscriptsResponse;

import java.util.ArrayList;
import java.util.List;

public class SentenceAdapter extends RecyclerView.Adapter<SentenceAdapter.SentenceViewHolder> {

    List<TranscriptsResponse> list = new ArrayList<>();


    public SentenceAdapter(List<TranscriptsResponse> newList) {
        this.list = newList;
    }

    @NonNull
    @Override
    public SentenceViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext())
                .inflate(R.layout.item_sentence_number, parent, false);
        return new SentenceViewHolder(view);
    }

    @Override
    public void onBindViewHolder(@NonNull SentenceViewHolder holder, int position) {

    }

    @Override
    public int getItemCount() {

        if (list != null) {
            return list.size();
        }
        return 0;
    }

    public class SentenceViewHolder extends RecyclerView.ViewHolder {

        public SentenceViewHolder(@NonNull View itemView) {
            super(itemView);
            this.tvSentence = itemView.findViewById(R.id.tvSentence);
        }

        TextView tvSentence;

    }


}
