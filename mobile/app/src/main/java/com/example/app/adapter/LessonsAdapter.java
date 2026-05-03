package com.example.app.adapter;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.bumptech.glide.Glide;
import com.example.app.R;
import com.example.app.data.remote.model.response.lessons.LessonsResponse;

import java.util.List;

public class LessonsAdapter extends RecyclerView.Adapter<LessonsAdapter.LessonsViewHolder> {

    private List<LessonsResponse> lessonsResponseList;

    public LessonsAdapter(List<LessonsResponse> lessonsResponseList) {
        this.lessonsResponseList = lessonsResponseList;
    }

    @NonNull
    @Override
    public LessonsViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.item_lesson_list, parent, false);
        return new LessonsViewHolder(view);
    }

    @Override
    public void onBindViewHolder(@NonNull LessonsViewHolder holder, int position) {
        LessonsResponse currentLesson = lessonsResponseList.get(position);
        holder.tvTitle.setText(currentLesson.getTitle());
        holder.tvDuration.setText(String.valueOf(currentLesson.getDuration()));
        holder.tvLevel.setText(currentLesson.getLevel());
        Glide.with(holder.itemView.getContext())
                .load(currentLesson.getThumbnailUrl())
                .placeholder(R.drawable.ic_placeholder)
                .error(R.drawable.ic_error)
                .into(holder.imgThumbnail);
    }

    @Override
    public int getItemCount() {
        if (lessonsResponseList != null){
            return lessonsResponseList.size();
        }
        return 0;
    }

    public class LessonsViewHolder extends RecyclerView.ViewHolder {
        private ImageView imgThumbnail;
        private TextView tvTitle;
        private TextView tvDuration;
        private TextView tvLevel;

        public LessonsViewHolder(@NonNull View itemView) {
            super(itemView);
            imgThumbnail = itemView.findViewById(R.id.imgThumbnail);
            tvTitle = itemView.findViewById(R.id.tvTitle);
            tvDuration = itemView.findViewById(R.id.tvDuration);
            tvLevel = itemView.findViewById(R.id.tvLevel);

        }


    }
}

