package com.example.exchanger.ui.conversion

import com.antonpriyma.android.exchanger.domain.conversion.models.Conversion
import io.reactivex.observers.DisposableSingleObserver

class ConversionListObserver(private val presenter: ConversionPresenter): DisposableSingleObserver<List<Conversion>>() {
  override fun onSuccess(articlesList: List<Conversion>) {
    presenter.showArticleList(articlesList)
  }

  override fun onError(e: Throwable) {
    e.printStackTrace()
  }
}
