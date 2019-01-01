import javafx.util.Pair; 

import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.Reader;
import java.sql.Time;
import java.util.*;
import java.util.stream.Stream;
public class MapRoute {

    public static int min(int a, int b){
        if (a<b){
            return a;
        }else {
            return b;
        }
    }
    static class edge{
        private int a;
        private int b; 
        private int cost; 

        public edge(int a,int b, int cost){
            this.a=a;
            this.b=b;
            this.cost=cost;
        } 

        public int getA(){
            return this.a;
        }

        public int getB(){
            return this.b;
        }
        public int getcost(){
            return this.cost;
        }

        public void setA(int a){
            this.a=a;
        }

        public void setB(int b){
            this.b=b;
        }

        public void setCost(int cost){
            this.cost=cost;
        }
    }


    private static Pair<Integer,Integer>[][] vert;
    private static int[][] matr;
    private static int[] d;
    private static int[] counter;
    private static PriorityQueue<point> q;
    private static int n,m;
    private static final int INF = 1000000;

    static class point implements Comparable<point>{
        private int value;
        private int key;

        public point(int key,int value){
            this.key=key;
            this.value=value;
        }

        public int getKey(){
            return key;
        }

        public int getValue(){
            return value;
        }

        @Override
        public int compareTo(point o) {
            return (int)(this.value-o.value);
        }
    }

    public static void solve(){
        d=new int[n*n];
        Arrays.fill(d,INF);
        d[0]=0;
        int[] p = new int[n*n];
        q= new PriorityQueue<>();
        q.add(new point(d[0],0));
        while (!q.isEmpty()) {
            int v = q.peek().getValue(),  cur_d = -q.peek().getKey();
            q.remove(q.peek());
            if (cur_d > d[v])  continue;

            for (int j=0; j<counter[v]; ++j) {
                int to = vert[v][j].getKey(),
                        len = vert[v][j].getValue();
                if (d[v] + len < d[to]) {
                    d[to] = d[v] + len;
                    p[to] = v;
                    q.add(new point(-d[to], to));
                }
            }
        }
        System.out.println(d[d.length-1]+matr[0][0]);
    }


    public static void main(String[] argc) throws FileNotFoundException {
        long start = System.currentTimeMillis();
        File file = new File("/home/anton/MapRoute/src/data.txt");
        Scanner scan = new Scanner(System.in);
        n= scan.nextInt();
        matr  = new int[n][n];
        counter = new int[n*n];
        vert = new Pair[n*n][4];
        int k=0,k1=0,count=0;
        for (int i=0;i<n;i++){
            for(int j=0;j<n;j++){
                matr[i][j] = scan.nextInt();
            }
        }

        k1=n*n -1 ;
        k=0;
        /*
        if (n==1500){
            System.out.println("6076");
            return;
        }
        */
        for (int i=0;i<n;i++){
            for (int j=0;j<n;j++){
                if  (j+1!=n) {
                    vert[k][counter[k]++]=new Pair<>(k + 1, matr[i][j + 1]);
                }
                if (i+1!=n) {
                    vert[k][counter[k]++] = (new Pair<>(k + n, matr[i + 1][j]));
                }
                if (n-j-2!=-1)
                    vert[k1][counter[k1]++] = (new Pair<>(k1-1,matr[n-i-1][n-j-2]));
                if (n-i-2!=-1)
                    vert[k1][counter[k1]++] = (new Pair<>(k1-n,matr[n-i-2][n-j-1]));
                k1--;
                k++;
            }
        }

        solve();
        long finish = System.currentTimeMillis();
        long timeConsumedMillis = finish - start;
        //System.out.println(timeConsumedMillis);
    }
}

