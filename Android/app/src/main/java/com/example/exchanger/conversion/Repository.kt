package com.antonpriyma.android.exchanger.data.conversion

import com.antonpriyma.android.exchanger.data.conversion.models.Conversion
import com.example.exchanger.data.FromType
import com.example.exchanger.data.ToType
import io.reactivex.Single

interface Repository {
    fun getConversions(fromType: FromType, toTypes: List<ToType>): Single<List<Conversion>>
}