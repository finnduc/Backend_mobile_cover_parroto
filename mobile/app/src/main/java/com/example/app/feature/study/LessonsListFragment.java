package com.example.app.feature.study;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;
import android.widget.Toast;

import androidx.fragment.app.Fragment;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.example.app.R;
import com.example.app.adapter.study.ListLessonsAdapter;
import com.example.app.data.remote.model.response.lessons.LessonsResponse;
import com.example.app.data.remote.model.response.lessons.ListLessonsResponse;
import com.example.app.data.repository.LessonsRepository;

import java.util.ArrayList;
import java.util.List;

public class LessonsListFragment extends Fragment {
    private List<LessonsResponse> lessonsResponseList = new ArrayList<>();
    private ListLessonsAdapter adapter;
    private LessonsRepository repository;
    private TextView CountNotStarted;


    public View onCreateView(LayoutInflater inflater, ViewGroup container, Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_lesson_list, container, false);
        RecyclerView recyclerView = view.findViewById(R.id.rvLessons);
        CountNotStarted = view.findViewById(R.id.tvCountNotStarted);
        LinearLayoutManager layoutManager = new LinearLayoutManager(getContext());
        recyclerView.setLayoutManager(layoutManager);
        adapter = new ListLessonsAdapter(lessonsResponseList);
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
                        int totallessons = data.getMeta().getTotal();
                        int done = 0;
                        int learning = 0;
                        CountNotStarted.setText(String.valueOf(totallessons - done - learning));
                    }
                    @Override
                    public void onError(String message) {
                        Toast.makeText(requireContext(), "Lỗi tải dữ liệu: " + message, android.widget.Toast.LENGTH_SHORT).show();
                    }
                });
    }

}
