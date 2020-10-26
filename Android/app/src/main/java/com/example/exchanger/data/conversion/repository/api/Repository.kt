package com.antonpriyma.android.exchanger.data.conversion.repository.api


import com.antonpriyma.android.exchanger.data.conversion.Repository
import com.antonpriyma.android.exchanger.data.conversion.models.Conversion
import com.example.exchanger.data.Api
import com.example.exchanger.data.FromType
import com.example.exchanger.data.ToType
import io.reactivex.Single
import javax.inject.Inject

class Repository @Inject constructor(private val api: Api) : Repository {
    override fun getConversions(
        days: Int,
        fromType: FromType,
        toTypes: List<ToType>
    ): Single<List<Conversion>> {
        var conv = api.getConversions(days, fromType, toTypes[0])
        return conv.map { it.data?.data }
    }
}