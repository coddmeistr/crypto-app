package com.example.megagigacryptoapp.data

import android.content.Context
import androidx.datastore.core.DataStore

import java.util.prefs.Preferences
import javax.inject.Inject

class UserPreferencesRepository @Inject constructor(private val dataStore: DataStore<Preferences>) {

    private companion object{

        const val TAG = "UserPreferenceRepo"
    }
}