import java.util.ArrayList;
import java.util.Collection;
import java.util.Iterator;


public class StringList implements Iterable<String>,Collection<String>{
    private ArrayList<String> data;

    public StringList() {
        data = new ArrayList<>();
    }

    public boolean add(String s) {
        return data.add(s);
    }

    @Override
    public boolean remove(Object s) {
        return data.remove(s);
    }

    @Override
    public boolean containsAll(Collection<?> c) {
        return data.containsAll(c);
    }

    @Override
    public boolean addAll(Collection<? extends String> c) {
        return data.addAll(c);
    }

    @Override
    public boolean removeAll(Collection<?> c) {
        return data.removeAll(c);
    }

    @Override
    public boolean retainAll(Collection<?> c) {
        return data.retainAll(c);
    }

    @Override
    public void clear() {
        data.clear();
    }

    public String get(int k) {
        String s = data.get(k);
        String nextS = data.get(k + 1);
        int n = s.length(), n1 = nextS.length(),j = 0;

        for (int i = n - 1; i >= 0 && j < n1; i--, j++) {
            if (s.charAt(i) != nextS.charAt(j)) {
                break;
            }
        }
        return s.substring(n-j,n);
    }

    @Override
    public boolean isEmpty(){
        return data.isEmpty();
    }


    @Override
    public int size() {
        return data.size();
    }

    @Override
    public boolean contains(Object o) {
        for (String i:this){
            if (i.equals(o)){
                return true;
            }
        }
        return false;
    }

    public Iterator<String> iterator(){
        return new StringIterator();
    }

    @Override
    public Object[] toArray() {
        return data.toArray();
    }

    @Override
    public <T> T[] toArray(T[] a) {
        return data.toArray(a);
    }

    @Override
    public String toString(){
        String s = new String();
        for (String i :this){
            if (!i.equals("")) {
                s += i + " ";
            }
        }
        s=s.substring(0,s.length()-1);
        return s;
    }

    /**************************Iterator***********************/
    private class StringIterator implements Iterator<String>{
        private int pos;

        public StringIterator(){
            pos=0;
        }

        public boolean hasNext(){
            return pos<data.size();
        }

        public String next(){
            if (pos==data.size()-1){
                return data.get(pos++);
            }
            return get(pos++);
        }
    }
}
