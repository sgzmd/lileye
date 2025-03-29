package com.sgzmd.lileye.service

import android.app.Notification
import android.content.Context
import android.os.Bundle
import android.os.UserHandle
import android.service.notification.StatusBarNotification
import androidx.test.core.app.ApplicationProvider
import com.sgzmd.lileye.model.Notification
import com.sgzmd.lileye.queue.MessageQueue
import org.junit.Assert.assertEquals
import org.junit.Before
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RobolectricTestRunner
import java.lang.reflect.Field

@RunWith(RobolectricTestRunner::class)
class NotificationListenerTest {

    private lateinit var notificationListener: NotificationListener
    private lateinit var testMessageQueue: MessageQueue

    @Before
    fun setUp() {
        // Use a test-friendly MessageQueue that allows inspection of messages.
        testMessageQueue = MessageQueue()
        notificationListener = NotificationListener()
        // For testing purposes we use reflection to set the private messageQueue field
        val field: Field = NotificationListener::class.java.getDeclaredField("messageQueue")
        field.isAccessible = true
        field.set(notificationListener, testMessageQueue)
    }

    @Test
    fun testOnNotificationPosted_validNotification() {
        val context = ApplicationProvider.getApplicationContext<Context>()

        // Prepare a Bundle with notification extras.
        val extras = Bundle().apply {
            putString("android.title", "Real Title")
            putString("android.text", "Real Text")
        }

        // Build a real Notification.
        val notification = Notification.Builder(context)
            .setContentTitle("Real Title")
            .setContentText("Real Text")
            .setExtras(extras)
            .build()

        val userHandle = UserHandle.getUserHandleForUid(1000)
        val sbn = StatusBarNotification(
            "com.example.test",
            "com.example.test",
            1,
            null,
            1000,
            1000,
            0,
            notification,
            userHandle,
            System.currentTimeMillis()
        )

        notificationListener.onNotificationPosted(sbn)

        val messages: List<com.sgzmd.lileye.model.Notification> = testMessageQueue.getMessages()
        assertEquals(1, messages.size)
        val message: com.sgzmd.lileye.model.Notification = messages[0]
        assertEquals("com.example.test", message.packageName)
        assertEquals("Real Title", message.title)
        assertEquals("Real Text", message.text)
        // Optionally, assert timestamp and extras.
    }
}