package com.example.lab3

import android.widget.Toast

class Contact(val name: String) {

    companion object {
        private var lastContactId = 0
        lateinit var  books: List<BuyBook>
        fun createContactsList(numContacts: Int): ArrayList<Contact> {
            val contacts = ArrayList<Contact>()
            for (i in 1..numContacts) {

                contacts.add(Contact("Cat " + ++lastContactId))
            }
            return contacts
        }
        fun createBooksList(returned_books: List<BuyBook>) {
            books = returned_books

        }

    }
}