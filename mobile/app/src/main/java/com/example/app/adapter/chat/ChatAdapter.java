package com.example.app.adapter.chat;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.bumptech.glide.Glide;
import com.example.app.R;
import com.example.app.data.remote.model.response.chat.ChatMessageResponse;
import de.hdodenhof.circleimageview.CircleImageView;

import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import java.util.Locale;
import java.util.TimeZone;

public class ChatAdapter extends RecyclerView.Adapter<RecyclerView.ViewHolder> {

    private static final int VIEW_TYPE_SENT = 1;
    private static final int VIEW_TYPE_RECEIVED = 2;

    private final List<ChatMessageResponse> messages = new ArrayList<>();
    private final String currentUserId;

    public ChatAdapter(String currentUserId) {
        this.currentUserId = currentUserId;
    }

    public void setMessages(List<ChatMessageResponse> newMessages) {
        messages.clear();
        if (newMessages != null) {
            messages.addAll(newMessages);
        }
        notifyDataSetChanged();
    }

    public void addMessage(ChatMessageResponse message) {
        for (ChatMessageResponse m : messages) {
            if (m.getId() == message.getId()) {
                return;
            }
        }
        messages.add(message);
        notifyItemInserted(messages.size() - 1);
    }

    public void prependMessages(List<ChatMessageResponse> oldMessages) {
        if (oldMessages == null || oldMessages.isEmpty()) return;
        List<ChatMessageResponse> toAdd = new ArrayList<>();
        for (ChatMessageResponse om : oldMessages) {
            boolean exists = false;
            for (ChatMessageResponse m : messages) {
                if (m.getId() == om.getId()) {
                    exists = true;
                    break;
                }
            }
            if (!exists) {
                toAdd.add(om);
            }
        }
        if (toAdd.isEmpty()) return;
        messages.addAll(0, toAdd);
        notifyItemRangeInserted(0, toAdd.size());
    }

    public List<ChatMessageResponse> getMessages() {
        return messages;
    }

    @Override
    public int getItemViewType(int position) {
        ChatMessageResponse msg = messages.get(position);
        if (msg.getUserId() != null && msg.getUserId().equals(currentUserId)) {
            return VIEW_TYPE_SENT;
        } else {
            return VIEW_TYPE_RECEIVED;
        }
    }

    @NonNull
    @Override
    public RecyclerView.ViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        if (viewType == VIEW_TYPE_SENT) {
            View view = LayoutInflater.from(parent.getContext())
                    .inflate(R.layout.item_chat_message_own, parent, false);
            return new SentMessageViewHolder(view);
        } else {
            View view = LayoutInflater.from(parent.getContext())
                    .inflate(R.layout.item_chat_message_other, parent, false);
            return new ReceivedMessageViewHolder(view);
        }
    }

    @Override
    public void onBindViewHolder(@NonNull RecyclerView.ViewHolder holder, int position) {
        ChatMessageResponse message = messages.get(position);
        String timeString = formatTime(message.getCreatedAt());

        if (holder instanceof SentMessageViewHolder) {
            SentMessageViewHolder sentHolder = (SentMessageViewHolder) holder;
            sentHolder.tvMessageContent.setText(message.getContent());
            sentHolder.tvTime.setText(timeString);
        } else if (holder instanceof ReceivedMessageViewHolder) {
            ReceivedMessageViewHolder receivedHolder = (ReceivedMessageViewHolder) holder;
            receivedHolder.tvMessageContent.setText(message.getContent());
            receivedHolder.tvTime.setText(timeString);
            
            String displayName = message.getUserName();
            if (displayName == null || displayName.trim().isEmpty()) {
                displayName = "Người dùng";
            }
            receivedHolder.tvUserName.setText(displayName);

            // Handle avatar
            String avatarUrl = message.getAvatarUrl();
            if (avatarUrl != null && !avatarUrl.isEmpty()) {
                receivedHolder.ivAvatar.setVisibility(View.VISIBLE);
                receivedHolder.tvAvatarFallback.setVisibility(View.GONE);
                Glide.with(receivedHolder.itemView.getContext())
                        .load(avatarUrl)
                        .placeholder(R.drawable.ic_avatar_placeholder)
                        .error(R.drawable.ic_avatar_placeholder)
                        .into(receivedHolder.ivAvatar);
            } else {
                receivedHolder.ivAvatar.setVisibility(View.GONE);
                receivedHolder.tvAvatarFallback.setVisibility(View.VISIBLE);
                receivedHolder.tvAvatarFallback.setText(getFallbackInitial(displayName, message.getUserId()));
            }
        }
    }

    @Override
    public int getItemCount() {
        return messages.size();
    }

    private String getFallbackInitial(String name, String userId) {
        String source = name != null && !name.trim().isEmpty() ? name.trim() : userId;
        if (source == null || source.isEmpty()) return "?";
        return source.substring(0, 1).toUpperCase();
    }

    private String formatTime(String isoString) {
        if (isoString == null || isoString.isEmpty()) return "";
        try {
            String formatted = isoString.replaceAll("Z$", "+0000");
            formatted = formatted.replaceAll("\\.\\d+(?=[+-])", "");
            SimpleDateFormat parser;
            if (formatted.contains("+")) {
                parser = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ssZ", Locale.getDefault());
            } else {
                parser = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss", Locale.getDefault());
            }
            parser.setTimeZone(TimeZone.getTimeZone("UTC"));
            Date date = parser.parse(formatted);
            if (date != null) {
                SimpleDateFormat formatter = new SimpleDateFormat("HH:mm", Locale.getDefault());
                formatter.setTimeZone(TimeZone.getDefault());
                return formatter.format(date);
            }
        } catch (Exception e) {
            try {
                if (isoString.length() >= 16) {
                    return isoString.substring(11, 16);
                }
            } catch (Exception ignored) {}
        }
        return "";
    }

    static class SentMessageViewHolder extends RecyclerView.ViewHolder {
        TextView tvMessageContent;
        TextView tvTime;

        SentMessageViewHolder(View itemView) {
            super(itemView);
            tvMessageContent = itemView.findViewById(R.id.tvMessageContent);
            tvTime = itemView.findViewById(R.id.tvTime);
        }
    }

    static class ReceivedMessageViewHolder extends RecyclerView.ViewHolder {
        CircleImageView ivAvatar;
        TextView tvAvatarFallback;
        TextView tvUserName;
        TextView tvMessageContent;
        TextView tvTime;

        ReceivedMessageViewHolder(View itemView) {
            super(itemView);
            ivAvatar = itemView.findViewById(R.id.ivAvatar);
            tvAvatarFallback = itemView.findViewById(R.id.tvAvatarFallback);
            tvUserName = itemView.findViewById(R.id.tvUserName);
            tvMessageContent = itemView.findViewById(R.id.tvMessageContent);
            tvTime = itemView.findViewById(R.id.tvTime);
        }
    }
}
