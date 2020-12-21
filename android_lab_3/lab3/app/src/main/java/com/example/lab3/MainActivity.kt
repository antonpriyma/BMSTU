package com.example.lab3

import DataAdapter
import android.annotation.SuppressLint
import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.util.Log
import android.view.View
import android.widget.Toast
import androidx.recyclerview.widget.LinearLayoutManager
import androidx.recyclerview.widget.RecyclerView
import io.reactivex.android.schedulers.AndroidSchedulers
import io.reactivex.disposables.Disposable
import io.reactivex.schedulers.Schedulers
import kotlinx.android.synthetic.main.activity_main.*
import kotlinx.android.synthetic.main.item_card.*

class MainActivity : AppCompatActivity() {
    lateinit var books: List<BuyBook>

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        val rvContacts = findViewById<View>(R.id.rvContacts) as RecyclerView
        books = emptyList()
        btn_search.setOnClickListener {
            if (edit_search.text.toString().isNotEmpty()) {
                books = BookQuery.create(edit_search.text.toString())!!

            }
        }
        val adapter = DataAdapter(books)
        rvContacts.adapter = adapter
        rvContacts.layoutManager = LinearLayoutManager(this)
    }

}