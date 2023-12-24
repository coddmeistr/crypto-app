package com.example.megagigacryptoapp.presentation.viewModel

import android.util.Log
import androidx.lifecycle.SavedStateHandle
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.megagigacryptoapp.data.MainRepository
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.update
import kotlinx.coroutines.launch
import java.net.PasswordAuthentication
import javax.inject.Inject



data class LoginState(
    val email: String = "",
    val password: String = "",
    val isFetch: Boolean = false,

)

@HiltViewModel
class LoginViewModel @Inject constructor(private val repository: MainRepository): ViewModel(){

    private val _uiState = MutableStateFlow(LoginState())

    val uiState = _uiState.asStateFlow()

    private val _rightPassword = MutableStateFlow(false)

    val  rightPassword = _rightPassword.asStateFlow()
    fun login(){

        _uiState.update {
            it.copy(isFetch = true)
        }
        viewModelScope.launch {
            val res = repository.login(uiState.value.email,uiState.value.password )
            if(res != null){
                _rightPassword.update { true }
            }

            _uiState.update {
                it.copy(isFetch = false)
            }
        }
    }

    fun updateLogin(s: String){
        _uiState.update { it.copy(email = s) }
    }

    fun updatePassword(s: String){
        _uiState.update { it.copy(password = s) }
    }
}