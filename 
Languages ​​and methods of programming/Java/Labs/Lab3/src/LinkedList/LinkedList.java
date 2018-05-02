package LinkedList;

import java.util.Iterator;

public class LinkedList<T> implements Iterable<T>{
    public LinkedList cur;
    public T data;
    public LinkedList next;

    public LinkedList(){
        this.data=null;
        this.next=null;
        cur=this;
    }

    public T getelement(){
        return this.data;
    }

    public void setnext(){
        this.next= new LinkedList();
    }

    public void setNext(LinkedList n){
        this.next=n;
    }

    public void add(T element){
        if (this.data==null){
            this.setelement(element);
        }
        LinkedList obj = new LinkedList();
        obj.setelement(element);
        this.setNext(obj);
    }

    public void setelement(T value){
        this.data = value;
    }

    public LinkedList getNext(){
        return this.next;
    }

    public Object getNextdata(){
        return this.next.data;
    }

    @Override
    public Iterator<T> iterator() {
        return new ListIterator();
    }

    public String toString(){
        String s = new String();
        s += this.data;
        return s;
    }

    private class ListIterator implements Iterator<T>{
        private LinkedList obj;

        public ListIterator(){
            obj = cur;
        }

        public boolean hasNext(){
            return (obj!=null);
        }

        public T next(){
            T value = (T) obj.data;
            obj = obj.next;
            return value;
        }
    }
}
