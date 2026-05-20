package com.example.app.data.remote;

import com.example.app.data.remote.api.ClerkApi;
import com.example.app.utils.Constants;

import okhttp3.OkHttpClient;
import okhttp3.logging.HttpLoggingInterceptor;
import retrofit2.Retrofit;
import retrofit2.converter.gson.GsonConverterFactory;

import java.util.concurrent.TimeUnit;

public class ClerkRetrofitClient {

    private static ClerkRetrofitClient instance;
    private final Retrofit retrofit;

    private ClerkRetrofitClient() {
        HttpLoggingInterceptor logging = new HttpLoggingInterceptor();
        logging.setLevel(HttpLoggingInterceptor.Level.BODY);

        OkHttpClient client = new OkHttpClient.Builder()
                .addInterceptor(logging)
                .connectTimeout(30, TimeUnit.SECONDS)
                .readTimeout(30, TimeUnit.SECONDS)
                .build();

        retrofit = new Retrofit.Builder()
                .baseUrl(Constants.CLERK_FRONTEND_API_URL + "/")
                .client(client)
                .addConverterFactory(GsonConverterFactory.create())
                .build();
    }

    public static synchronized ClerkRetrofitClient getInstance() {
        if (instance == null) {
            instance = new ClerkRetrofitClient();
        }
        return instance;
    }

    public ClerkApi getClerkApi() {
        return retrofit.create(ClerkApi.class);
    }
}
