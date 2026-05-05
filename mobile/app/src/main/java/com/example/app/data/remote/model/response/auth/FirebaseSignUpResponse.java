package com.example.app.data.remote.model.response.auth;

public class FirebaseSignUpResponse {
    private String idToken ;
    private String email ;
    private String refreshToken ;
    private String localId ;

    public String getIdToken() {
        return idToken;
    }

    public String getEmail() {
        return email;
    }

    public String getRefreshToken() {
        return refreshToken;
    }

    public String getLocalId() {
        return localId;
    }
}
