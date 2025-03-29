package com.sgzmd.lileye.queue

import android.util.Log
import androidx.annotation.VisibleForTesting
import com.sgzmd.lileye.model.Message
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.Job
import kotlinx.coroutines.launch
import java.util.concurrent.ConcurrentLinkedQueue

class MessageQueue {
    private val queue = ConcurrentLinkedQueue<Message>()
    private val scope = CoroutineScope(Dispatchers.IO + Job())
    private val TAG = "MessageQueue"

    fun addMessage(message: Message) {
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

    private fun processMessage(message: Message) {
        // Currently just logging, but can be extended to send to web service
        Log.d(TAG, "Processing message: $message")
    }

    @VisibleForTesting fun getMessages(): List<Message> {
        return queue.toList()
    }
} 