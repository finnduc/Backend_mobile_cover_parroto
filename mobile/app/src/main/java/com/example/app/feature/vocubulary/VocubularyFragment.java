package com.example.app.feature.vocubulary;

import android.os.Bundle;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;

import com.example.app.R;
import com.example.app.data.remote.RetrofitClient;
import com.example.app.data.remote.api.LessonsApi;
import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.lessons.LessonsResponse;
import com.example.app.data.remote.model.response.lessons.ListLessonsResponse;
import com.google.gson.Gson;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class VocubularyFragment extends Fragment {
    private TextView textView;
    private static final String TAG = "VocabularyFragment";
    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_vocabulary, container, false);
        textView = view.findViewById(R.id.Test);
        textView.setText("Vocabulary");
        fetchSingleLessonTest(1);
        return(view);
    }

    private void fetchLessonsTest() {
        LessonsApi lessonsApi = RetrofitClient.getInstance(requireContext()).getLessonsApi();
        lessonsApi.getLessons(10, 1, null,null).enqueue(new Callback<ApiResponse<ListLessonsResponse<LessonsResponse>>>() {
            @Override
            public void onResponse(Call<ApiResponse<ListLessonsResponse<LessonsResponse>>> call, Response<ApiResponse<ListLessonsResponse<LessonsResponse>>> response) {
                if (response.isSuccessful() && response.body() != null) {
                    String jsonResponse = new Gson().toJson(response.body());
                    Log.d(TAG, "API Success: " + jsonResponse);
                    textView.setText("THÀNH CÔNG!\n\n" + jsonResponse);
                } else {
                    String errorMsg = "LỖI: Response không thành công. Code: " + response.code();
                    Log.e(TAG, errorMsg);
                    textView.setText(errorMsg);
                }
            }

            @Override
            public void onFailure(Call<ApiResponse<ListLessonsResponse<LessonsResponse>>> call, Throwable t) {
                String errorMsg = "THẤT BẠI (Lỗi mạng/Parse): " + t.getMessage();
                Log.e(TAG, errorMsg, t);
                textView.setText(errorMsg);
            }

        });
    }
    private void fetchSingleLessonTest(int idToFetch) {
        LessonsApi lessonsApi = RetrofitClient.getInstance(requireContext()).getLessonsApi();

        // Gọi hàm lấy bài học số idToFetch
        lessonsApi.getLessonsDetail(idToFetch).enqueue(new Callback<ApiResponse<LessonsResponse>>() {
            @Override
            public void onResponse(Call<ApiResponse<LessonsResponse>> call, Response<ApiResponse<LessonsResponse>> response) {
                if (response.isSuccessful() && response.body() != null && response.body().issuccess()) {

                    // Lấy đối tượng bài học ra
                    LessonsResponse lesson = response.body().getData();

                    String resultText = "THÀNH CÔNG!\n\n" +
                            "ID: " + lesson.getId() + "\n" +
                            "Tiêu đề: " + lesson.getTitle() + "\n" +
                            "Độ khó: " + lesson.getLevel() + "\n"+
                            "Desciption:" + lesson.getDescription();

                    textView.setText(resultText);

                } else {
                    textView.setText("LỖI: Không tìm thấy bài học hoặc có lỗi xảy ra.");
                }
            }

            @Override
            public void onFailure(Call<ApiResponse<LessonsResponse>> call, Throwable t) {
                textView.setText("THẤT BẠI (Lỗi mạng): " + t.getMessage());
            }
        });
    }
}
