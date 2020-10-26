package com.example.exchanger.data

import com.antonpriyma.android.exchanger.data.conversion.models.Conversion
import com.antonpriyma.android.exchanger.data.conversion.models.Conversions
import com.google.gson.annotations.Expose
import com.google.gson.annotations.SerializedName

enum class ResponseCode(val code: String) {
    SUCCESS("Success"), EMPTY("")
}

data class Response(
    @SerializedName("Response")
    @Expose
    var response: ResponseCode = ResponseCode.EMPTY,

    @SerializedName("Message")
    @Expose
    var message: String = "",

    @SerializedName("HasWarning")
    @Expose
    var hasWarning: Boolean = false,

    @SerializedName("Data")
    @Expose
    var data: Data? = null
)

data class Data(
    @SerializedName("Aggregated")
    @Expose
    var aggregated: Boolean = false,

    @SerializedName("TimeFrom")
    @Expose
    var timeFrom: Int = 0,

    @SerializedName("TimeTo")
    @Expose
    var timeTo: Int = 0,

    @SerializedName("Data")
    @Expose
    var data: List<Conversion>? = null
)