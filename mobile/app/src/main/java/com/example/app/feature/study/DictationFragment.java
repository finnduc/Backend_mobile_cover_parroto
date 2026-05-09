package com.example.app.feature.study;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.webkit.WebChromeClient;
import android.webkit.WebSettings;
import android.webkit.WebView;
import android.webkit.WebViewClient;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;

import com.example.app.R;

public class DictationFragment extends Fragment {

    TextView toolbarTitle;
    TextView toolbarProgress;
    TextView timer;
    WebView webViewYoutube;

    @Nullable
    @Override
    public View onCreateView(
            @NonNull LayoutInflater inflater,
            @Nullable ViewGroup container,
            @Nullable Bundle savedInstanceState
    ) {

        View view = inflater.inflate(R.layout.fragment_dictation, container, false);

        toolbarTitle = view.findViewById(R.id.tvToolbarTitle);
        toolbarProgress = view.findViewById(R.id.tvProgress);
        timer = view.findViewById(R.id.tvTimer);
        webViewYoutube = view.findViewById(R.id.webViewYoutube);

        if (getArguments() != null) {

            String lessonTitle = getArguments().getString("lessonTitle");
            String lessonVideoUrl = getArguments().getString("lessonVideoUrl");
            int lessonDuration = getArguments().getInt("lessonDuration");

            toolbarTitle.setText(lessonTitle);
            toolbarProgress.setText("0% hoàn thành");
            timer.setText(lessonDuration + " Phút");

            setupYoutubeWebView(lessonVideoUrl);
        }

        return view;
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
}