package com.example.megagigacryptoapp.presentation.fragments

import android.os.Bundle
import android.text.Editable
import android.text.TextWatcher
import android.util.Log
import android.view.View
import androidx.fragment.app.Fragment
import androidx.fragment.app.viewModels
import androidx.lifecycle.lifecycleScope
import androidx.navigation.fragment.findNavController
import com.example.megagigacryptoapp.R
import com.example.megagigacryptoapp.databinding.LoginFragmentBinding
import com.example.megagigacryptoapp.presentation.viewModel.LoginViewModel
import dagger.hilt.android.AndroidEntryPoint
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext

@AndroidEntryPoint
class LoginFragments: Fragment(R.layout.login_fragment) {

  private lateinit var binding: LoginFragmentBinding


  private val viewModel: LoginViewModel by viewModels()


    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        binding = LoginFragmentBinding.bind(view)

        val uiState = viewModel.uiState



        viewLifecycleOwner.lifecycleScope.launch {
            viewModel.rightPassword.collect{
            if(it){
                Log.d("crypto","good")
                withContext(Dispatchers.Main){
                    navigate()
                }}
            }
        }

        binding.button2.setOnClickListener {
            viewModel.login()
        }

        binding.textInputEmail.setText(uiState.value.email)
        binding.textInputEmail.addTextChangedListener(object : TextWatcher{
            override fun afterTextChanged(p0: Editable?) {

            }

            override fun beforeTextChanged(p0: CharSequence?, p1: Int, p2: Int, p3: Int) {

            }

            override fun onTextChanged(p0: CharSequence?, p1: Int, p2: Int, p3: Int) {
                viewModel.updateLogin(p0.toString())
            }
        })



        binding.textInputEditPassword.setText(uiState.value.password)
        binding.textInputEditPassword.addTextChangedListener(object : TextWatcher{
            override fun afterTextChanged(p0: Editable?) {

            }

            override fun beforeTextChanged(p0: CharSequence?, p1: Int, p2: Int, p3: Int) {

            }

            override fun onTextChanged(p0: CharSequence?, p1: Int, p2: Int, p3: Int) {
                viewModel.updatePassword(p0.toString())
            }
        })



    }

    fun navigate(){
        findNavController().navigate(R.id.action_loginFragments_to_homeFragments)
    }





}