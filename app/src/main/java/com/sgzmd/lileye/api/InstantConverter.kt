package com.sgzmd.lileye.api

import com.google.gson.*
import java.lang.reflect.Type
import java.time.Instant

class InstantConverter : JsonSerializer<Instant>, JsonDeserializer<Instant> {
    override fun serialize(src: Instant?, typeOfSrc: Type?, context: JsonSerializationContext?): JsonElement =
        JsonPrimitive(src?.toString())

    override fun deserialize(json: JsonElement?, typeOfT: Type?, context: JsonDeserializationContext?): Instant =
        Instant.parse(json?.asString)
}