package com.sgzmd.lileye.service

import android.app.Notification
import android.os.Bundle
import android.service.notification.StatusBarNotification
import com.sgzmd.lileye.model.Message
import com.sgzmd.lileye.queue.MessageQueue
import kotlinx.coroutines.ExperimentalCoroutinesApi
import kotlinx.coroutines.test.TestScope
import kotlinx.coroutines.test.runTest
import org.junit.Before
import org.junit.Test
import org.junit.runner.RunWith
import org.mockito.Mock
import org.mockito.Mockito.verify
import org.mockito.Mockito.`when`
import org.robolectric.RobolectricTestRunner
import org.robolectric.annotation.Config
import java.util.UUID
import kotlin.test.assertEquals
import kotlin.test.assertNotNull

@OptIn(ExperimentalCoroutinesApi::class)
@RunWith(RobolectricTestRunner::class)
@Config(sdk = [29])
class NotificationListenerTest {

    @Mock
    private lateinit var messageQueue: MessageQueue

    private lateinit var notificationListener: NotificationListener
    private val testScope = TestScope()

    @Before
    fun setup() {
        notificationListener = NotificationListener()
        notificationListener.messageQueue = messageQueue
    }

    @Test
    fun `onNotificationPosted creates correct Message object`() = testScope.runTest {
        // Given
        val packageName = "com.test.app"
        val title = "Test Title"
        val text = "Test Text"
        val timestamp = System.currentTimeMillis()
        
        val notification = createTestNotification(title, text)
        val sbn = createTestStatusBarNotification(packageName, notification, timestamp)

        // When
        notificationListener.onNotificationPosted(sbn)

        // Then
        verify(messageQueue).addMessage(
            Message(
                packageName = packageName,
                title = title,
                text = text,
                timestamp = timestamp,
                extras = mapOf(
                    "android.title" to title,
                    "android.text" to text
                )
            )
        )
    }

    @Test
    fun `onNotificationPosted handles null title and text`() = testScope.runTest {
        // Given
        val packageName = "com.test.app"
        val timestamp = System.currentTimeMillis()
        
        val notification = createTestNotification(null, null)
        val sbn = createTestStatusBarNotification(packageName, notification, timestamp)

        // When
        notificationListener.onNotificationPosted(sbn)

        // Then
        verify(messageQueue).addMessage(
            Message(
                packageName = packageName,
                title = null,
                text = null,
                timestamp = timestamp,
                extras = emptyMap()
            )
        )
    }

    @Test
    fun `onNotificationPosted handles empty extras`() = testScope.runTest {
        // Given
        val packageName = "com.test.app"
        val timestamp = System.currentTimeMillis()
        
        val notification = Notification.Builder(notificationListener, "test_channel")
            .build()
        val sbn = createTestStatusBarNotification(packageName, notification, timestamp)

        // When
        notificationListener.onNotificationPosted(sbn)

        // Then
        verify(messageQueue).addMessage(
            Message(
                packageName = packageName,
                title = null,
                text = null,
                timestamp = timestamp,
                extras = emptyMap()
            )
        )
    }

    @Test
    fun `onNotificationPosted handles exception gracefully`() = testScope.runTest {
        // Given
        val packageName = "com.test.app"
        val notification = Notification.Builder(notificationListener, "test_channel")
            .build()
//        val sbn = createTestStatusBarNotification(packageName, notification, System.currentTimeMillis())

        // When
        notificationListener.onNotificationPosted(sbn)

        // Then
        // Verify that the service didn't crash and still attempted to process the notification
        verify(messageQueue).addMessage(
            Message(
                packageName = packageName,
                title = null,
                text = null,
                timestamp = sbn.postTime,
                extras = emptyMap()
            )
        )
    }

    private fun createTestNotification(title: String?, text: String?): Notification {
        val extras = Bundle().apply {
            title?.let { putString("android.title", it) }
            text?.let { putString("android.text", it) }
        }
        
        return Notification.Builder(notificationListener, "test_channel")
            .setContentTitle(title)
            .setContentText(text)
            .setExtras(extras)
            .build()
    }

//    private fun createTestStatusBarNotification(
//        packageName: String,
//        notification: Notification,
//        timestamp: Long
//    ): StatusBarNotification {
//        return StatusBarNotification(
//            packageName,
//            packageName,
//            UUID.randomUUID().hashCode(), // Use a valid Int value for id
//            null,
//            0, // Use a valid Int value for userId
//            notification,
//            null,
//            null,
//            timestamp,
//            timestamp
//        )
//    }
}