package com.example.app.feature.auth;

import android.os.Bundle;

import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.navigation.Navigation;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Button;
import android.widget.EditText;
import android.widget.TextView;

import com.example.app.R;

public class SignupFragment extends Fragment {
    @Override
    @Nullable
    public View onCreateView(LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_signup, container, false);
        TextView Signup = view.findViewById(R.id.btnRegister);
        EditText getFullname = view.findViewById(R.id.getFullName);
        EditText getUsername = view.findViewById(R.id.getUsername);
        EditText getPassword = view.findViewById(R.id.getPassword);
        EditText getConfirmPassword = view.findViewById(R.id.getConfirmPassword);

        Signup.setOnClickListener(v -> {
            String Fullname , Username, Password, Confirmpassword;
            Fullname = getFullname.getText().toString().trim();
            Username = getUsername.getText().toString().trim();
            Password = getPassword.getText().toString().trim();
            Confirmpassword = getConfirmPassword.getText().toString().trim();

            boolean isvalid = true;

            if (Username.isEmpty()) {
                getUsername.setError("Vui lòng nhập Email");
                isvalid = false;
            }

            if (Password.isEmpty()) {
                getPassword.setError("Vui lòng nhập mật khẩu");
                isvalid = false;
            }

            if (!android.util.Patterns.EMAIL_ADDRESS.matcher(Username).matches()){
                getUsername.setError("Email không hợp lệ");
                isvalid = false;
            }

            if (Password.length()<=5){
                getPassword.setError("Mật khẩu phải có ít nhất 6 ký tự");
                isvalid = false;
            }

            if (!Password.equals(Confirmpassword)){
                getConfirmPassword.setError("Mật khẩu không khớp");
                isvalid = false;
            }

            if (isvalid){

            }

                }
            )
        ;


        Login(view);
        return view;
    }

    private void Login(View view) {
        view.findViewById(R.id.tvLogin).setOnClickListener(v -> {
            Navigation.findNavController(v)
                    .navigate(R.id.action_signupFragment_to_loginFragment);
        });
    }

}