package com.example.app.feature.auth;

import android.os.Bundle;

import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.navigation.Navigation;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import com.example.app.R;


public class LoginFragment extends Fragment {
    @Override
    @Nullable
     public View onCreateView(LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
         View view = inflater.inflate(R.layout.fragment_login, container, false);
         register(view);
         return view;
    }

    private void register(View view){
        view.findViewById(R.id.tvRegister).setOnClickListener(v ->{
            Navigation.findNavController(v)
                    .navigate(R.id.action_LoginFragment_to_signupFragment);
                }
                );
    }

}