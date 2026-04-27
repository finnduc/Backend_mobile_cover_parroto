package com.example.app.data.remote.model.request;
import com.google.gson.annotations.SerializedName;
public class RegisterRequest {
    private String email;
    private String password;
    private Boolean returnSecureToken;

    public RegisterRequest(String email, String password, String returnSecureToken) {
        this.email = email;
        this.password = password;
        this.returnSecureToken = true;
    }
}
