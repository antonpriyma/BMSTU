package com.example.lab3

import android.util.Log
import com.google.gson.JsonArray
import com.google.gson.JsonObject
import java.util.*


internal object BookBuilder {
    fun jsonParse(root: JsonObject): List<BuyBook> {
        val bookData: MutableList<BuyBook> = ArrayList()
        var item: JsonObject
        var volumeInfo: JsonObject
        var imageLinks: JsonObject
        var saleInfo: JsonObject
        var listPrice: JsonObject
        val itemArray: JsonArray
        var authorsArray: JsonArray
        var categoryArray: JsonArray
        var title: String
        var authors: String
        var category: String
        var smallThumbnail: String
        var avgRating: String
        var price: String?
        var description: String
        var buyLink: String
        var previewLink: String
        itemArray = root.getAsJsonArray("items")
        for (i in 0 until itemArray.size()) {
            val element_item = itemArray[i]
            item = element_item.asJsonObject
            volumeInfo = item.getAsJsonObject("volumeInfo")
            val element_title = volumeInfo["title"]
            title = element_title.asString
            authors = ""
            if (volumeInfo.has("authors") && !volumeInfo["authors"].isJsonNull) {
                authorsArray = volumeInfo.getAsJsonArray("authors")
                if (authorsArray.size() > 1) {
                    for (j in 0 until authorsArray.size()) {
                        val element_author = authorsArray[j]
                        authors += """
                            ${element_author.asString}
                            
                            """.trimIndent()
                    }
                } else {
                    val element_author = authorsArray[0]
                    authors += element_author.asString
                }
            } else {
                authors = "Unknown"
            }
            category = ""
            if (volumeInfo.has("categories") && !volumeInfo["categories"].isJsonNull) {
                categoryArray = volumeInfo.getAsJsonArray("categories")
                if (categoryArray.size() > 1) {
                    for (j in 0 until categoryArray.size()) {
                        val element_category = categoryArray[j]
                        category += element_category.asString + "|"
                    }
                } else {
                    val element_category = categoryArray[0]
                    category += element_category.asString
                }
            } else {
                category = "No Category"
            }
            avgRating =
                if (volumeInfo.has("averageRating") && !volumeInfo["averageRating"].isJsonNull) {
                    val element_rating = volumeInfo["averageRating"]
                    element_rating.asString
                } else {
                    "0"
                }
            if (volumeInfo.has("imageLinks") && !volumeInfo["imageLinks"].isJsonNull) {
                imageLinks = volumeInfo.getAsJsonObject("imageLinks")
                val element_smallThumbnail = imageLinks["smallThumbnail"]
                smallThumbnail = element_smallThumbnail.asString
            } else {
                smallThumbnail =
                    "https://www.sylvansport.com/wp/wp-content/uploads/2018/11/image-placeholder-1200x800.jpg"
            }
            description =
                if (volumeInfo.has("description") && !volumeInfo["description"].isJsonNull) {
                    val element_description = volumeInfo["description"]
                    element_description.asString
                } else {
                    "No description for this Book!"
                }
            previewLink =
                if (volumeInfo.has("previewLink") && !volumeInfo["previewLink"].isJsonNull) {
                    val element_preview = volumeInfo["previewLink"]
                    element_preview.asString
                } else {
                    " "
                }
            saleInfo = item.getAsJsonObject("saleInfo")
            price = " "
            buyLink = " "
            if (saleInfo.has("saleability") && !saleInfo["saleability"].isJsonNull) {
                val element_saleability = saleInfo["saleability"]
                if (element_saleability.asString == "FOR_SALE") {
                    if (saleInfo.has("listPrice")) {
                        listPrice = saleInfo.getAsJsonObject("listPrice")
                        val element_price = listPrice["amount"]
                        price = java.lang.Double.toString(element_price.asDouble)
                    }
                    if (saleInfo.has("currencyCode")) {
                        val element_currencyCode = saleInfo["currencyCode"]
                        price += element_currencyCode.asString
                    }
                    if (saleInfo.has("buyLink")) {
                        val element_buyLink = saleInfo["buyLink"]
                        buyLink = element_buyLink.asString
                    }
                } else {
                    price = "Free"
                    buyLink = " "
                }
            }
            Log.d("inside", title)
            bookData.add(
                BuyBook(
                    smallThumbnail,
                    title,
                    category,
                    avgRating,
                    authors,
                    price!!,
                    description,
                    buyLink,
                    previewLink
                )
            )
        }
        return bookData
    }
}
