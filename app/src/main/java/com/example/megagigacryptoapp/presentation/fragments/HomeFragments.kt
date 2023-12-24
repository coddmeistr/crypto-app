package com.example.megagigacryptoapp.presentation.fragments

import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import androidx.navigation.fragment.NavHostFragment
import androidx.navigation.fragment.findNavController
import androidx.navigation.ui.NavigationUI
import androidx.navigation.ui.setupWithNavController
import com.example.megagigacryptoapp.R
import com.example.megagigacryptoapp.databinding.HomeBinding
import dagger.hilt.android.AndroidEntryPoint

@AndroidEntryPoint
class HomeFragments : Fragment(R.layout.home) {

lateinit var binding: HomeBinding

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        binding = HomeBinding.bind(view)


        val navController = (childFragmentManager.findFragmentById(R.id.top_content_container) as NavHostFragment).navController

        NavigationUI.setupWithNavController(binding.bottomNavigation, navController)



// Для обновления верхней части экрана (фрагмента)
        
    }
}