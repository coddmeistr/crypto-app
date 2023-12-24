package com.example.megagigacryptoapp.presentation.fragments

import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import androidx.navigation.fragment.findNavController
import com.example.megagigacryptoapp.R
import com.example.megagigacryptoapp.databinding.ChartFragmentBinding

class ChartFragment: Fragment(R.layout.chart_fragment) {

    lateinit var  binding: ChartFragmentBinding



    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        val id = this.arguments?.getInt("id",1) ?: 1

        binding = ChartFragmentBinding.bind(view)

        binding.textView10.text = id.toString()

        binding.topAppBar.setNavigationOnClickListener {
            findNavController().navigateUp()
        }
    }
}