package com.example.megagigacryptoapp.presentation.fragments

import android.graphics.Color
import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import com.example.megagigacryptoapp.R
import com.example.megagigacryptoapp.databinding.AccountFragmentBinding
import com.example.megagigacryptoapp.databinding.LoginFragmentBinding
import com.example.megagigacryptoapp.presentation.adapter.CurrencyCardAdapter
import com.example.megagigacryptoapp.repositoryOfData.Data
import com.github.mikephil.charting.data.PieData
import com.github.mikephil.charting.data.PieDataSet
import com.github.mikephil.charting.data.PieEntry
import dagger.hilt.EntryPoint
import dagger.hilt.android.AndroidEntryPoint
import javax.inject.Inject

@AndroidEntryPoint
class AccountFragment : Fragment(R.layout.account_fragment) {

    lateinit var binding: AccountFragmentBinding

    lateinit var adapter: CurrencyCardAdapter

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        binding = AccountFragmentBinding.bind(view)
        adapter = CurrencyCardAdapter()
        init()
        initChar()
    }


    private fun init(){
        binding.apply {
            recyclerCurrencyView.adapter = adapter
            adapter.setItem(Data.currencyList)
        }
    }

    private fun initChar(){
        val entries = ArrayList<PieEntry>()
        // Замените эти значения своими данными
        entries.add(PieEntry(30f, "Label 1"))
        entries.add(PieEntry(40f, "Label 2"))
        entries.add(PieEntry(30f, "Label 3"))

        val dataSet = PieDataSet(entries, "Your Chart Title")

        dataSet.colors = listOf(Color.rgb(255, 0, 0), Color.rgb(0, 255, 0), Color.rgb(0, 0, 255))

        val data = PieData(dataSet)

        binding.pieBalanceChart.data = data
        binding.pieBalanceChart.centerText = "$10000"
        binding.pieBalanceChart.setCenterTextColor(Color.parseColor("#3C69BF"))
        binding.pieBalanceChart.setHoleColor(Color.parseColor("#3C69BF"))
        binding.pieBalanceChart.description.isEnabled = false
        binding.pieBalanceChart.legend.isEnabled = false
        //binding.pieBalanceChart.setOnChartValueSelectedListener()
        binding.pieBalanceChart.invalidate() // Обновляем график
    }



}