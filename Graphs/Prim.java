import java.util.*;

public class Prim {
    private static final int INF = 100000;
    private static ArrayList<edge>[] g;
    static class edge implements Comparable<edge>{
        private int lenght;

        private int b;

        public edge (int y, int z){
            this.b=y;
            this.lenght=z;
        }

        public int compareTo(edge q) {
            return lenght < q.lenght ? -1 : lenght > q.lenght? 1 : 0;
        }

    }
    
    public static long mst(ArrayList<edge>[] edges, int[] pred) {
        int n = edges.length;
        Arrays.fill(pred, -1);
        boolean[] vis = new boolean[n];
        int[] prio = new int[n];
        Arrays.fill(prio, INF);
        prio[0] = 1;
        Queue<edge> q = new PriorityQueue<edge>();
        q.add(new edge(0, 0));
        long res = 0;
        while (!q.isEmpty()) {
            edge cur = q.poll();
            int u = cur.b;
            if (vis[u])
                continue;

            vis[u] = true;
            res += cur.lenght;
            for (edge e : edges[u]) {
                int v = e.b;
                if (!vis[v] && prio[v]>e.lenght)
                {
                    prio[v] = e.lenght;
                    pred[v] = u;
                    q.add(new edge(v,prio[v]));
                }
            }
        }
        return res;
    }

    public static void main(String[] argc){
        Scanner in = new Scanner(System.in);
        int n = in.nextInt(), m = in.nextInt();
        g = new ArrayList[n];
        int[] min_e= new int[n];
        int[] sel_e= new int[n];
        int count=0;
        Queue<edge> q = new PriorityQueue<edge>();
        q.add(new edge(0,0));

        for(int i=0;i<n;i++){
            g[i]=new ArrayList<>();
            min_e[i]=INF;
            sel_e[i]=-1;
        }

        for(int i=0;i<m;i++){
            int x=in.nextInt();
            int y=in.nextInt();
            int z=in.nextInt();

            edge buf = new edge(y,z);
            g[x].add(buf);
            buf = new edge(x,z);
            g[y].add(buf);
        }
        count = (int) mst(g,new int[n]);
        System.out.println(count);

    }
}

