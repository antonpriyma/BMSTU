import java.util.AbstractSet;
import java.util.Arrays;
import java.util.Iterator;

public class IntSparseSet extends AbstractSet<Integer> {
    public int[] dense;
    public int[] sparce;
    public int high;
    public int size;
    private int low;
    private int count;

    public IntSparseSet(int low, int high) {
        this.size=high-low;
        sparce = new int[size];
        Arrays.fill(sparce, high + 1);
        dense = new int[size];
        Arrays.fill(dense, high + 1);
        this.high = high;
        this.low=low;
        this.size=0;
        this.count=0;
    }
    
    public int size() {
        return this.count;
    }

 
    public Iterator<Integer> iterator() {
        return new SparceIterator();
    }

    private class SparceIterator implements Iterator<Integer> {
        private int index;
       
        public Integer next() {
            return dense[index++];
        }
       
        public boolean hasNext() {
            return index<count;
        }

        public SparceIterator() {
            this.index=0;
        }

        public void remove(){
            IntSparseSet.this.remove(dense[index-1]);
        }
    }

    public boolean contains(int x){
        if (this.sparce[x-this.low]<this.high+1) {
            if (this.dense[sparce[x - this.low]] == x) return true;
        }
        return false;
    }
    public boolean remove(Object x) {
        int a = (Integer) x;

        if (a > high|| a < low)
            return false;
        if (sparce[a-low] != high+1&& dense[sparce[a-low]] != high+1) {
            count--;
            dense[sparce[a-low]] = dense[count];
            sparce[dense[count]-low] = sparce[a-low];
            sparce[a-low]=high+1;

            return true;
        }
        else
            return false;
    }
    
    public boolean add(Integer a){
        int x = (int)a;
        if (x>high || x<low || this.sparce[x-this.low]!=this.high+1)
            return false;
        else {
            sparce[x - this.low] = this.count;
            dense[this.count] = x;
            this.count++;
            return true;
        }
    }

    public void clear(){
        this.count=0;
        Arrays.fill(sparce, high + 1);
        Arrays.fill(dense, high + 1);
    }


}

