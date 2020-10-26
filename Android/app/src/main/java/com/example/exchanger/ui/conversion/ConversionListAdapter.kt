package com.example.exchanger.ui.conversion

import android.content.Context
import android.os.Build
import android.view.LayoutInflater
import android.view.MenuItem
import android.view.View
import android.view.ViewGroup
import android.widget.PopupMenu
import android.widget.TextView
import androidx.annotation.RequiresApi
import androidx.recyclerview.widget.RecyclerView
import com.antonpriyma.android.exchanger.domain.conversion.models.Conversion
import com.example.exchanger.R
import java.time.format.DateTimeFormatter

class ConversionListAdapter(
    private val context: Context,
    private val conversions: List<Conversion>
) :
    androidx.recyclerview.widget.RecyclerView.Adapter<ConversionViewHolder>() {

    override fun onCreateViewHolder(
        parent: ViewGroup,
        viewType: Int
    ): ConversionViewHolder {
        val v = LayoutInflater.from(context)
            .inflate(R.layout.list_item_conversion, parent, false)
        return ConversionViewHolder(v)
    }

    @RequiresApi(Build.VERSION_CODES.O)
    override fun onBindViewHolder(holder: ConversionViewHolder, position: Int) {
        val conv = conversions[position]
        holder.itemView.setOnClickListener(holder)
        holder.bind(conv)
    }

    override fun getItemCount() = conversions.size
}

class ConversionViewHolder(v: View) : RecyclerView.ViewHolder(v), View.OnClickListener,
    PopupMenu.OnMenuItemClickListener {
    private var mTitleView: TextView? = null
    private var mYearView: TextView? = null
    private val popupMenu: PopupMenu? = null
    private var max: Float = Float.MIN_VALUE
    private var min: Float = Float.MIN_VALUE

    init {
        mTitleView = itemView.findViewById(R.id.date_text)
        mYearView = itemView.findViewById(R.id.course_text)

    }

    @RequiresApi(Build.VERSION_CODES.O)
    fun bind(conv: Conversion) {
        var formatter = DateTimeFormatter.ofPattern("dd-MM")
        max = conv.high
        min = conv.low

        mTitleView?.text = conv.date.format(formatter)
        mYearView?.text = conv.open.toString()
    }

    override fun onClick(p0: View?) {
        showMenu(p0)
    }

    private fun showMenu(v: View?) {
        val popupMenu: PopupMenu? = PopupMenu(v?.context, v)
        popupMenu?.inflate(R.menu.popup_menu)
        popupMenu?.setOnMenuItemClickListener(this)
        popupMenu?.show()
    }

    override fun onMenuItemClick(p0: MenuItem?): Boolean {
        when (p0?.itemId) {
            R.id.max -> {
                mYearView?.text = this.max.toString()
                return true
            }
            R.id.min -> {
                mYearView?.text = this.min.toString()
                return true
            }
        }
        return false
    }
}