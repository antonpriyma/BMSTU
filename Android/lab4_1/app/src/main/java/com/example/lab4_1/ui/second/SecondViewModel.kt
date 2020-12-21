package com.example.lab4_1.ui.second

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel

class SecondViewModel : ViewModel() {

    private val _text = MutableLiveData<String>().apply {
        value = "click the button!"
    }
    val text: LiveData<String> = _text
}