package com.example.megagigacryptoapp.presentation.adapter

import android.graphics.Color
import android.graphics.drawable.GradientDrawable
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.recyclerview.widget.RecyclerView
import com.example.megagigacryptoapp.R
import com.example.megagigacryptoapp.databinding.CryptoCardBinding
import com.github.mikephil.charting.data.Entry
import com.github.mikephil.charting.data.LineData
import com.github.mikephil.charting.data.LineDataSet
import com.github.mikephil.charting.renderer.DataRenderer
import javax.inject.Inject


class CryptoCardAdapter(private val action: (id: Int)-> Unit): RecyclerView.Adapter<CryptoCardAdapter.CryptoCardHolder>() {

    private var cryptoCardList =  ArrayList<CryptoCard>()

    class CryptoCardHolder(itemView: View, private val action: (id: Int) -> Unit): RecyclerView.ViewHolder(itemView){

        val binding = CryptoCardBinding.bind(itemView)

        fun bind(cryptoCard: CryptoCard){
            val b = 3
            val entries: MutableList<Entry> = ArrayList()
            for (data in cryptoCard.miniChart) {
                // turn your data into Entry objects
                entries.add(Entry(data.x, data.y))
            }
            val dataSet = LineDataSet(entries,"label1")
            dataSet.mode = LineDataSet.Mode.CUBIC_BEZIER
            dataSet.cubicIntensity = 0.2f
            dataSet.setDrawValues(false)
            val lineData = LineData(dataSet)

            itemView.setOnClickListener{
                action(cryptoCard.id)
            }
            with(binding){
                textViewFullNameCrypto.setText(cryptoCard.fullName)
                textViewCryptoCurse.setText(cryptoCard.course)
                textViewShortNameCrypto.setText(cryptoCard.shortName)
                textViewStatsForDay.setText(cryptoCard.statistics)



                chartMini.axisLeft.isEnabled = false
                chartMini.axisRight.isEnabled = false

                // Убрать подписи значений наверху графика
                chartMini.legend.isEnabled = false

                // Убрать сетку (горизонтальные и вертикальные линии)
                chartMini.xAxis.setDrawGridLines(false)
                chartMini.axisLeft.setDrawGridLines(false)
                chartMini.axisRight.setDrawGridLines(false)

                // Убрать подписи значений на осях
                chartMini.xAxis.setDrawLabels(false)
                chartMini.axisLeft.setDrawLabels(false)

                chartMini.description.isEnabled = false
                chartMini.xAxis.setDrawAxisLine(false)
                // Включить градиент
                val topColor = Color.TRANSPARENT
                val bottomColor = Color.BLUE
                val gradientDrawable = GradientDrawable(
                    GradientDrawable.Orientation.TOP_BOTTOM, intArrayOf(topColor, bottomColor)
                )
                chartMini.background = gradientDrawable
                chartMini.setScaleEnabled(false)
                chartMini.setPinchZoom(false)
                chartMini.isHighlightPerTapEnabled = false

// Запретить прокрутку
                chartMini.setDragEnabled(false)

// Запретить двойное нажатие для масштабирования
                chartMini.setDoubleTapToZoomEnabled(false)

                chartMini.data = lineData
                chartMini.invalidate()
            }
        }
    }

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): CryptoCardHolder {
        val view = LayoutInflater.from(parent.context).inflate(R.layout.crypto_card, parent,false)
        return CryptoCardHolder(view, action)
    }

    override fun getItemCount(): Int {
        return cryptoCardList.size
    }

    override fun onBindViewHolder(holder: CryptoCardHolder, position: Int) {
        holder.bind(cryptoCardList[position])
    }
    fun setItems(list: ArrayList<CryptoCard>){
        this.cryptoCardList = list
        notifyDataSetChanged()
    }


}