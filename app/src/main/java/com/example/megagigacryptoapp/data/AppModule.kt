package com.example.megagigacryptoapp.data

import android.content.Context
import com.example.megagigacryptoapp.service.RemoteData
import com.example.megagigacryptoapp.service.Service
import com.jakewharton.retrofit2.converter.kotlinx.serialization.asConverterFactory
import dagger.Module
import dagger.Provides
import dagger.hilt.InstallIn
import dagger.hilt.components.SingletonComponent
import kotlinx.serialization.json.Json
import okhttp3.MediaType.Companion.toMediaType
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import javax.inject.Singleton



@Module
@InstallIn(SingletonComponent::class)
object AppModule {

    @Provides
    fun baseUrl():String = "https://android-kotlin-fun-mars-server.appspot.com/"


    @Provides
    fun provideJson(): Json = Json { ignoreUnknownKeys = true }

    @Provides
    @Singleton
    fun provideMainRetrofit(baseUrl: String): Retrofit = Retrofit.Builder()
        .baseUrl(baseUrl)
        .addConverterFactory(GsonConverterFactory.create())
        .build()


    @Provides
    @Singleton
    fun provideService(retrofit: Retrofit): Service = retrofit.create(Service::class.java)


    @Provides
    @Singleton
    fun provideRemoteData(service: Service): RemoteData = RemoteData(service)

    @Provides
    @Singleton
    fun provideMainRepository(remoteData: RemoteData) : MainRepository = MainRepository(remoteData)



}