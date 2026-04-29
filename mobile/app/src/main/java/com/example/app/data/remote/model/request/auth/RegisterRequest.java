package com.example.app.data.remote.model.request.auth;

public class RegisterRequest {
    private String email;
    private String password;
    private Boolean returnSecureToken;

    public RegisterRequest(String email, String password, Boolean returnSecureToken) {
        this.email = email;
        this.password = password;
        this.returnSecureToken = returnSecureToken;
    }
}
