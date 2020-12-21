package com.example.lab6_1

import android.os.Bundle
import com.google.android.material.floatingactionbutton.FloatingActionButton
import com.google.android.material.snackbar.Snackbar
import androidx.appcompat.app.AppCompatActivity
import android.view.Menu
import android.view.MenuItem

class MainActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        setSupportActionBar(findViewById(R.id.toolbar))
    }

    override fun onCreateOptionsMenu(menu: Menu): Boolean {
        menuInflater.inflate(R.menu.menu_main, menu)
        return true
    }

    override fun onOptionsItemSelected(item: MenuItem): Boolean {
        val id = item.getItemId()
        if(id == R.id.action_settings){
            supportFragmentManager
                    .beginTransaction()
                    .replace(R.id.container, SettingsFragment())
                    .addToBackStack( "tag" )
                    .commit()
            supportActionBar?.setDisplayHomeAsUpEnabled(true)
        }
        return super.onOptionsItemSelected(item)
    }
}