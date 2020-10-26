package com.example.exchanger.ui.di.module

import com.antonpriyma.android.exchanger.data.conversion.Repository
import com.example.exchanger.data.Api
import dagger.Module
import dagger.Provides
import javax.inject.Singleton
import com.antonpriyma.android.exchanger.data.conversion.repository.api.Repository as apiRepository

@Module
class RepositoryModule {
    @Provides
    @Singleton
    internal fun provideArticleRepository(api: Api): Repository = apiRepository(api)
}