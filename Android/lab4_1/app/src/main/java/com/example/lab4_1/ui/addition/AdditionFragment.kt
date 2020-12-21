package com.example.lab4_1.ui.addition

import android.annotation.SuppressLint
import android.app.Fragment
import android.content.Intent
import android.net.Uri
import android.os.Bundle
import android.provider.AlarmClock
import android.provider.CalendarContract
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.Button
import android.widget.TextView
import androidx.fragment.app.Fragment
import androidx.lifecycle.Observer
import androidx.lifecycle.ViewModelProviders
import com.example.lab4_1.R
import java.util.Observer


class AdditionFragment : Fragment() {

    private lateinit var additionViewModel: AdditionViewModel

    override fun onCreateView(
            inflater: LayoutInflater,
            container: ViewGroup?,
            savedInstanceState: Bundle?
    ): View? {
        additionViewModel =
                ViewModelProviders.of(this).get(AdditionViewModel::class.java)
        val root = inflater.inflate(R.layout.fragment_addition, container, false)
        val textView: TextView = root.findViewById(R.id.text_slideshow)
        additionViewModel.text.observe(viewLifecycleOwner, Observer {
            textView.text = it
        })
        val button_web: Button = root.findViewById(R.id.web)
        button_web.setOnClickListener(object : View.OnClickListener{
            override fun onClick(v: View?) {
                val i = Intent(Intent.ACTION_VIEW, Uri.parse("https://www.booktalk.org/"))
                startActivity(i)
            }})
        val button_alarm: Button = root.findViewById(R.id.alarm)
        button_alarm.setOnClickListener(object : View.OnClickListener{
            @SuppressLint("ShortAlarm")
            override fun onClick(v: View?) {
                val i = Intent(AlarmClock.ACTION_SET_ALARM).apply {
                    putExtra(AlarmClock.EXTRA_MESSAGE, "alarm message")
                    putExtra(AlarmClock.EXTRA_HOUR, 1)
                    putExtra(AlarmClock.EXTRA_MINUTES, 5)
                }
                startActivity(i)
            }})
        val button_map: Button = root.findViewById(R.id.map)
        button_map.setOnClickListener(object : View.OnClickListener{
            @SuppressLint("ShortMap")
            override fun onClick(v: View?) {
                val i = Intent(Intent.ACTION_VIEW).apply {
                    data = Uri.parse("geo:47.6,-122.3")
                }
                startActivity(i)
            }})

        val button_calender: Button = root.findViewById(R.id.calendar)
        button_calender.setOnClickListener(object : View.OnClickListener{
            @SuppressLint("button_calender")
            override fun onClick(v: View?) {
                val i = Intent(Intent.ACTION_INSERT).apply {
                    data = CalendarContract.Events.CONTENT_URI
                    putExtra(CalendarContract.Events.TITLE, "title")
                    putExtra(CalendarContract.Events.EVENT_LOCATION, "location")
                    putExtra(CalendarContract.EXTRA_EVENT_BEGIN_TIME, 20.00)
                    putExtra(CalendarContract.EXTRA_EVENT_END_TIME, 23.30)
                }
                startActivity(i)
            }})
        return root


    }

}