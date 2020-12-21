package com.example.lab4_1.ui.addition

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel

class AdditionViewModel : ViewModel() {

    private val _text = MutableLiveData<String>().apply {
        value = "You here because you clicked the button"
    }
    val text: LiveData<String> = _text
}