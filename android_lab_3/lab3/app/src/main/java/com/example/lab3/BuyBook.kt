package com.example.lab3

import com.google.gson.JsonElement
import java.io.Serializable


class BuyBook(
    private var mThumbnail: String,
    private var mTitle: String,
    private var mCategory: String,
    private var mRating: String,
    private var mAuthors: String,
    private var mPrice: String,
    private var mDescription: String,
    private var mBuyLink: String,
    private var mPreviewLink: String
) :
    Serializable {
    fun getmThumbnail(): String {
        return mThumbnail
    }

    fun setmThumbnail(mThumbnail: String) {
        this.mThumbnail = mThumbnail
    }

    fun getmTitle(): String {
        return mTitle
    }

    fun setmTitle(mTitle: JsonElement) {
        this.mTitle = mTitle.toString()
    }

    fun getmCategory(): String {
        return mCategory
    }

    fun setmCategory(mCategory: String) {
        this.mCategory = mCategory
    }

    fun getmRating(): String {
        return mRating
    }

    fun setmRating(mRating: String) {
        this.mRating = mRating
    }

    fun getmAuthors(): String {
        return mAuthors
    }

    fun setmAuthors(mAuthors: String) {
        this.mAuthors = mAuthors
    }

    fun getmDescription(): String {
        return mDescription
    }

    fun setmDescription(mDescription: String) {
        this.mDescription = mDescription
    }

    fun getmPrice(): String {
        return mPrice
    }

    fun setmPrice(mPrice: String) {
        this.mPrice = mPrice
    }

    fun getmBuyLink(): String {
        return mBuyLink
    }

    fun setmBuyLink(mBuyLink: String) {
        this.mBuyLink = mBuyLink
    }

    fun getmPreviewLink(): String {
        return mPreviewLink
    }

    fun setmPreviewLink(mPreviewLink: String) {
        this.mPreviewLink = mPreviewLink
    }

}