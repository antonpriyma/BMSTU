import java.util.*;

public class Kruskal{
    private static ArrayList<edge> g = new ArrayList<>();
    private static ArrayList<vertex> g1=new ArrayList<>();
    //private static ArrayList<Map<Integer,Integer>> g1 = new ArrayList<>();
 

    private static void radixsort(int[] arr, int[] ch) {
        int[]counter=new int[65536];
        int[]helper=new int[arr.length];
        for(int i=0;i<arr.length;i++)
            counter[ch[arr[i]] & 0x0000ffff]++;
        for(int i=1;i<counter.length;i++) {
            counter[i]+=counter[i-1];
        } 
        for(int i=arr.length-1;i>=0;i--) { 
            helper[ --counter[ch[arr[i]]&0x0000ffff]]=arr[i];
        }
        System.arraycopy(helper,0,arr,0,helper.length);
        counter=new int[32768];
        for(int i =0;i<helper.length;i++)
            counter[ch[helper[i]]>> 16]++;
        for(int i=1;i<counter.length;i++){
            counter[i]+=counter[i-1];
        }
        for(int i=helper.length-1;i>=0;i--){
            arr[--counter[ch[helper[i]]>>16]]=helper[i];
        }
    }

    static class vertex{
        private int x,y;



        public vertex(int a,int b){
            this.x=a;
            this.y=b;
        }


    }

    static class edge  {
        private int length;
        int a;
        int b;

        public edge(int a, int b) {
            this.a = a;
            this.b = b;
            this.length = ((g1.get(a).x - g1.get(b).x)*(g1.get(a).x - g1.get(b).x) + (g1.get(a).y - g1.get(b).y)*(g1.get(a).y - g1.get(b).y)) ;
        }

    }



    public static void main(String[] argc){
        //long start = System.currentTimeMillis();
        //long finish;
        Scanner in = new Scanner(System.in);
        int n = in.nextInt(),vertex1,vertex2;
        int[] tree_id= new int[n];
        int arr[]  = new int[(n*(n-1))/2];
        for (int i=0;i<n;i++){
            tree_id[i]=i;
            vertex1=in.nextInt();vertex2=in.nextInt();
            g1.add(new vertex(vertex1,vertex2));
        }
        int count1=0;
        int ch[]=new int[(n*(n-1)/2)];
        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                arr[count1]=count1++;
                g.add(new edge(i,j));
            }
        }

        for(int i=0;i<g.size();i++){
            ch[i]=g.get(i).length;
        }

        float res=0;
        radixsort(arr,ch);

        int counter=0,e=0;
        double weight= (double)0;
        while (counter<tree_id.length-1) {
            int u1, u2;
            u1=g.get(arr[e]).a;
            u2=g.get(arr[e]).b;
            if(tree_id[u1]!=tree_id[u2]) {
                int mem=tree_id[u2];
                counter+=1;
                weight+=Math.sqrt(g.get(arr[e]).length);
                for(int i=0; i<tree_id.length; i++) if (tree_id[i]==mem) tree_id[i]=tree_id[u1];
            }
            e+=1;
        }

        System.out.printf("%.2f\n",weight);
        //finish=System.currentTimeMillis()-start;
        //System.out.println(finish+"mc");
    }
}

