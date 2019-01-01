import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;

public class EqDist {
    private static ArrayList<Integer>[] g;
    private static boolean[] used;
    private static int timer,tim[],fup[];
    private static int count=0;
    private static int[] p,comp;
    private static HashMap<Integer,HashMap<Integer,Integer>> s = new HashMap<>();
    private static boolean[] equal;
    private static StringBuilder result = new StringBuilder();

    public static void main(String[] argc){
        Scanner scan  = new Scanner(System.in);
        int n,m;
        n=scan.nextInt();
        m=scan.nextInt();
        used= new boolean[n];
        g= new ArrayList[n];
        tim=new int[n];
        fup=new int[n];
        p=new int[n];
        equal=new boolean[n];
        int[] vertex;
        int[] p = new int[n];
        comp = new int[n];
        int[] u = new int[n];



        for (int i=0;i<n;i++){
            g[i]=new ArrayList<>();
            s.put(i,new HashMap<>());
            equal[i]=true;
        }

        for (int i=0;i<m;i++){
            int x,x1;
            x= scan.nextInt(); x1 = scan.nextInt();
            g[x].add(x1);
            g[x1].add(x);
        }

        m=scan.nextInt();
        vertex=new int[m];
        for (int i=0;i<m;i++){
            vertex[i]=scan.nextInt();
            equal[vertex[i]]=false;
        }
        int[][] d = new int[n][n];

        for (int i=0;i<n;i++){
            p[i]=100000;
            for(int j=0;j<n;j++){
                d[i][j]=100000;
            }
            d[i][i]=0;
        }
        
        for(int k=0;k<m;k++) {
            for (int i=0;i<n;i++){
                u[i]=-1;
                p[i]=0;
            }
            for (int i = 0; i < n; ++i) {
                int v = -1;
                for (int j = 0; j < n; ++j) {
                    if (v == -1 && u[j] < 0) {
                        v = j;
                    }
                    if (u[j] < 0 && (v == -1 || d[vertex[k]][j] < d[vertex[k]][v]))
                        v = j;
                }
                if (d[vertex[k]][v] == 100000)
                    break;
                u[v] = 1;

                for (int j = 0; j < g[v].size(); ++j) {
                    int to = g[v].get(j),
                            len = 1;
                    if (to==v) continue;
                    if (d[vertex[k]][v] + len < d[vertex[k]][to]) {
                        d[vertex[k]][to] = d[vertex[k]][v] + len;
                        p[to] = v;
                    }
                }
            }
        }
        
        int component = comp[vertex[0]];

        for (int i =0;i<m;i++){
            if (component!=comp[vertex[i]]){
                System.out.println("-");
                System.exit(0);
            }
        }
      

        for (int i=0;i<n;i++){
            if(comp[vertex[0]]==comp[i] && i!=vertex[0] && d[vertex[0]][i]!=100000) {
                p[i] = d[vertex[0]][i];
            }
            else {
                equal[i]=false;
            }
        }

        for(int i=1;i<m;i++){
            for(int j=0;j<n;j++){
                if(p[j]!=-1) {
                    if (j != vertex[i]) {
                        if (p[j] != d[vertex[i]][j]) {
                            equal[j] = false;

                        }
                    }
                }
            }
        }
        count=0;

        for (int i=0;i<n;i++) {
            if (equal[i]) {
                result.append(i+" ");
            } else {
                count++;
            }
        }
        if (count==n){
            System.out.println("-");
            System.exit(0);
        }
        System.out.println(result);
    }
}


