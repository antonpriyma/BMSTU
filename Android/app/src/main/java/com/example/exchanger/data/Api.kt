package com.example.exchanger.data

import com.antonpriyma.android.exchanger.data.conversion.models.Conversions
import io.reactivex.Single
import retrofit2.http.GET
import retrofit2.http.Query

@JvmSuppressWildcards
interface Api {
    @GET("data/v2/histoday")
    fun getConversions(@Query("fsym") fromType: FromType, @Query("tsym") toType: ToType): Single<Response>
}


// TODO: remove duplicate code
enum class FromType(var type: String){
    BTC("BTC")
}

enum class ToType(var type: String){
    EUR("EUR"), USD("USD"), RUB("RUB")
}