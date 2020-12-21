package com.example.lab3

import android.net.Uri
import android.util.Log
import com.google.gson.JsonObject
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import retrofit2.http.GET
import retrofit2.http.Url


interface BookQuery {
    @GET
    fun getBooks(@Url url: String?): Call<JsonObject?>?
    companion object {
        fun create(q:String): List<BuyBook>? {

            val baseUrl = "https://www.googleapis.com/books/v1/"
            val retrofit = Retrofit.Builder().baseUrl(baseUrl)
                .addConverterFactory(GsonConverterFactory.create()).build()
            val service = retrofit.create(BookQuery::class.java)

            val baseUri = Uri.parse("https://www.googleapis.com/books/v1/volumes?q=$q")
            val builder: Uri.Builder = baseUri.buildUpon()
            val finalRequest = builder.toString()
            Log.d("Request", finalRequest)
            val call = service.getBooks(finalRequest)
            val books: List<BuyBook> = emptyList()
            var titles: Array<String?> = arrayOfNulls(10)
            call?.enqueue(object : Callback<JsonObject?> {
                override fun onResponse(
                    call: Call<JsonObject?>,
                    response: Response<JsonObject?>
                ) {
                    val root = response.body()

                    if (root?.get("totalItems")!!.asString !== "0") {
                        val itemArray = root?.getAsJsonArray("items")
                        if (itemArray != null) {
                            for (i in 0 until itemArray.size()) {
                                val element_item = itemArray[i]
                                val item = element_item.asJsonObject
                                val volumeInfo = item.getAsJsonObject("volumeInfo")
                                val element_title = volumeInfo["title"]

                                titles[i] = element_title.toString()
                            }
                            titles[0]?.let { Log.d("books", it) }
                            Log.d("books!!", titles.size.toString())
                        }
                    }
                }

                override fun onFailure(call: Call<JsonObject?>, t: Throwable) {
                    t.printStackTrace()
                }

            })
            return books
        }
    }
}