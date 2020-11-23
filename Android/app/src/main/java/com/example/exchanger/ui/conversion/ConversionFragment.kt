package com.example.exchanger.ui.conversion

import android.preference.PreferenceManager
import android.view.Menu
import android.view.MenuInflater
import android.view.MenuItem
import android.view.View
import android.widget.AdapterView
import androidx.navigation.fragment.findNavController
import com.antonpriyma.android.exchanger.domain.conversion.models.Conversion
import com.example.exchanger.R
import com.example.exchanger.mvp.CleanFragment
import com.example.exchanger.ui.App
import com.example.exchanger.ui.conversion.di.component.DaggerConversionComponent
import kotlinx.android.synthetic.main.conversations_fragment.*

class ConversionFragment : CleanFragment<ConversionPresenter>(), ConversionView {
    override fun getLayout(): Int {
        return R.layout.conversations_fragment
    }

    override fun initialiseView() {
        conversions_recycler_view.apply {
            setHasFixedSize(true)
            layoutManager =
                androidx.recyclerview.widget.LinearLayoutManager(context)
        }



        crypto_value.setSelection(
            PreferenceManager.getDefaultSharedPreferences(context).getInt("crypto_selected", 0)
        )

        crypto_value.onItemSelectedListener = object : AdapterView.OnItemSelectedListener {
            override fun onItemSelected(
                parent: AdapterView<*>,
                view: View,
                position: Int,
                id: Long
            ) {
                val selectedItem = parent.getItemAtPosition(position).toString()
                val prefs = PreferenceManager.getDefaultSharedPreferences(view.context)
                prefs.edit().putString("crypto_value", selectedItem).apply()
                prefs.edit().putInt("crypto_selected", position).apply()
                presenter.update(
                    prefs.getString("days_limit", "10").toIntOrNull(),
                    prefs.getString("value", ""),
                    prefs.getString("crypto_value", "BTC")
                )
            } // to close the onItemSelected

            override fun onNothingSelected(p0: AdapterView<*>?) {

            }
        }
    }

    override fun hideConversions() {
        conversions_recycler_view.adapter = null
    }

    override fun showProgress() {
        progress_bar.visibility = View.VISIBLE
    }

    override fun hideProgress() {
        progress_bar.visibility = View.GONE
    }

    override fun showArticleList(articles: List<Conversion>) {
        conversions_recycler_view.adapter = context?.let { ConversionListAdapter(it, articles) }
    }

    override fun initInjector() {
        DaggerConversionComponent.builder()
            .appComponent((activity?.application as App).applicationComponent)
            .build()
            .inject(this)
    }

    override fun onCreateOptionsMenu(menu: Menu, inflater: MenuInflater) {
        inflater.inflate(R.menu.main_menu, menu);
        super.onCreateOptionsMenu(menu,inflater);
    }



    override fun onOptionsItemSelected(item: MenuItem): Boolean {
        return when (item.itemId) {
            R.id.update -> {
                val prefs = PreferenceManager.getDefaultSharedPreferences(context)
                presenter.update(
                    prefs.getString("days_limit", "10").toInt(),
                    prefs.getString("value", ""),
                    prefs.getString("crypto_value", "BTC")
                )
                return true
            }
            R.id.settings -> {
                findNavController().navigate(R.id.action_listFragment_to_settingsFragment)
                return true
            }
            else -> super.onOptionsItemSelected(item)
        }
    }
}