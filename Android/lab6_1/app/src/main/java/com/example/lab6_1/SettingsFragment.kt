package com.example.lab6_1

import android.os.Bundle
import android.util.Log
import android.widget.Toast
import androidx.preference.EditTextPreference
import androidx.preference.Preference
import androidx.preference.PreferenceFragmentCompat
import androidx.preference.SwitchPreferenceCompat

class SettingsFragment : PreferenceFragmentCompat() {

    override fun onCreatePreferences(savedInstanceState: Bundle?, rootKey: String?) {
        setPreferencesFromResource(R.xml.root_preferences, rootKey)
        val editName: EditTextPreference? = findPreference("name")
        editName!!.onPreferenceChangeListener =
            Preference.OnPreferenceChangeListener { _, newValue ->
                Log.d("inside", newValue.toString())
                if (newValue.toString() == "nasty"){
                    Toast.makeText(this.context, "nasty means ужасный not настя!!1!", Toast.LENGTH_SHORT).show()
                    false
                }else{
                    true
                }

            }
    }
}