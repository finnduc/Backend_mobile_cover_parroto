package com.example.app.feature.study;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.example.app.R;
import com.example.app.adapter.study.LessonSectionAdapter;
import com.example.app.data.remote.model.response.categories.CategoryResponse;
import com.example.app.data.remote.model.response.categories.ListCategoryResponse;
import com.example.app.data.remote.model.response.lessons.LessonsResponse;
import com.example.app.data.remote.model.response.lessons.ListLessonsResponse;
import com.example.app.data.repository.CategoriesRepository;
import com.example.app.data.repository.LessonsRepository;

import java.util.ArrayList;
import java.util.List;

public class StudyFragment extends Fragment {

    private RecyclerView rvLessonSections;
    private LessonSectionAdapter sectionAdapter;
    private List<LessonSection> sectionList;
    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_study, container, false);
        rvLessonSections = view.findViewById(R.id.rvLessonSections);
        rvLessonSections.setLayoutManager(new LinearLayoutManager(getContext()));
        sectionList = new ArrayList<>();
        sectionAdapter = new LessonSectionAdapter(sectionList);
        rvLessonSections.setAdapter(sectionAdapter);
        loadDataFromApi();
        return view;
    }

    private void loadDataFromApi() {
        CategoriesRepository categoryRepo = new CategoriesRepository(requireContext());
        LessonsRepository lessonsRepo = new LessonsRepository(requireContext());

        categoryRepo.getCategory(10, 1,
                new CategoriesRepository.categoryCallback<ListCategoryResponse<CategoryResponse>>() {
            @Override
            public void onSuccess(ListCategoryResponse<CategoryResponse> data) {
                if (data != null && data.getData() != null) {
                    List<CategoryResponse> categoryList = data.getData();

                    for (CategoryResponse category : categoryList) {

                        int categoryId = 0;
                        try {
                            categoryId = Integer.parseInt(category.getId());
                        } catch (NumberFormatException e) {
                            continue;
                        }
                        lessonsRepo.getLessons(5, 1, categoryId, null,
                                new LessonsRepository.lessonsCallback<ListLessonsResponse<LessonsResponse>>() {
                            @Override
                            public void onSuccess(ListLessonsResponse<LessonsResponse> lessonData) {
                                if (lessonData != null && lessonData.getData() != null) {
                                    LessonSection newSection = new LessonSection(
                                            category.getName(),
                                            lessonData.getMeta().getTotal(),
                                            lessonData.getData()
                                    );

                                    // 5. Thêm vào danh sách tổng và báo Adapter cập nhật UI
                                    sectionList.add(newSection);
                                    sectionAdapter.notifyDataSetChanged();
                                }
                            }
                            @Override
                            public void onError(String message) {
                            }
                        });
                    }
                }
            }

            @Override
            public void onError(String message) {
                Toast.makeText(getContext(), "Lỗi tải danh mục: " + message, Toast.LENGTH_SHORT).show();
            }
        });
    }
}
