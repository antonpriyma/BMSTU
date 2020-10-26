package com.example.exchanger.ui.di.module

import com.antonpriyma.android.exchanger.data.conversion.Repository
import com.example.exchanger.domain.conversion.GetConversionsListUsecase
import dagger.Module
import dagger.Provides
import io.reactivex.Scheduler
import io.reactivex.android.schedulers.AndroidSchedulers
import io.reactivex.schedulers.Schedulers
import javax.inject.Named
import javax.inject.Singleton

@Module
class UseCaseModule {
  @Provides
  @Singleton
  @Named("ioScheduler")
  internal fun provideIoScheduler() = Schedulers.io()

  @Provides
  @Singleton
  @Named("mainThreadScheduler")
  internal fun provideMainThreadScheduler() = AndroidSchedulers.mainThread()

  @Provides
  @Singleton
  internal fun provideGetConversionsListUseCase(repository: Repository, @Named("ioScheduler") ioScheduler: Scheduler, @Named("mainThreadScheduler") mainThreadScheduler: Scheduler): GetConversionsListUsecase =
    GetConversionsListUsecase(repository, ioScheduler, mainThreadScheduler)
}
