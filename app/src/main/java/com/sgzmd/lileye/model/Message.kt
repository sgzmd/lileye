package com.sgzmd.lileye.model

import android.os.Parcelable
import kotlinx.parcelize.Parcelize

@Parcelize
data class Message(
    val packageName: String,
    val title: String?,
    val text: String?,
    val timestamp: Long,
    val extras: Map<String, String> = emptyMap()
) : Parcelable 