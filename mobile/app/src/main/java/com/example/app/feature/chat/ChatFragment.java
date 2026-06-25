package com.example.app.feature.chat;

import android.net.Uri;
import android.os.Bundle;
import android.os.Handler;
import android.os.Looper;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.EditText;
import android.widget.ImageButton;
import android.widget.ProgressBar;
import android.widget.TextView;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.navigation.Navigation;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.example.app.R;
import com.example.app.adapter.chat.ChatAdapter;
import com.example.app.data.local.TokenManager;
import com.example.app.data.remote.ClerkAuthBridge;
import com.example.app.data.remote.model.response.ApiResponse;
import com.example.app.data.remote.model.response.chat.ChatHistoryResponse;
import com.example.app.data.remote.model.response.chat.ChatMessageResponse;
import com.example.app.data.repository.ChatRepository;
import com.example.app.utils.BaseCallback;
import com.example.app.utils.Constants;
import com.google.gson.Gson;

import java.util.Collections;
import java.util.List;
import java.util.concurrent.TimeUnit;

import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;
import okhttp3.sse.EventSource;
import okhttp3.sse.EventSourceListener;
import okhttp3.sse.EventSources;

public class ChatFragment extends Fragment {

    private enum ConnectionStatus {
        CONNECTED,
        CONNECTING,
        DISCONNECTED,
        ERROR
    }

    private ChatRepository chatRepository;
    private ChatAdapter chatAdapter;
    
    private RecyclerView rvMessages;
    private EditText etMessage;
    private ImageButton btnSend;
    private View viewConnectionStatus;
    private TextView tvConnectionStatus;
    private ProgressBar pbLoadingMore;

    private OkHttpClient okHttpClient;
    private EventSource eventSource;
    private Gson gson;

    private final Handler mainHandler = new Handler(Looper.getMainLooper());
    private int reconnectAttempts = 0;
    private final Runnable reconnectRunnable = this::connectSse;

    private boolean hasMore = false;
    private Integer nextId = null;
    private boolean isLoadingMore = false;
    private boolean isInitialLoading = false;

    private String currentUserId;

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_chat, container, false);

        rvMessages = view.findViewById(R.id.rvMessages);
        etMessage = view.findViewById(R.id.etMessage);
        btnSend = view.findViewById(R.id.btnSend);
        viewConnectionStatus = view.findViewById(R.id.viewConnectionStatus);
        tvConnectionStatus = view.findViewById(R.id.tvConnectionStatus);
        pbLoadingMore = view.findViewById(R.id.pbLoadingMore);

        TokenManager tokenManager = TokenManager.getInstance(requireContext());
        currentUserId = tokenManager.getUserId();

        chatRepository = new ChatRepository(requireContext());
        chatAdapter = new ChatAdapter(currentUserId);

        LinearLayoutManager layoutManager = new LinearLayoutManager(getContext());
        layoutManager.setStackFromEnd(true);
        rvMessages.setLayoutManager(layoutManager);
        rvMessages.setAdapter(chatAdapter);

        gson = new Gson();
        okHttpClient = new OkHttpClient.Builder()
                .readTimeout(0, TimeUnit.MILLISECONDS)
                .build();

        // Pagination scroll listener
        rvMessages.addOnScrollListener(new RecyclerView.OnScrollListener() {
            @Override
            public void onScrolled(@NonNull RecyclerView recyclerView, int dx, int dy) {
                super.onScrolled(recyclerView, dx, dy);
                if (layoutManager.findFirstVisibleItemPosition() == 0) {
                    if (hasMore && !isLoadingMore && nextId != null) {
                        loadOlderMessages();
                    }
                }
            }
        });

        btnSend.setOnClickListener(v -> sendMessage());

        if (tokenManager.hasToken()) {
            loadInitialHistory();
            connectSse();
        }

        return view;
    }

    @Override
    public void onViewCreated(@NonNull View view, @Nullable Bundle savedInstanceState) {
        super.onViewCreated(view, savedInstanceState);
        TokenManager tokenManager = TokenManager.getInstance(requireContext());
        if (!tokenManager.hasToken()) {
            Toast.makeText(getContext(), "Vui lòng đăng nhập để sử dụng chat cộng đồng", Toast.LENGTH_SHORT).show();
            Navigation.findNavController(view).navigate(R.id.LoginFragment);
        }
    }

    private void loadInitialHistory() {
        if (isInitialLoading) return;
        isInitialLoading = true;
        chatRepository.getChatHistory(null, 20, new BaseCallback<ApiResponse<ChatHistoryResponse>>() {
            @Override
            public void onSuccess(ApiResponse<ChatHistoryResponse> response) {
                isInitialLoading = false;
                if (!isAdded()) return;
                if (response != null && response.getData() != null) {
                    ChatHistoryResponse data = response.getData();
                    hasMore = data.isHasMore();
                    nextId = data.getNextId();

                    List<ChatMessageResponse> messages = data.getMessages();
                    // History returns newest first, so we reverse it to display oldest at the top, newest at the bottom
                    if (messages != null) {
                        Collections.reverse(messages);
                        chatAdapter.setMessages(messages);
                    }
                }
            }

            @Override
            public void onError(String message) {
                isInitialLoading = false;
                if (!isAdded()) return;
                Toast.makeText(getContext(), "Không thể tải lịch sử chat: " + message, Toast.LENGTH_SHORT).show();
            }
        });
    }

    private void loadOlderMessages() {
        if (isLoadingMore || nextId == null) return;
        isLoadingMore = true;
        pbLoadingMore.setVisibility(View.VISIBLE);

        LinearLayoutManager layoutManager = (LinearLayoutManager) rvMessages.getLayoutManager();
        int firstVisiblePosition = layoutManager != null ? layoutManager.findFirstVisibleItemPosition() : 0;
        View firstVisibleView = layoutManager != null ? layoutManager.findViewByPosition(firstVisiblePosition) : null;
        int topOffset = firstVisibleView != null ? firstVisibleView.getTop() : 0;

        chatRepository.getChatHistory(nextId, 20, new BaseCallback<ApiResponse<ChatHistoryResponse>>() {
            @Override
            public void onSuccess(ApiResponse<ChatHistoryResponse> response) {
                if (!isAdded()) return;
                isLoadingMore = false;
                pbLoadingMore.setVisibility(View.GONE);

                if (response != null && response.getData() != null) {
                    ChatHistoryResponse data = response.getData();
                    hasMore = data.isHasMore();
                    nextId = data.getNextId();

                    List<ChatMessageResponse> messages = data.getMessages();
                    if (messages != null && !messages.isEmpty()) {
                        // Reverse older messages so they maintain correct chronological order when prepended
                        Collections.reverse(messages);
                        chatAdapter.prependMessages(messages);

                        // Preserve scroll position
                        if (layoutManager != null) {
                            layoutManager.scrollToPositionWithOffset(firstVisiblePosition + messages.size(), topOffset);
                        }
                    }
                }
            }

            @Override
            public void onError(String message) {
                if (!isAdded()) return;
                isLoadingMore = false;
                pbLoadingMore.setVisibility(View.GONE);
                Toast.makeText(getContext(), "Không thể tải tin nhắn cũ: " + message, Toast.LENGTH_SHORT).show();
            }
        });
    }

    private void connectSse() {
        updateStatusUI(ConnectionStatus.CONNECTING);
        new Thread(() -> {
            String token = ClerkAuthBridge.getTokenBlocking();
            if (token == null || token.isEmpty()) {
                mainHandler.post(() -> {
                    if (!isAdded()) return;
                    updateStatusUI(ConnectionStatus.ERROR);
                    scheduleReconnect();
                });
                return;
            }
            mainHandler.post(() -> {
                if (!isAdded()) return;
                startSse(token);
            });
        }).start();
    }

    private void startSse(String token) {
        if (eventSource != null) {
            eventSource.cancel();
        }

        String rawBaseUrl = Constants.BASE_URL;
        if (rawBaseUrl.endsWith("/")) {
            rawBaseUrl = rawBaseUrl.substring(0, rawBaseUrl.length() - 1);
        }
        String sseUrl = rawBaseUrl + "/chat/events?stream=messages&token=" + Uri.encode(token);

        Request request = new Request.Builder()
                .url(sseUrl)
                .header("Accept", "text/event-stream")
                .build();

        EventSource.Factory factory = EventSources.createFactory(okHttpClient);
        eventSource = factory.newEventSource(request, new EventSourceListener() {
            @Override
            public void onOpen(@NonNull EventSource eventSource, @NonNull Response response) {
                reconnectAttempts = 0;
                mainHandler.post(() -> {
                    if (!isAdded()) return;
                    updateStatusUI(ConnectionStatus.CONNECTED);
                });
            }

            @Override
            public void onEvent(@NonNull EventSource eventSource, @Nullable String id, @Nullable String type, @NonNull String data) {
                if ("chat.message.created".equals(type)) {
                    try {
                        ChatMessageResponse msg = gson.fromJson(data, ChatMessageResponse.class);
                        mainHandler.post(() -> {
                            if (!isAdded()) return;
                            chatAdapter.addMessage(msg);
                            rvMessages.smoothScrollToPosition(chatAdapter.getItemCount() - 1);
                        });
                    } catch (Exception e) {
                        e.printStackTrace();
                    }
                }
            }

            @Override
            public void onClosed(@NonNull EventSource eventSource) {
                mainHandler.post(() -> {
                    if (!isAdded()) return;
                    updateStatusUI(ConnectionStatus.DISCONNECTED);
                });
            }

            @Override
            public void onFailure(@NonNull EventSource eventSource, @Nullable Throwable t, @Nullable Response response) {
                mainHandler.post(() -> {
                    if (!isAdded()) return;
                    updateStatusUI(ConnectionStatus.ERROR);
                    scheduleReconnect();
                });
            }
        });
    }

    private void scheduleReconnect() {
        mainHandler.removeCallbacks(reconnectRunnable);
        reconnectAttempts++;
        long delay = Math.min(30000L, 1000L * (1L << Math.min(reconnectAttempts, 5)));
        mainHandler.postDelayed(reconnectRunnable, delay);
    }

    private void updateStatusUI(ConnectionStatus status) {
        if (!isAdded()) return;
        int color;
        String statusText;
        switch (status) {
            case CONNECTED:
                color = 0xFF34A853; // Green
                statusText = "Đang trực tuyến";
                break;
            case CONNECTING:
                color = 0xFFFBBC05; // Yellow
                statusText = "Đang kết nối...";
                break;
            case DISCONNECTED:
            case ERROR:
            default:
                color = 0xFFEA4335; // Red
                statusText = "Mất kết nối";
                break;
        }
        tvConnectionStatus.setText(statusText);
        if (viewConnectionStatus.getBackground() != null) {
            viewConnectionStatus.getBackground().setTint(color);
        }
    }

    private void sendMessage() {
        String text = etMessage.getText().toString().trim();
        if (text.isEmpty()) return;

        btnSend.setEnabled(false);
        chatRepository.sendMessage(text, new BaseCallback<ApiResponse<ChatMessageResponse>>() {
            @Override
            public void onSuccess(ApiResponse<ChatMessageResponse> response) {
                if (!isAdded()) return;
                btnSend.setEnabled(true);
                etMessage.setText("");
                if (response != null && response.getData() != null) {
                    chatAdapter.addMessage(response.getData());
                    rvMessages.smoothScrollToPosition(chatAdapter.getItemCount() - 1);
                }
            }

            @Override
            public void onError(String message) {
                if (!isAdded()) return;
                btnSend.setEnabled(true);
                Toast.makeText(getContext(), "Không thể gửi tin nhắn: " + message, Toast.LENGTH_SHORT).show();
            }
        });
    }

    @Override
    public void onDestroy() {
        super.onDestroy();
        mainHandler.removeCallbacks(reconnectRunnable);
        if (eventSource != null) {
            eventSource.cancel();
        }
    }
}
