package com.example.lab4_1.ui.first

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel

class FirstViewModel : ViewModel() {

    private val _text = MutableLiveData<String>().apply {
        value = "Second Fragment"
    }
    val text: LiveData<String> = _text
}