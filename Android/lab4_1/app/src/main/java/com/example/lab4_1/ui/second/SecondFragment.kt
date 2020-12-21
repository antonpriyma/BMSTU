package com.example.lab4_1.ui.second

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.Button
import android.widget.TextView
import androidx.fragment.app.Fragment
import androidx.fragment.app.FragmentTransaction
import androidx.lifecycle.Observer
import androidx.lifecycle.ViewModelProviders
import androidx.navigation.Navigation
import com.example.lab4_1.R
import com.example.lab4_1.ui.addition.AdditionFragment

class SecondFragment : Fragment() {

    private lateinit var secondViewModel: SecondViewModel

    override fun onCreateView(
            inflater: LayoutInflater,
            container: ViewGroup?,
            savedInstanceState: Bundle?
    ): View? {
        secondViewModel =
                ViewModelProviders.of(this).get(SecondViewModel::class.java)
        val root = inflater.inflate(R.layout.fragment_second, container, false)
        val button: Button = root.findViewById(R.id.button)
        button.setOnClickListener(object : View.OnClickListener{
            override fun onClick(v: View?) {
                Navigation.createNavigateOnClickListener(R.id.action_second_to_second_third)
            }})
        val textView: TextView = root.findViewById(R.id.text_gallery)
        secondViewModel.text.observe(viewLifecycleOwner, Observer {
            textView.text = it
        })

        return root
    }
}