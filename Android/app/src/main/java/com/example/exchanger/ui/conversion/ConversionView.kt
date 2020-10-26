package com.example.exchanger.ui.conversion

import com.antonpriyma.android.exchanger.domain.conversion.models.Conversion

interface ConversionView {
    fun initialiseView()
    fun hideConversions()
    fun showProgress()
    fun hideProgress()
    fun showArticleList(articles: List<Conversion>)
}