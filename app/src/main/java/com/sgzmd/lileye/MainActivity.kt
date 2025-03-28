package com.sgzmd.lileye

import android.content.Intent
import android.os.Bundle
import android.provider.Settings
import android.widget.Button
import android.widget.Toast
import androidx.appcompat.app.AppCompatActivity
import com.sgzmd.lileye.service.NotificationListener

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
            }
        }
    }

    private fun isNotificationServiceEnabled(): Boolean {
        val pkgName = packageName
        val flat = Settings.Secure.getString(contentResolver, "enabled_notification_listeners")
        return flat?.contains(pkgName) == true
    }

    private fun requestNotificationPermission() {
        val intent = Intent(Settings.ACTION_NOTIFICATION_LISTENER_SETTINGS)
        startActivity(intent)
        Toast.makeText(
            this,
            "Please enable notification access for LilEye",
            Toast.LENGTH_LONG
        ).show()
    }

    private fun stopNotificationListener() {
        val intent = Intent(this, NotificationListener::class.java)
        stopService(intent)
        toggleButton.text = "Start Listening"
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