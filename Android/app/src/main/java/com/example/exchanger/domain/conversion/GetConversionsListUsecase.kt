package com.example.exchanger.domain.conversion


import androidx.annotation.IntegerRes
import com.antonpriyma.android.exchanger.data.conversion.Repository
import com.antonpriyma.android.exchanger.domain.UseCase
import com.antonpriyma.android.exchanger.domain.conversion.models.Conversion
import com.example.exchanger.data.FromType
import com.example.exchanger.data.ToType
import io.reactivex.Scheduler
import io.reactivex.Single
import java.time.ZoneId
import java.util.*
import javax.inject.Inject

class GetConversionsListUsecase @Inject constructor(private val conversions: Repository,
                                                    subscribeScheduler: Scheduler,
                                                    postExecutionScheduler: Scheduler) : UseCase<List<Conversion>, Params>(subscribeScheduler, postExecutionScheduler) {

  override fun buildUseCaseSingle(params: Params): Single<List<Conversion>> {
      var conv = conversions.getConversions(params.days, params.fromType, params.toTypes)
      return conv.map {
          it.map{ Conversion(date = Date(it.time*1000).toInstant().atZone(ZoneId.systemDefault()).toLocalDate(),high = it.high, low = it.low,  open = it.open, close = it.close)}
      }
  }
  }

data class Params(
    val fromType: FromType,
    val toTypes: List<ToType>,
    val days: Int
)