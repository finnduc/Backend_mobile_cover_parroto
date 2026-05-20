package com.example.app.data.remote.model.response.clerk;

import com.google.gson.annotations.SerializedName;
import java.util.List;

public class ClerkClient {
    @SerializedName("sessions")
    private List<ClerkSession> sessions;

    public List<ClerkSession> getSessions() { return sessions; }

    public String getJwt() {
        if (sessions != null && !sessions.isEmpty()) {
            ClerkSession session = sessions.get(0);
            if (session.getLastActiveToken() != null) {
                return session.getLastActiveToken().getJwt();
            }
        }
        return null;
    }

    public String getSessionId() {
        if (sessions != null && !sessions.isEmpty()) {
            return sessions.get(0).getId();
        }
        return null;
    }
}
