package com.example.megagigacryptoapp

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.view.View
import androidx.navigation.NavController
import androidx.navigation.findNavController
import androidx.navigation.fragment.NavHostFragment
import dagger.hilt.android.AndroidEntryPoint

@AndroidEntryPoint
class MainActivity : AppCompatActivity() {




    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)



        val  controller = (supportFragmentManager.findFragmentById(R.id.nav_host_fragment) as NavHostFragment)
            .navController


        controller.navigate(R.id.loginFragments);
    }


    fun onTextViewClicked(view: View) {
        // Обработка события клика на тексте
        // Здесь вы можете определить, что должно произойти при нажатии на текст
    }
}