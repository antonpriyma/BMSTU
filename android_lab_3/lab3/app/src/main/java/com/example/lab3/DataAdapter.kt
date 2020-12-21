import android.media.Image
import android.provider.Settings.Global.getString
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.Button
import android.widget.ImageView
import android.widget.TextView
import android.widget.Toast
import androidx.recyclerview.widget.RecyclerView
import com.example.lab3.BuyBook
import com.example.lab3.Contact
import com.example.lab3.R

class DataAdapter (private val mBooks: List<BuyBook>) : RecyclerView.Adapter<DataAdapter.ViewHolder>()
{
    inner class ViewHolder(listItemView: View) : RecyclerView.ViewHolder(listItemView) {

        val messageButton = itemView.findViewById<Button>(R.id.message_button)
        val nameTextView = itemView.findViewById<TextView>(R.id.info_text)
        val imageCat = itemView.findViewById<ImageView>(R.id.person_photo)
    }

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): DataAdapter.ViewHolder {
        val context = parent.context
        val inflater = LayoutInflater.from(context)
        val contactView = inflater.inflate(R.layout.item_card, parent, false)
        return ViewHolder(contactView)
    }

    override fun onBindViewHolder(viewHolder: DataAdapter.ViewHolder, position: Int) {
        val contact: BuyBook = mBooks.get(position)
        val textView = viewHolder.nameTextView
        textView.setText(contact.getmTitle())
        val button = viewHolder.messageButton
        val iCat = viewHolder.imageCat
        button.setOnClickListener {
            if( button.text == "new"){
                iCat.setImageResource(R.drawable.cat1)
                button.text = "new one"
            }
            else{
                iCat.setImageResource(R.drawable.cat)
                button.text = "new"
            }

        }
    }

    override fun getItemCount(): Int {
        return mBooks.size
    }


}