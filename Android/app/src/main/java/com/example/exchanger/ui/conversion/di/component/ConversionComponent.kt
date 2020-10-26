package com.example.exchanger.ui.conversion.di.component

import com.example.exchanger.mvp.scope.PerActivity
import com.example.exchanger.ui.conversion.ConversionActivity
import com.example.exchanger.ui.di.component.AppComponent
import dagger.Component
import com.example.exchanger.ui.conversion.di.module.ArticlesModule

@PerActivity
@Component(dependencies = [AppComponent::class], modules = [ArticlesModule::class])
interface ConversionComponent {
  fun inject(conversionActivity: ConversionActivity)
}