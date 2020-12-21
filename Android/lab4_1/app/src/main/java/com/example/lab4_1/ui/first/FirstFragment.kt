package com.example.lab4_1.ui.first

import android.app.Fragment
import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.fragment.app.Fragment
import androidx.lifecycle.Observer
import androidx.lifecycle.ViewModelProviders
import com.example.lab4_1.R
import java.util.Observer

class FirstFragment : Fragment() {

    private lateinit var firstViewModel: FirstViewModel

    override fun onCreateView(
            inflater: LayoutInflater,
            container: ViewGroup?,
            savedInstanceState: Bundle?
    ): View? {
        firstViewModel =
                ViewModelProviders.of(this).get(FirstViewModel::class.java)
        val root = inflater.inflate(R.layout.fragment_first, container, false)
        val textView: TextView = root.findViewById(R.id.text_home)
        firstViewModel.text.observe(viewLifecycleOwner, Observer {
            textView.text = it
        })
        return root
    }
}