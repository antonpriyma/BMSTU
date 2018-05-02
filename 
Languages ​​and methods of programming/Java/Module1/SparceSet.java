import java.util.AbstractSet;
import java.util.ArrayList;
import java.util.Iterator;

public class SparseSet<T extends Hintable> extends AbstractSet<T> {
    public ArrayList<T> dense;
    public int high;
    public int size;
    private int low;
    private int count;

    public SparseSet() {
        this.size=high-low;;
        dense = new ArrayList<>();
        this.size=0;
        count=0;
    }
    @Override
    public int size() {
        return count;
    }

    @Override
    public Iterator<T> iterator() {
        return new SparceIterator();
    }

    private class SparceIterator implements Iterator<T> {
        private int index;
        @Override
        public T next() {
            return dense.get(index++);
        }
        @Override
        public boolean hasNext() {
            return index<count;
        }

        public SparceIterator() {
            this.index=0;
        }

        public void remove(){
            SparseSet.this.remove(dense.get(index-1));
        }
    }

    public boolean contains(T x){
        if ( count>0 && dense.get(x.hint())==x ) return true;
        else
        return false;
    }
    public boolean remove(T x) {
        if (contains(x)){
            count--;
            dense.get(count).setHint(x.hint());
            dense.set(x.hint(),dense.get(count));
            return true;
        }
        return false;
    }

    public boolean add(T a){
        if (!contains(a)) {
            dense.add(a);
            a.setHint(count++);
            return true;
        }
        return false;
    }

    public void clear(){
        this.count=0;
    }


}

