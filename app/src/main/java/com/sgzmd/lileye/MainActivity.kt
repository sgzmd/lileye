package com.sgzmd.lileye

import android.content.Intent
import android.os.Bundle
import android.provider.Settings
import android.util.Log
import android.widget.Button
import android.widget.Toast
import androidx.appcompat.app.AppCompatActivity
import androidx.core.app.NotificationManagerCompat
import com.sgzmd.lileye.service.NotificationListener
import androidx.core.content.edit

private const val NOTIFICATION_LISTENER_PREFS = "NotificationListenerPrefs"
private const val IS_NOTIFICATION_LISTENER_ENABLED = "isNotificationListenerEnabled"

class MainActivity : AppCompatActivity() {
    private lateinit var toggleButton: Button

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        toggleButton = findViewById(R.id.toggleButton)
        updateButtonState()
        toggleButton.setOnClickListener {
            if (isNotificationServiceEnabled()) {
                stopNotificationListener()
            } else {
                requestNotificationPermission()
                startNotificationListener()
            }
            updateButtonState()
        }
    }

    private fun hasNotificationPermission() =
        Settings.Secure.getString(contentResolver, "enabled_notification_listeners")
            ?.contains(packageName) == true

    private fun isNotificationServiceRunning() =
        NotificationManagerCompat.getEnabledListenerPackages(this).contains(packageName)

    private fun isNotificationServiceEnabled() =
        getSharedPreferences(NOTIFICATION_LISTENER_PREFS, MODE_PRIVATE)
            .getBoolean(IS_NOTIFICATION_LISTENER_ENABLED, false)

    private fun startNotificationListener() {
        startService(Intent(this, NotificationListener::class.java))
        updateNotificationListenerSetting(true)
        Toast.makeText(this, "Notification listener started", Toast.LENGTH_SHORT).show()
        Log.d("MainActivity", "Notification listener started successfully")
    }

    private fun stopNotificationListener() {
        stopService(Intent(this, NotificationListener::class.java))
        updateNotificationListenerSetting(false)
        Toast.makeText(this, "Notification listener stopped", Toast.LENGTH_SHORT).show()
        Log.d("MainActivity", "Notification listener stopped successfully")
    }

    private fun updateNotificationListenerSetting(enabled: Boolean) {
        getSharedPreferences(NOTIFICATION_LISTENER_PREFS, MODE_PRIVATE)
            .edit() {
                putBoolean(IS_NOTIFICATION_LISTENER_ENABLED, enabled)
            }
        Log.d("MainActivity", "Notification listener setting updated to $enabled")
    }

    private fun requestNotificationPermission() {
        if (!hasNotificationPermission()) {
            startActivity(Intent(Settings.ACTION_NOTIFICATION_LISTENER_SETTINGS))
            Toast.makeText(
                this,
                "Please enable notification access for LilEye",
                Toast.LENGTH_LONG
            ).show()
        }
    }

    private fun updateButtonState() {
        toggleButton.text =
            if (isNotificationServiceRunning() && isNotificationServiceEnabled()) "Stop Listening" else "Start Listening"
    }

    override fun onResume() {
        super.onResume()
        updateButtonState()
    }
}