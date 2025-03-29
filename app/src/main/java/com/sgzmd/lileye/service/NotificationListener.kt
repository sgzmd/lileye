package com.sgzmd.lileye.service

import android.provider.Settings
import android.service.notification.NotificationListenerService
import android.service.notification.StatusBarNotification
import android.util.Log
import com.sgzmd.lileye.model.Notification
import com.sgzmd.lileye.queue.MessageQueue
import java.time.Instant

class NotificationListener : NotificationListenerService() {
    var messageQueue: MessageQueue = MessageQueue()
        private set

    override fun onNotificationPosted(sbn: StatusBarNotification) {
        try {
            val notification = sbn.notification
            val extras = notification.extras

            val message = Notification(
                packageName = sbn.packageName,
                title = extras.getString("android.title")?.toString(),
                text = extras.getString("android.text")?.toString(),
                timestamp = Instant.ofEpochMilli(sbn.postTime),
                extras = extras.keySet().associateWith { extras.get(it)?.toString() ?: "" },
                deviceId = Settings.Secure.ANDROID_ID
            )

            messageQueue.addMessage(message)
            Log.d(TAG, "Received notification from ${sbn.packageName}")
        } catch (e: Exception) {
            Log.e(TAG, "Error processing notification", e)
            // Even if there's an error, try to create a basic message
            messageQueue.addMessage(
                Notification(
                    packageName = sbn.packageName,
                    title = null,
                    text = null,
                    timestamp = Instant.ofEpochMilli(sbn.postTime),
                    extras = emptyMap(),
                    deviceId = Settings.Secure.ANDROID_ID
                )
            )
        }
    }

    override fun onNotificationRemoved(sbn: StatusBarNotification) {
        // We don't need to handle notification removal
    }

    companion object {
        private const val TAG = "NotificationListener"
    }
} 