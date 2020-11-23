package com.example.exchanger.ui.conversion

import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
import androidx.navigation.NavController
import androidx.navigation.Navigation
import androidx.navigation.Navigation.findNavController
import com.example.exchanger.R
import com.example.exchanger.ui.App
import com.example.exchanger.ui.conversion.di.component.DaggerConversionComponent

class ConversionActivity : AppCompatActivity() {

    private lateinit var navController: NavController
    override fun onSupportNavigateUp() =
        findNavController(this, R.id.navHostFragment).navigateUp()

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_conversions)
        navController = Navigation.findNavController(this, R.id.navHostFragment);
    }


}