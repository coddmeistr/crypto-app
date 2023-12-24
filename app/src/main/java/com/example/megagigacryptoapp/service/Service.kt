package com.example.megagigacryptoapp.service

import retrofit2.Response
import retrofit2.http.GET
import retrofit2.http.Header
import retrofit2.http.Query

interface Service {

    @GET("login")
    suspend fun login(@Query("email") email: String,@Query("password") password: String):Response<USerDataFromApi>


    @GET("refresh")
    suspend fun refreshToken(@Header("Authorization") slowlyKey: String): Response<USerDataFromApi>
}