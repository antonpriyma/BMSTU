package com.example.exchanger.ui.conversion

import android.os.Bundle
import android.preference.PreferenceManager
import android.widget.Toast
import androidx.activity.OnBackPressedCallback
import androidx.navigation.fragment.NavHostFragment
import androidx.preference.Preference
import androidx.preference.PreferenceFragmentCompat
import com.example.exchanger.R


class SettingsFragment : PreferenceFragmentCompat() {

    override fun onCreatePreferences(savedInstanceState: Bundle?, rootKey: String?) {
        setPreferencesFromResource(R.xml.preferences, rootKey)
    }

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        bindPreferenceSummaryToValue(findPreference("days_limit"))

        val onBackPressedCallback = object : OnBackPressedCallback(true) {
            override fun handleOnBackPressed() {
                NavHostFragment.findNavController(this@SettingsFragment).navigate(R.id.action_settingsFragment_to_listFragment)
            }
        }
        requireActivity().onBackPressedDispatcher.addCallback(
            this, onBackPressedCallback
        )
    }

    private val sBindPreferenceSummaryToValueListener =
        Preference.OnPreferenceChangeListener { preference, value ->
            var res: Boolean = true
            if (preference.key == "days_limit") {
                res = value.toString() != "" && value.toString().matches(Regex("\\d*"))


            }

            if (!res) {
                val t = Toast.makeText(context, "Невалидные данные", Toast.LENGTH_LONG)
                t.show()
            }
            res
        }


    private fun bindPreferenceSummaryToValue(preference: Preference?) {
        preference?.onPreferenceChangeListener = sBindPreferenceSummaryToValueListener

        sBindPreferenceSummaryToValueListener.onPreferenceChange(
            preference,
            PreferenceManager
                .getDefaultSharedPreferences(preference?.context)
                .getString(preference?.key, "")
        )
    }
}


