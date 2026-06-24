package com.example.app;

import static org.junit.Assert.assertTrue;

import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.charset.StandardCharsets;

import org.junit.Test;

public class SignupNavigationTest {
    @Test
    public void successfulSignupNavigatesToStudyInsteadOfLogin() throws Exception {
        String signupFragment = readFile(
                "src/main/java/com/example/app/feature/auth/SignupFragment.java");
        String navGraph = readFile("src/main/res/navigation/nav_graph.xml");

        assertTrue(
                "Signup success should navigate with the Signup -> Study action.",
                signupFragment.contains("R.id.action_signupFragment_to_StudyFragment"));
        assertTrue(
                "Navigation graph should define a Signup -> Study action.",
                navGraph.contains("action_signupFragment_to_StudyFragment")
                        && navGraph.contains("app:destination=\"@id/StudyFragment\""));
    }

    private String readFile(String path) throws Exception {
        return new String(Files.readAllBytes(Path.of(path)), StandardCharsets.UTF_8);
    }
}
