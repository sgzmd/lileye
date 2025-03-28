package com.sgzmd.lileye.service

import android.service.notification.NotificationListenerService
import android.service.notification.StatusBarNotification
import android.util.Log
import com.sgzmd.lileye.model.Message
import com.sgzmd.lileye.queue.MessageQueue

class NotificationListener : NotificationListenerService() {
    var messageQueue: MessageQueue = MessageQueue()
        private set

    override fun onNotificationPosted(sbn: StatusBarNotification) {
        try {
            val notification = sbn.notification
            val extras = notification.extras

            val message = Message(
                packageName = sbn.packageName,
                title = extras.getString("android.title")?.toString(),
                text = extras.getString("android.text")?.toString(),
                timestamp = sbn.postTime,
                extras = extras.keySet().associateWith { extras.get(it)?.toString() ?: "" }
            )

            messageQueue.addMessage(message)
            Log.d(TAG, "Received notification from ${sbn.packageName}")
        } catch (e: Exception) {
            Log.e(TAG, "Error processing notification", e)
            // Even if there's an error, try to create a basic message
            messageQueue.addMessage(
                Message(
                    packageName = sbn.packageName,
                    title = null,
                    text = null,
                    timestamp = sbn.postTime,
                    extras = emptyMap()
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