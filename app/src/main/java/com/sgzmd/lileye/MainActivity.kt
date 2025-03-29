package com.sgzmd.lileye

import android.content.Intent
import android.os.Bundle
import android.provider.Settings
import android.widget.Button
import android.widget.Toast
import androidx.appcompat.app.AppCompatActivity
import com.sgzmd.lileye.service.NotificationListener

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
        }
    }

    private fun hasNotificationPermission(): Boolean {
        val pkgName = packageName
        val flat = Settings.Secure.getString(contentResolver, "enabled_notification_listeners")
        return flat?.contains(pkgName) == true
    }

    private fun isNotificationServiceEnabled(): Boolean {
        val sharedPreferences = getSharedPreferences(NOTIFICATION_LISTENER_PREFS, MODE_PRIVATE)
        return sharedPreferences.getBoolean(IS_NOTIFICATION_LISTENER_ENABLED, false)
    }

    private fun startNotificationListener() {
        val intent = Intent(this, NotificationListener::class.java)
        startService(intent)

        updateNotificationListenerSetting(true)

        Toast.makeText(this, "Notification listener started", Toast.LENGTH_SHORT)
            .show()
    }

    private fun updateNotificationListenerSetting(enabled: Boolean) {
        val sharedPreferences = getSharedPreferences(NOTIFICATION_LISTENER_PREFS, MODE_PRIVATE)
        val editor = sharedPreferences.edit()
        editor.putBoolean(IS_NOTIFICATION_LISTENER_ENABLED, enabled)
        editor.apply()
    }

    private fun requestNotificationPermission() {
        if (!hasNotificationPermission()) {
            val intent = Intent(Settings.ACTION_NOTIFICATION_LISTENER_SETTINGS)
            startActivity(intent)
            Toast.makeText(
                this,
                "Please enable notification access for LilEye",
                Toast.LENGTH_LONG
            ).show()
        }
    }

    private fun stopNotificationListener() {
        val intent = Intent(this, NotificationListener::class.java)
        stopService(intent)
        updateNotificationListenerSetting(false)
        Toast.makeText(this, "Notification listener stopped", Toast.LENGTH_SHORT).show()
    }

    private fun updateButtonState() {
        toggleButton.text = if (isNotificationServiceEnabled()) {
            "Stop Listening"
        } else {
            "Start Listening"
        }
    }

    override fun onResume() {
        super.onResume()
        updateButtonState()
    }
} 