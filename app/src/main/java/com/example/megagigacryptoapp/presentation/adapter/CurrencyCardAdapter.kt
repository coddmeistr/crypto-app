package com.example.megagigacryptoapp.presentation.adapter

import android.content.Context
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.appcompat.content.res.AppCompatResources.getDrawable
import androidx.core.content.ContextCompat
import androidx.recyclerview.widget.RecyclerView
import com.example.megagigacryptoapp.R
import com.example.megagigacryptoapp.databinding.CurrencyCardBinding
import java.util.zip.Inflater
import kotlin.coroutines.coroutineContext

class CurrencyCardAdapter: RecyclerView.Adapter<CurrencyCardAdapter.CurrencyCardHolder>(){

    private var currencyCardList = ArrayList<CurrencyCard>()


    class CurrencyCardHolder(itemView: View, private val context: Context): RecyclerView.ViewHolder(itemView){
        val binding = CurrencyCardBinding.bind(itemView)

        fun bind(currencyCard: CurrencyCard){
            binding.currencyImage.setImageDrawable(ContextCompat.getDrawable(context,currencyCard.idImage))
            binding.currencyName.text = currencyCard.nameCurrency
            binding.currencyValue.text = currencyCard.valueCurrency
            binding.percentCurrencyValue.text = currencyCard.percentage
        }

    }

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): CurrencyCardHolder {
        val view = LayoutInflater.from(parent.context).inflate(R.layout.currency_card, parent, false)

        return  CurrencyCardHolder(view, parent.context)
    }

    override fun getItemCount(): Int {
        return currencyCardList.size
    }

    override fun onBindViewHolder(holder: CurrencyCardHolder, position: Int) {
        holder.bind(currencyCardList[position])
    }

    fun setItem(list: ArrayList<CurrencyCard>){
        this.currencyCardList = list
        notifyDataSetChanged()
    }
}