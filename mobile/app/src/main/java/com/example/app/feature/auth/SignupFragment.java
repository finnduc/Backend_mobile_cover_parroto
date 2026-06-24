package com.example.app.feature.auth;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.EditText;
import android.widget.TextView;
import android.widget.Toast;

import androidx.annotation.Nullable;
import androidx.appcompat.app.AlertDialog;
import androidx.fragment.app.Fragment;
import androidx.navigation.Navigation;

import com.example.app.R;
import com.example.app.data.repository.AuthRepository;

public class SignupFragment extends Fragment {
    private AuthRepository authRepository;

    @Override
    @Nullable
    public View onCreateView(LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_signup, container, false);
        TextView signup = view.findViewById(R.id.btnRegister);
        EditText getFullname = view.findViewById(R.id.getFullName);
        EditText getUsername = view.findViewById(R.id.getUsername);
        EditText getPassword = view.findViewById(R.id.getPassword);
        EditText getConfirmPassword = view.findViewById(R.id.getConfirmPassword);
        authRepository = new AuthRepository(requireContext());

        signup.setOnClickListener(v -> {
            String fullname = getFullname.getText().toString().trim();
            String username = getUsername.getText().toString().trim();
            String password = getPassword.getText().toString().trim();
            String confirmPassword = getConfirmPassword.getText().toString().trim();

            boolean isValid = true;
            if (username.isEmpty()) {
                getUsername.setError("Vui long nhap Email");
                isValid = false;
            }
            if (password.isEmpty()) {
                getPassword.setError("Vui long nhap mat khau");
                isValid = false;
            }
            if (!android.util.Patterns.EMAIL_ADDRESS.matcher(username).matches()) {
                getUsername.setError("Email khong hop le");
                isValid = false;
            }
            if (password.length() <= 5) {
                getPassword.setError("Mat khau phai co it nhat 6 ky tu");
                isValid = false;
            }
            if (!password.equals(confirmPassword)) {
                getConfirmPassword.setError("Mat khau khong khop");
                isValid = false;
            }

            if (isValid) {
                register(username, password, fullname);
            }
        });

        setupLoginNavigation(view);
        return view;
    }

    private void register(String email, String password, String fullname) {
        authRepository.Register(email, password, fullname, new AuthRepository.authCallBack<String>() {
            @Override
            public void onSuccess(String data) {
                if (isAdded() && getView() != null && getContext() != null) {
                    Toast.makeText(getContext(), "Dang ky thanh cong!", Toast.LENGTH_SHORT).show();
                    Navigation.findNavController(getView()).navigate(R.id.action_signupFragment_to_StudyFragment);
                }
            }

            @Override
            public void onError(String message) {
                if (isAdded() && getContext() != null) {
                    Toast.makeText(getContext(), "Loi dang ki " + message, Toast.LENGTH_SHORT).show();
                }
            }

            @Override
            public void onNeedsVerification() {
                if (isAdded() && getContext() != null) {
                    showVerificationDialog();
                }
            }
        });
    }

    private void setupLoginNavigation(View view) {
        view.findViewById(R.id.tvLogin).setOnClickListener(v -> {
            Navigation.findNavController(v).navigate(R.id.action_signupFragment_to_loginFragment);
        });
    }

    private void showVerificationDialog() {
        EditText input = new EditText(requireContext());
        input.setHint("Ma xac minh email");
        new AlertDialog.Builder(requireContext())
                .setTitle("Xac minh email")
                .setMessage("Nhap ma xac minh Clerk da gui toi email cua ban.")
                .setView(input)
                .setPositiveButton("Xac minh", (dialog, which) -> {
                    String code = input.getText().toString().trim();
                    if (code.isEmpty()) {
                        Toast.makeText(requireContext(), "Vui long nhap ma xac minh", Toast.LENGTH_SHORT).show();
                        showVerificationDialog();
                        return;
                    }
                    verifyRegistration(code);
                })
                .setNegativeButton("Huy", null)
                .show();
    }

    private void verifyRegistration(String code) {
        authRepository.verifyRegistration(code, new AuthRepository.authCallBack<String>() {
            @Override
            public void onSuccess(String data) {
                if (isAdded() && getView() != null && getContext() != null) {
                    Toast.makeText(getContext(), "Dang ky thanh cong!", Toast.LENGTH_SHORT).show();
                    Navigation.findNavController(getView()).navigate(R.id.action_signupFragment_to_StudyFragment);
                }
            }

            @Override
            public void onError(String message) {
                if (isAdded() && getContext() != null) {
                    Toast.makeText(getContext(), "Loi xac minh " + message, Toast.LENGTH_SHORT).show();
                }
            }
        });
    }
}
