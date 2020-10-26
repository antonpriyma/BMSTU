package com.example.exchanger.ui.di.component

import android.content.Context
import com.example.exchanger.data.Api
import com.example.exchanger.domain.conversion.GetConversionsListUsecase
import com.example.exchanger.ui.App
import dagger.Component
import com.example.exchanger.ui.di.module.AppModule
import com.example.exchanger.ui.di.module.IOModule
import com.example.exchanger.ui.di.module.RepositoryModule
import com.example.exchanger.ui.di.module.UseCaseModule

import javax.inject.Singleton

@Singleton
@Component(modules = [AppModule::class, IOModule::class, RepositoryModule::class, UseCaseModule::class])
interface AppComponent {
  fun inject(app: App)
  fun getApplicationContext(): Context
  fun getEndpoint(): Api

  fun getArticlesListUseCase(): GetConversionsListUsecase
}
