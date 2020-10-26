package com.antonpriyma.android.exchanger.data.conversion.models

import com.google.gson.annotations.Expose
import com.google.gson.annotations.SerializedName

data class Conversions(
    @SerializedName("Data")
    @Expose
    var conversions: List<Conversion>? = null
)
