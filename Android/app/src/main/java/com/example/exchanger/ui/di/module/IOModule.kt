package com.example.exchanger.ui.di.module

import android.provider.Telephony.TextBasedSmsColumns.BODY
import com.example.exchanger.BuildConfig
import com.example.exchanger.data.Api
import com.google.gson.GsonBuilder
import dagger.Module
import dagger.Provides
import okhttp3.OkHttpClient
import okhttp3.logging.HttpLoggingInterceptor
import retrofit2.Retrofit
import retrofit2.adapter.rxjava2.RxJava2CallAdapterFactory
import retrofit2.converter.gson.GsonConverterFactory
import java.util.logging.Level
import javax.inject.Singleton

@Module
class IOModule {
  private val api: Api

  init {
    val okHttpBuilder = OkHttpClient.Builder()

    if (BuildConfig.BUILD_TYPE == "debug") {
      okHttpBuilder.addInterceptor { chain ->
        println(chain.request())
        chain.proceed(chain.request())
      }
    }

    val okHttpClient = okHttpBuilder

    val gson = GsonBuilder().create()

    val logging = HttpLoggingInterceptor()
    logging.setLevel(HttpLoggingInterceptor.Level.BODY);
    okHttpBuilder.addInterceptor(logging)



    val retrofit = Retrofit.Builder()
      .addCallAdapterFactory(RxJava2CallAdapterFactory.create())
      .addConverterFactory(GsonConverterFactory.create(gson))
      .baseUrl("https://min-api.cryptocompare.com")
      .client(okHttpClient.build())
      .build()

    api = retrofit.create(Api::class.java)
  }

  @Provides
  @Singleton
  internal fun provideEndpoint(): Api = api
}