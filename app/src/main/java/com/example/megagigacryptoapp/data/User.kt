package com.example.megagigacryptoapp.data

import android.provider.ContactsContract.CommonDataKinds.Email

data class User(
    val fastKey: String,
    val slowlyKey: String,
    val email: String,
    val name: String,
    val password : String,
)