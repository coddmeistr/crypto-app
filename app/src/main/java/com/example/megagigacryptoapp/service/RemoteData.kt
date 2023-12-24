package com.example.megagigacryptoapp.service

import android.util.Log
import javax.inject.Inject
import javax.inject.Singleton


class RemoteData @Inject constructor(private val service: Service) {

    suspend fun login(email: String, password: String): USerDataFromApi?{
        val res = service.login(email, password)

        if(res.isSuccessful){
            val data = res.body()

            if (data != null) {
                return USerDataFromApi(data.fastKey, data.slowlyKey)
            }
        }else{
            return USerDataFromApi("rere","rer")
        }
        return USerDataFromApi("rere","rer")
    }

    suspend fun refreshToken(slowlyToken: String): USerDataFromApi?{
        val res = service.refreshToken(slowlyToken)

        if(res.isSuccessful){
            val data = res.body()

            if(data != null){
                return USerDataFromApi(data.fastKey ,data.slowlyKey)
            }
        }

        return null
    }
}