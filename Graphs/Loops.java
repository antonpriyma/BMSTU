import com.sun.corba.se.impl.orbutil.graph.Graph;

import java.util.*;

public class Loops {

    private static int count = 0;

    static class Vert implements Comparable<Vert> {
        private int value;
        private int result;
        private Vert sdom=this;
        private boolean used = false;
        private boolean isResultable;
        private Vert dom=null;
        private Vert label=this;
        private Vert parent=null;
        private Vert ancector=null;
        int n;
        private ArrayList<Vert> bucket;
        private ArrayList<Vert> a;
       private ArrayList<Vert> b;


        public boolean isResultable() {
            return isResultable;
        }

        public boolean isUsed() {
            return used;
        }

        public int getValue() {
            return value;
        }

        public int getResult() {
            return result;
        }

        public Vert getSdom() {
            return sdom;
        }

        public ArrayList<Vert> getA() {
            return a;
        }
        
        public ArrayList<Vert> getB() {
            return b;
        }

        public ArrayList<Vert> getBucket() {
            return bucket;
        }

        public int getN() {
            return n;
        }

        public Vert getAncector() {
            return ancector;
        }

        public Vert getDom() {
            return dom;
        }

        public Vert getLabel() {
            return label;
        }

        public Vert getParent() {
            return parent;
        }

        public void setA(ArrayList<Vert> a) {
            this.a = a;
        }

        public void setAncector(Vert ancector) {
            this.ancector = ancector;
        }

        public void setB(ArrayList<Vert> b) {
            this.b = b;
        }

        public void setBucket(ArrayList<Vert> bucket) {
            this.bucket = bucket;
        }

        public void setDom(Vert dom) {
            this.dom = dom;
        }

        public void setLabel(Vert label) {
            this.label = label;
        }

        public void setN(int n) {
            this.n = n;
        }

        public void setParent(Vert parent) {
            this.parent = parent;
        }

        public void setResult(int result) {
            this.result = result;
        }

        public void setSdom(Vert sdom) {
            this.sdom = sdom;
        }

        public void setValue(int value) {
            this.value = value;
        }

        public void setUsed(boolean used) {
            this.used = used;
        }

        public void setResultable(boolean resultable) {
            isResultable = resultable;
        }

        @Override
        public int compareTo(Vert graph) {
            return (this.n - graph.n);
        }

        public int vecComp(Vert o, Vert o1){
            return o.compareTo(o1);
        }
    }

    class VecComp implements Comparator<Vert>{
        public int compare(Vert a, Vert b){
            return a.compareTo(b);
        }
    }


    public static final Vert FindMin(ArrayList<Vert> g, Vert a) {
        SearchAndCut(g, a);
        return a.label;
    }

    public static void dfs(ArrayList<Vert> g) {
        dfsSecond(g.get(0));
        for (int i = 0; i < g.size(); i++) {
            if (!g.get(i).isUsed()) {
                g.remove(i--);
            } else {
                for (int j = 0; j < g.get(i).getB().size(); j++) {
                    if (!g.get(i).b.get(j).used) {
                        g.get(i).b.remove(j);
                        j--;
                    }
                }
            }
        }
    }

    public static void dfsSecond(Vert v) {
        v.setN(Loops.count);
        Loops.count++;
        v.setUsed(true);
        for (Vert i : v.getA()) {
            if (!i.isUsed()) {
                i.setParent(v);
                dfsSecond(i);
            }
        }

    }


    public static final Vert SearchAndCut(ArrayList<Vert> g, Vert a) {
        Vert root = new Vert();
        if (a.getAncector() == null) {
            root = a;
        } else {
            root=SearchAndCut(g,a.getAncector());
            if (a.getAncector().getLabel().getSdom().getN() < a.getLabel().getSdom().getN()) {
                a.setLabel(a.getAncector().getLabel());
            }
            a.setAncector(root);
        }

        return root;
    }

    public final static void Dominators(ArrayList<Vert> g) {
        int n = g.size() - 1;
        for (int i = n; i > 0; i--) {
            Vert w = g.get(i);
            for (Vert v : w.b) {
                Vert u = FindMin(g, v);
                if (u.getSdom().getN() < w.getSdom().getN())
                    w.setSdom(u.getSdom());
            }
            w.setAncector(w.parent);
            w.getSdom().getBucket().add(w);
            for (Vert v : w.getParent().getBucket()) {
                Vert u = FindMin(g, v);
                if (u.getSdom() == v.getSdom()) {
                    v.setDom(w.parent);
                } else v.setDom(u);
            }
            w.getParent().getBucket().clear();
        }
        n++;
        for (int i = 1; i < n; i++) {
            Vert w = g.get(i);
            if (w.getDom() != w.getSdom())
                w.setDom(w.getDom().getDom());
        }
        g.get(0).setDom(null);
    }


    public static int FindLoops(ArrayList<Vert> v) {
        int loops = 0;
        for (Vert i : v) {
            for (Vert j : i.getB()) {
                while (j != i && j != null)
                    j=j.dom;
                if (j==i) {
                    loops++;
                    break;
                }
            }
        }
        return loops;
    }

    public static void main(String[] argc) {
        Scanner in = new Scanner(System.in);

        int n = in.nextInt(), N = 0;
        ArrayList<Vert> graph = new ArrayList<>();
        ArrayList<Integer> op = new ArrayList<>();
        ArrayList<Integer> mark = new ArrayList<>();
        for (int i = 0; i < n; i++){
            graph.add(new Vert());
            graph.get(i).setA(new ArrayList<>());
            graph.get(i).setB(new ArrayList<>());
            graph.get(i).setBucket(new ArrayList<>());
        }
        in.nextLine();

        Map<Integer, Integer> g = new TreeMap<Integer, Integer>(Integer::compareTo);

        for (int i = 0; i < n; i++) {
            String buf[] = in.nextLine().split(" ");
            graph.get(i).setValue(Integer.parseInt(buf[0]));
            g.put(graph.get(i).getValue(), i);
            if ((buf[1].equals("ACTION") || buf[1].equals("BRANCH")) && i < n - 1) {
                graph.get(i).getA().add(graph.get(i + 1));
                graph.get(i + 1).b.add(graph.get(i));
            }
            if (buf.length != 2) {
                graph.get(i).setResult(Integer.parseInt(buf[2]));
                graph.get(i).setResultable(true);
            } else {
                graph.get(i).setResultable(false);
                graph.get(i).setResult(1);
            }
        }
        for (Vert v : graph) {
            if (v.isResultable() && v.getResult() != -1) {
                v.getA().add(graph.get(g.get(v.getResult())));
                graph.get(g.get(v.getResult())).getB().add(v);
            }
        }
        dfs(graph);
        Collections.sort(graph,Vert::compareTo);
        Dominators(graph);
        System.out.println(FindLoops(graph));

    }

}



