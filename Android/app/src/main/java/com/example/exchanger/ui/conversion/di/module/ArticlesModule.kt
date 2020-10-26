package com.example.exchanger.ui.conversion.di.module

import com.example.exchanger.domain.conversion.GetConversionsListUsecase
import com.example.exchanger.mvp.scope.PerActivity
import com.example.exchanger.ui.conversion.ConversionPresenter
import dagger.Module
import dagger.Provides


@Module
class ArticlesModule {
  @PerActivity
  @Provides
  internal fun provideArticlesPresenter(getArticlesListUseCase: GetConversionsListUsecase) = ConversionPresenter(getArticlesListUseCase)
}
