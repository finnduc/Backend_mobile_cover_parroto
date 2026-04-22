package com.example.app.feature.auth;

import android.os.Bundle;

import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.navigation.Navigation;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import com.example.app.R;

public class SignupFragment extends Fragment {
    @Override
    @Nullable
    public View onCreateView(LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_signup, container, false);
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