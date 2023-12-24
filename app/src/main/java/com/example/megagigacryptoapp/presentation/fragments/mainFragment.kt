package com.example.megagigacryptoapp.presentation.fragments

import android.os.Bundle
import android.view.View
import androidx.core.os.bundleOf
import androidx.fragment.app.Fragment
import androidx.navigation.fragment.findNavController
import com.example.megagigacryptoapp.R
import com.example.megagigacryptoapp.databinding.LoginFragmentBinding
import com.example.megagigacryptoapp.databinding.MainFragmentBinding
import com.example.megagigacryptoapp.presentation.adapter.CryptoCardAdapter
import com.example.megagigacryptoapp.repositoryOfData.Data
import dagger.hilt.android.AndroidEntryPoint

@AndroidEntryPoint
class mainFragment:Fragment(R.layout.main_fragment) {
    lateinit var binding: MainFragmentBinding

    lateinit var cryptoCardAdapter: CryptoCardAdapter


    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        binding = MainFragmentBinding.bind(view)
        cryptoCardAdapter = CryptoCardAdapter { id ->
            findNavController().navigate(
                R.id.action_mainFragment_to_chartFragment,
                bundleOf("id" to id)
            )
        }
        init()

    }

    private fun init(){
        binding.apply {
            RecyclerViewCryptoCard.adapter = cryptoCardAdapter

            cryptoCardAdapter.setItems(Data.arrayOfCryptoCard)
        }
    }

}