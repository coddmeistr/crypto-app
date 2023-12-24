package com.example.megagigacryptoapp.presentation.adapter

import android.util.ArrayMap
import com.example.megagigacryptoapp.repositoryOfData.Coordinates

data class CryptoCard(
    val id: Int,
    val imageId: Int,
    val fullName: String,
    val shortName: String,
    val statistics: String,
    val course: String,
    val miniChart: Array<Coordinates>,
)
