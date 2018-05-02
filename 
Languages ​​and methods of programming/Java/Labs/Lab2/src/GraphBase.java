import javafx.util.Pair;

import java.util.*;

public class GraphBase {
    private static ArrayList<Integer>[] g;
    private static ArrayList<Integer>[] gr;
    private static boolean[] used;
    private static int[] color;
    private static ArrayList<Pair<Integer,Integer>> s;
    private static ArrayList<Integer> order, component;
    private static boolean[] base;
    public static void dfs1 (int v) {
        used[v] = true;
        for (int  i=0; i<g[v].size(); ++i)
            if (!used[ g[v].get(i) ])
                dfs1 (g[v].get(i));
        order.add(v);
    }

    public static void dfs2 (int v,int c) {
        color[v] = c;
        component.add(v);
        for (int i=0; i<gr[v].size(); ++i)
            if (color[ gr[v].get(i) ]==-1)
                dfs2 (gr[v].get(i),c);
    }


    public static void main(String[] argc) {
        int n = 0, m = 0, x = 0, x1 = 0;
        Scanner in = new Scanner(System.in);
        n=in.nextInt();
        m=in.nextInt();
        s= new ArrayList<>();
        g = new ArrayList[n];
        gr = new ArrayList[n];
        order=new ArrayList<>();
        component=new ArrayList<>();
        used = new boolean[n];
        color = new int[n];
        Arrays.fill(used, false);

        for (int i = 0; i < n; i++) {
            color[i]=-1;
            g[i] = new ArrayList<>();
            gr[i] = new ArrayList<>();
        }//иницилизация

        for (int i = 0; i < m; i++) {
            x = in.nextInt();
            x1 = in.nextInt();
            g[x].add(x1);
            gr[x1].add(x);
        }//считывание, создание списков смежности

        for (int i = 0; i < n; ++i)
            if (!used[i])
                dfs1(i);

        Arrays.fill(used, false);
        int c =0;
        for (int i = 0; i < n; ++i) {
            int v = order.get(n - i - 1);
            if (color[v]==-1) {
                dfs2(v,c++);
            }
        }
        base= new boolean[c];
        Set<Integer> help=new HashSet<>();

        for(int i = 0; i < g.length; i++)
            for(int j = 0; j < g[i].size(); j++) {
                int to = g[i].get(j);
                if (color[i]!=color[to]){
                    //Map <Integer,Integer> buf = new HashMap<>();
                    s.add(new Pair<>(i,to));
                    help.add(i);
                }
            }
        Arrays.fill(base,true);

        Integer[] help1= new Integer[n];
        help.toArray(help1);
        boolean flag =true;
        for (int i=0;i<s.size();i++){
            base[color[s.get(i).getValue()]]=false;//ставим все с ненулевой входящей false
        }

        /*
        for (int i=0;i<s.size();i++){
            int buf = help1[i];
            flag=true;
            for (int j=0;j<s.size();j++){
                if (s.get(help1[j])==buf){
                    flag = false;
                    base[color[buf]]=false;
                }
            }
           // if (flag==true) System.out.println(buf);
        }
        */
        flag=false;
        for(int i=base.length-1;i>=0;i--){
            if (base[i]==true){
                flag=true;
                for (int j=0;j<color.length;j++){
                    if (color[j]==i){
                        System.out.println(j);
                        break;
                    }
                }
            }
        }
        if (flag==false) System.out.println(0);

       // System.out.println(s);
    }
}
