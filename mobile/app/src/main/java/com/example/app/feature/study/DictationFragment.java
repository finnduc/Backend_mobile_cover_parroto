package com.example.app.feature.study;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.webkit.WebChromeClient;
import android.webkit.WebSettings;
import android.webkit.WebView;
import android.webkit.WebViewClient;
import android.widget.Button;
import android.widget.EditText;
import android.widget.ImageButton;
import android.widget.LinearLayout;
import android.widget.TextView;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.navigation.Navigation;

import com.example.app.R;
import com.example.app.data.remote.api.TranscriptsApi;
import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.transcripts.TranscriptsResponse;
import com.example.app.data.repository.TranscriptsRepository;

import java.util.List;

import retrofit2.Call;

public class DictationFragment extends Fragment {

    private TranscriptsRepository transcriptsRepository;

    private TextView toolbarTitle;
    private TextView toolbarProgress;
    private TextView timer;
    private WebView webViewYoutube;
    private ImageButton btnClose ;
    private EditText etInput;
    private Button btnStart;
    private LinearLayout btnSpeed;
    private TextView tvSpeed;
    private ImageButton btnReplay;
    private ImageButton btnPlaySentence;
    private androidx.appcompat.widget.SwitchCompat switchAutoStop;
    private boolean isPlaying = false;
    private static final float[] SPEED_LEVELS = {0.25f, 0.5f, 0.75f, 1.0f, 1.25f, 1.5f, 1.75f, 2.0f};
    private int speedIndex = 3;
    private float currentSpeed = 1.0f;
    private int currentSentenceIndex = 0;
    private LinearLayout llSentenceNumbers;

    @Nullable
    @Override
    public View onCreateView(
            @NonNull LayoutInflater inflater,
            @Nullable ViewGroup container,
            @Nullable Bundle savedInstanceState
    ) {

        View view = inflater.inflate(R.layout.fragment_dictation, container, false);

        transcriptsRepository = new TranscriptsRepository(requireContext());
        toolbarTitle = view.findViewById(R.id.tvToolbarTitle);
        toolbarProgress = view.findViewById(R.id.tvProgress);
        timer = view.findViewById(R.id.tvTimer);
        webViewYoutube = view.findViewById(R.id.webViewYoutube);
        btnClose = view.findViewById(R.id.btnClose);
        etInput = view.findViewById(R.id.etInput);
        btnStart = view.findViewById(R.id.btnStart);
        btnSpeed = view.findViewById(R.id.btnSpeed);
        tvSpeed = view.findViewById(R.id.tvSpeed);
        btnReplay = view.findViewById(R.id.btnReplay);
        btnPlaySentence = view.findViewById(R.id.btnPlaySentence);
        switchAutoStop = view.findViewById(R.id.switchAutoStop);
        etInput.setEnabled(false);
        setupListeners();


        if (getArguments() != null) {

            String lessonTitle = getArguments().getString("lessonTitle");
            String lessonVideoUrl = getArguments().getString("lessonVideoUrl");
            int lessonDuration = getArguments().getInt("lessonDuration");
            int lessonId = getArguments().getInt("lessonId",-1);
            if (lessonId == -1){

            }

            toolbarTitle.setText(lessonTitle);
            toolbarProgress.setText("0% hoàn thành");
            timer.setText(lessonDuration + " Phút");

            setupYoutubeWebView(lessonVideoUrl);
        }

        return view;
    }

    public void fetchTranscripts(int lessonId){
        transcriptsRepository.getTranscripts(lessonId,new TranscriptsRepository.TranscriptsCallback(){
            @Override
            public void onSuccess(List<TranscriptsResponse> data) {

            }

            @Override
            public void onError(String message) {
                System.out.println("Lỗi tải transcript: " + message);
            }
        });
    }



    private void setupYoutubeWebView(String lessonVideoUrl) {

        if (lessonVideoUrl == null || lessonVideoUrl.isEmpty()) {
            return;
        }

        WebSettings settings = webViewYoutube.getSettings();

        settings.setJavaScriptEnabled(true);
        settings.setDomStorageEnabled(true);
        settings.setMediaPlaybackRequiresUserGesture(false);
        settings.setMixedContentMode(WebSettings.MIXED_CONTENT_ALWAYS_ALLOW);

        String defaultUA = settings.getUserAgentString();
        String fakedUA = defaultUA.replace("; wv", "");
        settings.setUserAgentString(fakedUA);

        webViewYoutube.setWebChromeClient(new WebChromeClient());
        webViewYoutube.setWebViewClient(new WebViewClient());

        String embedUrl = lessonVideoUrl;

        if (embedUrl.startsWith("http://")) {
            embedUrl = embedUrl.replaceFirst("http://", "https://");
        }

        if (embedUrl.contains("watch?v=")) {
            embedUrl = embedUrl.replace("watch?v=", "embed/");
        } else if (embedUrl.contains("youtu.be/")) {
            embedUrl = "https://www.youtube.com/embed/"
                    + embedUrl.substring(
                    embedUrl.lastIndexOf("/") + 1
            );
        }

        embedUrl = embedUrl.replace(
                "youtube.com",
                "youtube-nocookie.com"
        );

        String appOrigin = "https://" + requireContext().getPackageName();

        String youtubeParams =
                "controls=1&modestbranding=1&rel=0&playsinline=1&enablejsapi=1&origin=" + appOrigin;

        if (embedUrl.contains("?")) {
            embedUrl += "&" + youtubeParams;
        } else {
            embedUrl += "?" + youtubeParams;
        }

        String html =
                "<!DOCTYPE html>" +
                        "<html style='margin:0;padding:0;height:100%;'>" +
                        "<body style='margin:0;padding:0;height:100%;background:#000;'>" +
                        "<iframe width='100%' height='100%' " +
                        "style='display:block;' " +
                        "src='" + embedUrl + "' " +
                        "allow='autoplay; encrypted-media; fullscreen' " +
                        "referrerpolicy='strict-origin-when-cross-origin' " +
                        "frameborder='0' allowfullscreen>" +
                        "</iframe>" +
                        "</body></html>";

        webViewYoutube.loadDataWithBaseURL(
                appOrigin,
                html,
                "text/html",
                "utf-8",
                null
        );
    }

    @Override
    public void onDestroyView() {
        super.onDestroyView();
        if (webViewYoutube != null) {
            webViewYoutube.loadUrl("about:blank");
            webViewYoutube.onPause();
            webViewYoutube.removeAllViews();
            webViewYoutube.destroy();
            webViewYoutube = null;
        }
    }

    public void setupListeners(){
        btnClose.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if (webViewYoutube != null) {
                    webViewYoutube.loadUrl("about:blank");
                    webViewYoutube.onPause();
                }
                Navigation.findNavController(v).popBackStack();
            }
        });

        btnStart.setOnClickListener(v -> {
            if(!isPlaying){
                isPlaying = true;
                btnStart.setText("Kiềm tra");
                etInput.setEnabled(true);
                etInput.requestFocus();
            }
            else {
                String userInput = etInput.getText().toString().trim();
            }
        });

        btnSpeed.setOnClickListener(v -> {
            speedIndex = (speedIndex + 1) % SPEED_LEVELS.length;
            currentSpeed = SPEED_LEVELS[speedIndex];
            tvSpeed.setText(currentSpeed + "x");
        });


        btnReplay.setOnClickListener(v -> {
        });
    }


    public void changeVideoSpeed(float speed) {
        if (webViewYoutube != null) {
            String jsCommand = "javascript:(function() { " +
                    "var player = document.getElementById('movie_player'); " +
                    "if(player) { player.setPlaybackRate(" + speed + "); } " +
                    "})()";
            webViewYoutube.evaluateJavascript(jsCommand, null);
        }
    }

    public void replayCurrentSentence() {
        // Giả sử câu hiện tại bắt đầu ở giây thứ 10
        int startTimeInSeconds = 10;


        if (webViewYoutube != null) {
            String jsCommand = "javascript:(function() { " +
                    "var player = document.getElementById('movie_player'); " +
                    "if(player) { player.seekTo(" + startTimeInSeconds + ", true); player.playVideo(); } " +
                    "})()";
            webViewYoutube.evaluateJavascript(jsCommand, null);
        }
    }

    public void checkAnswer(String userInput) {
        // Tạm thời hardcode đáp án để test
        String targetSentence = "princess mononoke";

        if (userInput.equalsIgnoreCase(targetSentence)) {
            btnStart.setText("Chính xác! Câu tiếp theo");
            etInput.setBackgroundResource(R.drawable.bg_sentence_active);
        } else {
            btnStart.setText("Sai rồi, thử lại");
        }
    }

    public void seekVideoTo(float seconds) {
        if (webViewYoutube != null) {
            String jsCommand = "javascript:(function() { " +
                    "var player = document.getElementById('movie_player'); " +
                    "if(player) { player.seekTo(" + seconds + ", true); player.playVideo(); } " +
                    "})()";
            webViewYoutube.evaluateJavascript(jsCommand, null);
        }
    }

    public int dpToPx(int dp) {
        float density = requireContext().getResources().getDisplayMetrics().density;
        return Math.round((float) dp * density);
    }


}