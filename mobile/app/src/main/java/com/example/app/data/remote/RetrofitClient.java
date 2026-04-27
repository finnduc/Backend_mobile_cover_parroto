package com.example.app.data.remote;

import android.content.Context;

import com.example.app.data.local.TokenManager;
import com.example.app.data.remote.api.AuthApi;
import com.example.app.data.remote.interceptor.AuthInterceptor;
import com.example.app.utils.Constants;

import okhttp3.OkHttpClient;
import okhttp3.logging.HttpLoggingInterceptor;
import retrofit2.Retrofit;
import retrofit2.converter.gson.GsonConverterFactory;

import java.util.concurrent.TimeUnit;

public class RetrofitClient {

    private static RetrofitClient instance;
    private final Retrofit retrofit;

    private RetrofitClient(Context context) {
        // 1. Logging — in request/response ra Logcat khi debug
        HttpLoggingInterceptor logging = new HttpLoggingInterceptor();
        logging.setLevel(HttpLoggingInterceptor.Level.BODY);

        // 2. OkHttpClient — gắn interceptor vào
        OkHttpClient client = new OkHttpClient.Builder()
                .addInterceptor(new AuthInterceptor(TokenManager.getInstance(context)))
                .addInterceptor(logging)
                .connectTimeout(30, TimeUnit.SECONDS)
                .readTimeout(30, TimeUnit.SECONDS)
                .writeTimeout(30, TimeUnit.SECONDS)
                .build();

        // 3. Retrofit — cấu hình chính
        retrofit = new Retrofit.Builder()
                .baseUrl(Constants.BASE_URL)
                .client(client)
                .addConverterFactory(GsonConverterFactory.create())
                .build();
    }

    public static synchronized RetrofitClient getInstance(Context context) {
        if (instance == null) {
            instance = new RetrofitClient(context);
        }
        return instance;
    }

    // Trả về các Api interface
    public AuthApi getAuthApi() {
        return retrofit.create(AuthApi.class);
    }
}