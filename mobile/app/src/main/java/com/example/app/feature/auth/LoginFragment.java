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
import android.widget.Toast;
import com.example.app.R;
import com.example.app.data.remote.model.response.auth.SyncResponse;
import com.example.app.data.repository.AuthRepository;


public class LoginFragment extends Fragment {
    @Override
    @Nullable
     public View onCreateView(LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
         View view = inflater.inflate(R.layout.fragment_login, container, false);
         TextView Login = view.findViewById(R.id.btnLogin);
         EditText getUsername = view.findViewById(R.id.getUsername);
         EditText getPassword = view.findViewById(R.id.getPassword);
        AuthRepository authRepository = new AuthRepository(requireContext());

        Login.setOnClickListener(v ->
         {
             boolean isvalid = true;
             String Username = getUsername.getText().toString().trim();
             String Password = getPassword.getText().toString().trim();
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
             if (isvalid){
                 Login.setEnabled(false);
                 Login.setText("Đang đăng nhập...");
                 authRepository.login(Username,Password,
                         new AuthRepository.AuthCallback<SyncResponse>(){
                             @Override
                             public void onSuccess(SyncResponse data) {
                                 Login.setEnabled(true);
                                 Login.setText("Đăng nhập");
                                 Toast.makeText(requireContext(),"Đăng nhập thành công!",Toast.LENGTH_SHORT).show();
                                 Navigation.findNavController(v).navigate(R.id.action_LoginFragment_to_shadowingFragment);
                         }
                             @Override
                             public void onError(String message) {
                                Login.setEnabled(true);
                                Login.setText("ĐĂNG NHẬP");
                                Toast.makeText(requireContext(),message,Toast.LENGTH_LONG).show();
                             }
                         }
                     );

             }
         });
         register(view);
         return view;
    }

    private void register(View view){
        view.findViewById(R.id.tvRegister).setOnClickListener(v ->{
            Navigation.findNavController(v)
                    .navigate(R.id.action_LoginFragment_to_signupFragment);
                });
    }
}