import java.util.*;


public class SkipList<K extends Comparable<K>, V> extends AbstractMap<K, V> {


  
    public SkipList(int levels) {
        this.levels = levels;
        head = new Node(null, levels);
    }

    public boolean isEmpty() {
        if (size == 0) return true;
        return false;
    }

    @Override
    public Set<Map.Entry<K, V>> entrySet() {
        return new NodeSet(size);
    }

    public boolean containsKey(Object k) {
        ArrayList<Node> help;
        Node y = head;
        ArrayList<Node> help1 = new ArrayList<>();
        for (int i = 0; i < levels; i++) {
            help1.add(null);
        }
        for (int i = levels - 1; i >= 0; i--) {
            while (y.getNext().get(i) != null && y.getNext().get(i).getKey().compareTo((K) k) < 0)
                y = y.getNext().get(i);
            help1.set(i, y);
        }
        help = help1;
        Node x = help.get(0).getNext().get(0);
        boolean verdict = x != null && x.getKey().equals((K) k);
        return verdict;
    }

    public V put(K k, V v) {
        Map.Entry<K, V> a = new AbstractMap.SimpleEntry<K, V>(k, v);
        Node x = new Node(a, levels);
        ArrayList<Node> help;
        Node y = head;
        ArrayList<Node> help1 = new ArrayList<>();
        for (int i = 0; i < levels; i++) {
            help1.add(null);
        }
        for (int i = levels - 1; i >= 0; i--) {
            while (y.getNext().get(i) != null && y.getNext().get(i).getKey().compareTo(k) < 0)
                y = y.getNext().get(i);
            help1.set(i, y);
        }
        help = help1;

        if (help.get(0).getNext().get(0) != null && help.get(0).getNext().get(0).getKey().equals((K) k)) {
            V buf = (V) help.get(0).getNext().get(0).a.getValue();
            help.get(0).getNext().get(0).a.setValue(v);
            return buf;
        } else {
            V n = get(k);
            int r = rand.nextInt() << 1, i = 0;
            for (; i < levels && r % 2 == 0; i++, r /= 2) {
                x.getNext().set(i, help.get(i).getNext().get(i));
                help.get(i).getNext().set(i, x);
            }
            while (i < levels) {
                x.getNext().set(i, null);
                i++;
            }
            size++;
            return n;
        }
    }

    public V get(Object k) {
        ArrayList<Node> help;
        Node y = head;
        ArrayList<Node> help1 = new ArrayList<>();
        for (int i = 0; i < levels; i++) {
            help1.add(null);
        }
        for (int i = levels - 1; i >= 0; i--) {
            while (y.getNext().get(i) != null && y.getNext().get(i).getKey().compareTo((K) k) < 0)
                y = y.getNext().get(i);
            help1.set(i, y);
        }
        help = help1;
        Node x = help.get(0).getNext().get(0);
        if (x != null && x.getKey().equals((K) k))
            return x.getValue();
        else
            return null;
    }

    public V remove(Object k) {
        ArrayList<Node> help;
        Node y = head;
        ArrayList<Node> help1 = new ArrayList<>();
        for (int i = 0; i < levels; i++) {
            help1.add(null);
        }
        for (int i = levels - 1; i >= 0; i--) {
            while (y.getNext().get(i) != null && y.getNext().get(i).getKey().compareTo((K) k) < 0)
                y = y.getNext().get(i);
            help1.set(i, y);
        }
        help = help1;
        Node x = help.get(0).getNext().get(0);
        if (x == null || !x.getKey().equals((K) k))
            return null;
        for (int i = 0; i < levels && help.get(i).getNext().get(i) == x; i++)
            help.get(i).getNext().set(i, x.getNext().get(i));
        size--;
        return (V) x.getValue();
    }
    
      private final Random rand = new Random();
    private Node head;
    private int size = 0;
    private int levels;


    private class Node implements Map.Entry<K, V> {
        ArrayList<Node> next;

        Map.Entry a;
        public Node(Map.Entry a, int levels) {
            this.a = a;
            next = new ArrayList<>();
            for (int i = 0; i < levels; i++) {
                next.add(null);
            }
        }

        public ArrayList<Node> getNext() {
            return this.next;
        }

        @Override
        public K getKey() {
            return (K) a.getKey();
        }

        @Override
        public V getValue() {
            return (V) a.getValue();
        }

        @Override
        public V setValue(V value) {
            a.setValue(value);
            return null;
        }
    }

    private class NodeSet<K, V> extends AbstractSet {
        private int size;

        public NodeSet(int size) {
            this.size = size;
        }

        public Iterator iterator() {
            return new NodeIterator();
        }


        public int size() {
            return size;
        }

        private class NodeIterator implements Iterator {
            private Node list;

            public NodeIterator() {
                list = SkipList.this.head;
            }

            public boolean hasNext() {
                return list.getNext().get(0) != null;
            }

            public Node next() {
                list = list.getNext().get(0);
                return list;
            }

            public void remove() {
                SkipList.this.remove(list.getKey());
            }
        }
    }
}
