package com.example.app.feature.setting.ui;

import androidx.appcompat.app.AlertDialog;
import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.navigation.Navigation;

import com.example.app.R;

public class SettingsFragment extends Fragment {

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater,
                             @Nullable ViewGroup container,
                             @Nullable Bundle savedInstanceState) {

        View view = inflater.inflate(R.layout.fragment_settings, container, false);

        setupProfileCard(view);
        setupMenuRows(view);
        setupLogout(view);

        return view;
    }

    private void setupProfileCard(View view) {
        view.findViewById(R.id.cardLogin).setOnClickListener(v -> {
            Navigation.findNavController(v)
                    .navigate(R.id.action_settingsFragment_to_loginFragment);
        });
    }

    // ──────────────────────────────────────────
    // Menu Rows
    // ──────────────────────────────────────────
    private void setupMenuRows(View view) {

        view.findViewById(R.id.rowNotes).setOnClickListener(v -> {
            // TODO: navigate sang NotesFragment
        });

        view.findViewById(R.id.rowProgress).setOnClickListener(v -> {
            // TODO: navigate sang ProgressFragment
        });

        view.findViewById(R.id.rowLeaderboard).setOnClickListener(v -> {
            // TODO: navigate sang LeaderboardFragment
        });
    }

    private void setupLogout(View view) {
        view.findViewById(R.id.cardLogout).setOnClickListener(v -> {
            showLogoutDialog();
        });
    }

    private void showLogoutDialog() {
        new AlertDialog.Builder(requireActivity())
                .setTitle("Đăng xuất")
                .setMessage("Bạn có chắc muốn đăng xuất không?")
                .setPositiveButton("Đăng xuất", (dialog, which) -> {
                    performLogout();
                })
                .setNegativeButton("Hủy", null)
                .show();
    }

    private void performLogout() {
        // Xóa session (SharedPreferences)
        requireContext()
                .getSharedPreferences("user_session", 0)
                .edit()
                .clear()
                .apply();

        // TODO: navigate về LoginFragment
        // Navigation.findNavController(requireView())
        //     .navigate(R.id.action_settingsFragment_to_loginFragment);
    }
}