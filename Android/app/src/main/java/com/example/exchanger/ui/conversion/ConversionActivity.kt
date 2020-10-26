package com.example.exchanger.ui.conversion

import android.content.Intent
import android.preference.PreferenceManager
import android.view.Menu
import android.view.MenuItem
import android.view.View
import androidx.navigation.findNavController
import com.antonpriyma.android.exchanger.domain.conversion.models.Conversion
import com.example.exchanger.R
import com.example.exchanger.data.FromType
import com.example.exchanger.data.ToType
import com.example.exchanger.ui.App
import com.example.exchanger.ui.conversion.di.component.DaggerConversionComponent
import com.example.exchanger.mvp.CleanActivity
import kotlinx.android.synthetic.main.activity_conversions.*

class ConversionActivity : CleanActivity<ConversionPresenter>(), ConversionView {
    override fun getLayout(): Int {
        return R.layout.activity_conversions
    }

    override fun initInjector() {
        DaggerConversionComponent.builder()
            .appComponent((application as App).applicationComponent)
            .build()
            .inject(this)
    }

    override fun initialiseView() {
        conversions_recycler_view.apply {
            setHasFixedSize(true)
            layoutManager =
                androidx.recyclerview.widget.LinearLayoutManager(this@ConversionActivity)
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
        conversions_recycler_view.adapter = ConversionListAdapter(this, articles)
    }

    override fun onCreateOptionsMenu(menu: Menu?): Boolean {
        super.onCreateOptionsMenu(menu)
        menuInflater.inflate(R.menu.main_menu, menu)
        return true
    }

    override fun onOptionsItemSelected(item: MenuItem): Boolean {
        return when (item.itemId) {
            R.id.update -> {
                val prefs = PreferenceManager.getDefaultSharedPreferences(this)
                presenter.update(prefs.getInt("days_limt", 10), prefs.getString("value", ""))
                return true
            }
            R.id.settings -> {
                val intent = Intent(this, SettingsActivity::class.java)
                startActivity(intent)
                return true
            }
            else -> super.onOptionsItemSelected(item)
        }
    }

}