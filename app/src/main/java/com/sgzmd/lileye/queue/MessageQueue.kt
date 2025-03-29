package com.sgzmd.lileye.queue

import android.util.Log
import androidx.annotation.VisibleForTesting
import com.sgzmd.lileye.api.RetrofitClient
import com.sgzmd.lileye.model.Notification
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.Job
import kotlinx.coroutines.launch
import java.util.concurrent.ConcurrentLinkedQueue

class MessageQueue {
    private val queue = ConcurrentLinkedQueue<Notification>()
    private val scope = CoroutineScope(Dispatchers.IO + Job())
    private val TAG = "MessageQueue"

    fun addMessage(message: Notification) {
        queue.offer(message)
        processQueue()
    }

    private fun processQueue() {
        scope.launch {
            while (queue.isNotEmpty()) {
                val message = queue.poll()
                message?.let { processMessage(it) }
            }
        }
    }

    private suspend fun processMessage(message: Notification) {
        try {
            Log.d(TAG, "Sending notification: $message")
            val response = RetrofitClient.notificationApi.createNotification(message)
            if (response.isSuccessful) {
                Log.d(TAG, "Notification sent successfully: $message")
            } else {
                Log.e(
                    TAG,
                    "Failed to send notification. Status code: ${response.code()}, notification: $message"
                )
                // Consider retrying or requeuing here
            }
        } catch (e: Exception) {
            Log.e(TAG, "Error sending notification: $message", e)
            // Optionally handle retry logic
        }
        // Currently just logging, but can be extended to send to web service
        Log.d(TAG, "Processing message: $message")
    }

    @VisibleForTesting
    fun getMessages(): List<Notification> {
        return queue.toList()
    }
} 