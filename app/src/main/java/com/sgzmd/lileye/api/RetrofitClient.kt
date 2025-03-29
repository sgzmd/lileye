package com.sgzmd.lileye.api

import com.google.gson.GsonBuilder
import com.sgzmd.lileye.model.NotificationApi
import okhttp3.OkHttpClient
import okhttp3.logging.HttpLoggingInterceptor
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import java.time.Instant

object RetrofitClient {
    private const val BASE_URL = "http://10.0.2.2:8080" // Emulator-friendly localhost

    private val gson = GsonBuilder()
        .registerTypeAdapter(Instant::class.java, InstantConverter())
        .create()

    private val client = OkHttpClient.Builder()
        .addInterceptor(HttpLoggingInterceptor().apply {
            level = HttpLoggingInterceptor.Level.BODY
        })
        .build()

    val notificationApi: NotificationApi by lazy {
        Retrofit.Builder()
            .baseUrl(BASE_URL)
            .client(client)
            .addConverterFactory(GsonConverterFactory.create(gson))
            .build()
            .create(NotificationApi::class.java)
    }
}