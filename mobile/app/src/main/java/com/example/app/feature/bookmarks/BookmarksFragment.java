package com.example.app.feature.bookmarks;

import android.os.Bundle;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageButton;
import android.widget.ImageView;
import android.widget.LinearLayout;
import android.widget.TextView;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.navigation.Navigation;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.example.app.R;
import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.bookmarks.BookmarksModel;
import com.example.app.data.remote.model.response.bookmarks.BookmarksResponse;
import com.example.app.data.remote.model.response.lessons.LessonsResponse;
import com.example.app.data.repository.BookMarksRepository;
import com.example.app.data.repository.LessonsRepository;
import com.example.app.dialog.ChooseModeBottomSheet;
import com.example.app.utils.BaseCallback;
import com.bumptech.glide.Glide;

import java.util.ArrayList;
import java.util.List;

public class BookmarksFragment extends Fragment {

    private static final String TAG = "BookmarksFragment";

    private BookMarksRepository bookMarksRepository;
    private LessonsRepository lessonsRepository;
    private RecyclerView rvBookmarks;
    private LinearLayout layoutEmpty;
    private BookmarkAdapter adapter;
    private List<BookmarkItem> bookmarkItems = new ArrayList<>();

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_bookmarks, container, false);

        bookMarksRepository = new BookMarksRepository(requireContext());
        lessonsRepository = new LessonsRepository(requireContext());

        rvBookmarks = view.findViewById(R.id.rvBookmarks);
        layoutEmpty = view.findViewById(R.id.layout_empty);

        rvBookmarks.setLayoutManager(new LinearLayoutManager(requireContext()));
        adapter = new BookmarkAdapter();
        rvBookmarks.setAdapter(adapter);

        loadBookmarks();

        return view;
    }

    private void loadBookmarks() {
        bookMarksRepository.getBookmarks(1, 100, new BaseCallback<ApiResponse<List<BookmarksModel>>>() {
            @Override
            public void onSuccess(ApiResponse<List<BookmarksModel>> data) {
                if (!isAdded() || getView() == null) return;

                if (data == null || data.getData() == null || data.getData().isEmpty()) {
                    showEmpty();
                    return;
                }

                bookmarkItems.clear();
                for (BookmarksModel bookmark : data.getData()) {
                    loadLessonForBookmark(bookmark);
                }
            }

            @Override
            public void onError(String message) {
                Log.e(TAG, "Error loading bookmarks: " + message);
                if (isAdded()) {
                    showEmpty();
                }
            }
        });
    }

    private void loadLessonForBookmark(BookmarksModel bookmark) {
        lessonsRepository.getLessonsDetail(bookmark.getLessonId(), new LessonsRepository.lessonsCallback<LessonsResponse>() {
            @Override
            public void onSuccess(LessonsResponse lesson) {
                if (!isAdded()) return;

                BookmarkItem item = new BookmarkItem();
                item.bookmarkId = bookmark.getId();
                item.lessonId = bookmark.getLessonId();
                item.title = lesson.getTitle();
                item.description = lesson.getDescription();
                item.thumbnailUrl = lesson.getThumbnailUrl();
                item.videoUrl = lesson.getVideoUrl();
                item.duration = lesson.getDuration();
                item.level = lesson.getLevel();

                bookmarkItems.add(item);
                adapter.notifyDataSetChanged();

                if (!bookmarkItems.isEmpty()) {
                    rvBookmarks.setVisibility(View.VISIBLE);
                    layoutEmpty.setVisibility(View.GONE);
                }
            }

            @Override
            public void onError(String message) {
                Log.e(TAG, "Error loading lesson: " + message);
            }
        });
    }

    private void showEmpty() {
        rvBookmarks.setVisibility(View.GONE);
        layoutEmpty.setVisibility(View.VISIBLE);
    }

    private void removeBookmark(int position) {
        BookmarkItem item = bookmarkItems.get(position);
        bookMarksRepository.deleteBookmark(item.bookmarkId, new BaseCallback<ApiResponse<BookmarksResponse>>() {
            @Override
            public void onSuccess(ApiResponse<BookmarksResponse> data) {
                if (!isAdded()) return;
                bookmarkItems.remove(position);
                adapter.notifyDataSetChanged();
                if (bookmarkItems.isEmpty()) {
                    showEmpty();
                }
                Toast.makeText(requireContext(), "Đã bỏ lưu bài học", Toast.LENGTH_SHORT).show();
            }

            @Override
            public void onError(String message) {
                Log.e(TAG, "Error removing bookmark: " + message);
                if (isAdded()) {
                    Toast.makeText(requireContext(), "Lỗi: " + message, Toast.LENGTH_SHORT).show();
                }
            }
        });
    }

    private class BookmarkItem {
        int bookmarkId;
        int lessonId;
        String title;
        String description;
        String thumbnailUrl;
        String videoUrl;
        int duration;
        String level;
    }

    private class BookmarkAdapter extends RecyclerView.Adapter<BookmarkAdapter.ViewHolder> {

        @NonNull
        @Override
        public ViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
            View view = LayoutInflater.from(parent.getContext())
                    .inflate(R.layout.item_bookmark, parent, false);
            return new ViewHolder(view);
        }

        @Override
        public void onBindViewHolder(@NonNull ViewHolder holder, int position) {
            BookmarkItem item = bookmarkItems.get(position);

            holder.tvTitle.setText(item.title);
            holder.tvDescription.setText(item.description);

            if (item.thumbnailUrl != null && !item.thumbnailUrl.isEmpty()) {
                Glide.with(holder.itemView.getContext())
                        .load(item.thumbnailUrl)
                        .placeholder(R.drawable.ic_placeholder)
                        .error(R.drawable.ic_error)
                        .into(holder.ivThumbnail);
            }

            holder.btnRemoveBookmark.setOnClickListener(v -> {
                removeBookmark(holder.getAdapterPosition());
            });

            holder.itemView.setOnClickListener(v -> {
                ChooseModeBottomSheet bottomSheet = new ChooseModeBottomSheet();
                Bundle bundle = new Bundle();
                bundle.putInt("lessonId", item.lessonId);
                bundle.putString("lessonTitle", item.title);
                bundle.putString("lessonDescription", item.description);
                bundle.putString("lessonThumbnailUrl", item.thumbnailUrl);
                bundle.putString("lessonVideoUrl", item.videoUrl);
                bundle.putInt("lessonDuration", item.duration);
                bundle.putString("lessonLevel", item.level);
                bottomSheet.setArguments(bundle);
                bottomSheet.show(getChildFragmentManager(), "ChooseModeBottomSheet");
            });
        }

        @Override
        public int getItemCount() {
            return bookmarkItems.size();
        }

        class ViewHolder extends RecyclerView.ViewHolder {
            ImageView ivThumbnail;
            TextView tvTitle;
            TextView tvDescription;
            ImageButton btnRemoveBookmark;

            ViewHolder(View view) {
                super(view);
                ivThumbnail = view.findViewById(R.id.ivThumbnail);
                tvTitle = view.findViewById(R.id.tvTitle);
                tvDescription = view.findViewById(R.id.tvDescription);
                btnRemoveBookmark = view.findViewById(R.id.btnRemoveBookmark);
            }
        }
    }
}
