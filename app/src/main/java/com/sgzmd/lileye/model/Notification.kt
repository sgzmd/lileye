package com.sgzmd.lileye.model

import android.os.Parcelable
import kotlinx.parcelize.Parcelize
import retrofit2.Response
import retrofit2.http.Body
import retrofit2.http.POST
import java.time.Instant

@Parcelize
data class Notification(
    val packageName: String,
    val title: String?,
    val text: String?,
    val timestamp: Instant,
    val extras: Map<String, String> = emptyMap(),
    val deviceId: String,
) : Parcelable

interface NotificationApi {
    @POST("/api/notifications")
    suspend fun createNotification(@Body notification: Notification): Response<Void>
}