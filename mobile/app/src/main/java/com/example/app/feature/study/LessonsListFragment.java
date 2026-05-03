package com.example.app.feature.study;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Toast;

import androidx.fragment.app.Fragment;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.example.app.R;
import com.example.app.adapter.LessonsAdapter;
import com.example.app.data.remote.model.response.lessons.LessonsResponse;
import com.example.app.data.remote.model.response.lessons.ListLessonsResponse;
import com.example.app.data.repository.LessonsRepository;

import java.util.ArrayList;
import java.util.List;

public class LessonsListFragment extends Fragment {
    private List<LessonsResponse> lessonsResponseList = new ArrayList<>();
    private LessonsAdapter adapter;
    private LessonsRepository repository;



    public View onCreateView(LayoutInflater inflater, ViewGroup container, Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_lesson_list, container, false);
        RecyclerView recyclerView = view.findViewById(R.id.rvLessons);
        LinearLayoutManager layoutManager = new LinearLayoutManager(getContext());
        recyclerView.setLayoutManager(layoutManager);
        adapter = new LessonsAdapter(lessonsResponseList);
        repository = new LessonsRepository(requireContext());
        recyclerView.setAdapter(adapter);
        fetchLessons();
        return view;
    }

    public void fetchLessons(){
        repository.getLessons(100, 1, 1, null,
                new LessonsRepository.lessonsCallback<ListLessonsResponse<LessonsResponse>>(){
                    @Override
                    public void onSuccess(ListLessonsResponse<LessonsResponse> data) {
                        lessonsResponseList.clear();
                        if (data != null && data.getData() != null) {
                            lessonsResponseList.addAll(data.getData());
                        }
                        adapter.notifyDataSetChanged();
                    }
                    @Override
                    public void onError(String message) {
                        Toast.makeText(requireContext(), "Lỗi tải dữ liệu: " + message, android.widget.Toast.LENGTH_SHORT).show();
                    }
                });
    }
}
