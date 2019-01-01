import java.util.ArrayList;
import java.util.Scanner;

public class BridgeNum {
    private static ArrayList<Integer>[] g;
    private static boolean[] used;
    private static int timer,tim[],fup[];
    public static int count=0;

    private static int min(int a,int b){
    if (a<b) return a;
    return b;
    }


    public static void dfs(int v,int p) {
        used[v] = true;
        tim[v] = fup[v] = timer++;
        for (int i = 0; i < g[v].size(); ++i) {
            int to = g[v].get(i);
            if (to == p) continue;
            if (used[to])
                fup[v] = min(fup[v], tim[to]);
            else {
                dfs(to, v);
                fup[v] = min(fup[v], fup[to]);
                if (fup[to] > tim[v])
                    count++;
            }
        }
    }

    public static void main(String[] argc){
        Scanner scan  = new Scanner(System.in);
        int n,m;
        n=scan.nextInt();
        m=scan.nextInt();
        used= new boolean[n];
        g= new ArrayList[n];
        tim=new int[n];
        fup=new int[n];

        for (int i=0;i<n;i++){
            g[i]=new ArrayList<>();
        }

        for (int i=0;i<m;i++){
            int x,x1;
            x= scan.nextInt(); x1 = scan.nextInt();
            g[x].add(x1);
            g[x1].add(x);
        }

        int timer =0;
        for (int i=0;i<n;++i){
            if (!used[i]){
                dfs(i,-1);
            }
        }

        System.out.println(count);
    }
}

