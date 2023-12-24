package com.example.megagigacryptoapp.service

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
data class USerDataFromApi(
    @SerialName("accessToken") val fastKey: String,
    @SerialName("refreshToken") val slowlyKey: String
)
