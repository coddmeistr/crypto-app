package com.example.megagigacryptoapp.data

import com.example.megagigacryptoapp.service.RemoteData
import com.example.megagigacryptoapp.service.Service
import com.example.megagigacryptoapp.service.USerDataFromApi
import javax.inject.Inject
import javax.inject.Singleton


class MainRepository @Inject constructor(private val remoteDataFromApi: RemoteData){

    suspend fun login(email: String , password: String): USerDataFromApi? = remoteDataFromApi.login(email, password)

    suspend fun refresh(slowlyToken: String): USerDataFromApi? = remoteDataFromApi.refreshToken(slowlyToken)



}