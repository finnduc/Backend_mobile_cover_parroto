package com.example.app.data.remote.model.request.auth;

public class TokenRequest {

    private String email;
    private String password;

    public TokenRequest(String email, String password) {
        this.email = email;
        this.password = password;
    }

}
