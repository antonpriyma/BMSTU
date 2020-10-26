package com.antonpriyma.android.exchanger.domain.conversion.models

import java.time.LocalDate

data class Conversion(
    var date: LocalDate,
    var high: Float,
    var low: Float,
    var open: Float,
    var close: Float
)
