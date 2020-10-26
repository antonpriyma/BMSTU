package com.example.exchanger.ui

import android.app.Application
import com.example.exchanger.ui.di.component.AppComponent
import com.example.exchanger.ui.di.component.DaggerAppComponent
import com.example.exchanger.ui.di.module.AppModule


class App : Application() {
    val applicationComponent: AppComponent by lazy {
        DaggerAppComponent.builder()
            .appModule(AppModule(this))
            .build()
    }

    override fun onCreate() {
        super.onCreate()
        initInjector()
    }

    private fun initInjector() {
        applicationComponent.inject(this)
    }
}