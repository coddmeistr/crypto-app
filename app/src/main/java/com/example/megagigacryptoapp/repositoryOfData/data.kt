package com.example.megagigacryptoapp.repositoryOfData

import androidx.collection.arrayMapOf
import com.example.megagigacryptoapp.R
import com.example.megagigacryptoapp.presentation.adapter.CryptoCard
import com.example.megagigacryptoapp.presentation.adapter.CurrencyCard

object Data {
    val arrayOfCryptoCard = arrayListOf<CryptoCard>(
        CryptoCard(1,
            R.drawable.bitcoin,
            "Bitcoin",
            "BTC",
            "+15.55%",
            "$120.20",
            arrayOf(
                Coordinates(1.0F,4.0F),
                Coordinates(2.0F,5.0F),
                Coordinates(3.0F,6.0F),
                Coordinates(4.0F,8.0F)
            )
        ),
        CryptoCard(2,R.drawable.bitcoin,
            "Bitcoin",
            "BTC",
            "+15.55%",
            "$120.20",
            arrayOf(
                Coordinates(1.0F,4.0F),
                Coordinates(2.0F,5.0F),
                Coordinates(3.0F,1.0F),
                Coordinates(4.0F,8.0F)
            )),

        CryptoCard(3,R.drawable.bitcoin,
            "Bitcoin",
            "BTC",
            "+15.55%",
            "$120.20",
            arrayOf(
                Coordinates(1.0F,4.0F),
                Coordinates(2.0F,8.0F),
                Coordinates(3.0F,2.0F),
                Coordinates(4.0F,11.0F)
            )),
        CryptoCard(4,R.drawable.bitcoin,
            "Bitcoin",
            "BTC",
            "+15.55%",
            "$120.20",
            arrayOf(
                Coordinates(1.0F,4.0F),
                Coordinates(2.0F,5.0F),
                Coordinates(3.0F,6.0F),
                Coordinates(4.0F,8.0F)
            ))
    )

    val currencyList = arrayListOf<CurrencyCard>(
        CurrencyCard(
            0,
            R.drawable.bitcoin,
            "bitcion",
            "%40",
            "$ 2.000",
        ),
        CurrencyCard(
            1,
            R.drawable.litecoin,
            "litecoin",
            "%10",
            "$ 1.000",
        ),
        CurrencyCard(
            2,
            R.drawable.dash,
            "dash",
            "%5",
            "$ 500",
        ),
        CurrencyCard(
            3,
            R.drawable.eos,
            "eos",
            "%40",
            "$ 2.000",
        ),CurrencyCard(
            4,
            R.drawable.ethereum,
            "ethereum",
            "%2",
            "$ 200",
        )

    )
}