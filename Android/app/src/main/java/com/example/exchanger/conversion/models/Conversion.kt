package com.antonpriyma.android.exchanger.data.conversion.models

import com.google.gson.annotations.Expose
import com.google.gson.annotations.SerializedName

enum class ConversionType(val Type: String){
    DEFAULT("default")
}

data class Conversion(
    @SerializedName("time")
    @Expose
    var time: Long = 0,

    @SerializedName("high")
    @Expose
    var high: Float = Float.MIN_VALUE,

    @SerializedName("low")
    @Expose
    var low: Float = Float.MIN_VALUE,

    @SerializedName("open")
    @Expose
    var open: Float = Float.MIN_VALUE,


    @SerializedName("close")
    @Expose
    var close: Float = Float.MIN_VALUE,

    @SerializedName("volumefrom")
    @Expose
    var volumeFrom: Float = Float.MIN_VALUE,

    @SerializedName("volumeto")
    @Expose
    var volumeTo: Float = Float.MIN_VALUE,

    @SerializedName("conversionType")
    @Expose
    var conversionType: ConversionType = ConversionType.DEFAULT,

    @SerializedName("conversionSymbol")
    @Expose
    var conversionSymbol: String = ""
)