import java.io.*;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Scanner;

public class MaxComponent {
    //private static ArrayList<Integer>[] g,o1;
    private static boolean[] used;
    private static ArrayList<Integer> used1;
    private static int[] comp;
    private static int[] o;
    private static ArrayList<Integer>[] o1,o2;
    private static int count=0;
    private static StringBuilder s = new StringBuilder();


    public static void dfs(int v){
        comp[v]=count;
        if (o1[v]!=null)
        for(int i=0;i<o1[v].size();i++){

            int to=o1[v].get(i);
            comp[to]=count;


            if (!used[to]){
                used[to]=true;
                if (o1[to]!=null)
                dfs(to);
            }
        }
    }

    public static void main(String[] argc) throws FileNotFoundException {

        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int x,x1;
        used=new boolean[n];
        used1=new ArrayList<>();
        comp=new int[n];
        int count1[],count2[];
        o =new int[n];
        o2 = new ArrayList[n];
        int max=0,max1=0,max2=0;
        //g = new ArrayList[n];
        o1=new ArrayList[n];
        int m = in.nextInt();
        if (m==0){
            System.out.println("graph {\n" +
                    "\t0 [color = red]");
            for (int i=1;i<n;i++){
                System.out.println("\t"+i);
            }
            System.out.println("}");
            System.exit(0);
        }


        for(int i=0;i<m;i++){
            //in.nextLine();
            //String s = in.nextLine();
            x=in.nextInt();
            x1=in.nextInt();
                if (o1[x] != null) {
                        o1[x].add(x1);
                        o[x]++;

                } else {
                        o1[x] = new ArrayList<>();
                        o1[x].add(x1);
                        o[x] = 1;
                }
            if (o1[x1] != null) {
                o1[x1].add(x);
                o[x1]++;

            } else {
                o1[x1] = new ArrayList<>();
                o1[x1].add(x);
                o[x1] = 1;
            }

            if (o2[x]!=null){
                    o2[x].add(x1);
            }
            else {
                    o2[x]=new ArrayList<>();
                    o2[x].add(x1);
            }



                //g[x].add(x1);

                //g[x1].add(x);

        }
        for(int i=0;i<n;i++){
            if (!used[i]){
                count++;
                used[i]=true;
                dfs(i);
            }
        }

        count1=new int[count];
        count2=new int[count];

        for (int  i =0;i<n;i++) {
            count1[comp[i] - 1]++;
            if (o[i] > 0) {
                count2[comp[i] - 1] += o[i];
            }
        }

        for(int i=0;i<n;i++){
            if (count1[comp[i]-1]>max1 || (count2[comp[i]-1]>max2 && count1[comp[i]-1]==max1)){
                max=comp[i]-1;
                max1=count1[comp[i]-1];
                max2=count2[comp[i]-1];
            }
        }

        System.out.println("graph {");

        for(int i=0;i<n;i++){

            if (comp[i]==max+1){
                s.append("\t"+i+"[color = red]"+ "\n");
               // for(int j=0;j<o[i];j++){
               //     System.out.println("\t"+i+" -- "+k.get(j)+"[color = red]");
                //}
            }
            else {
                s.append("\t"+i+ "\n");
                //for(int j=0;j<o[i];j++){
                //    System.out.println("\t"+i+" -- "+k.get(j));
               // }
            }
        }


        //System.out.println("программа выполнялась " + timeSpent + " миллисекунд");

        for (int i=0;i<n;i++){
            ArrayList<Integer> k = null;
            if (o2[i]!=null) {
                k = o2[i];
            }
            if (o2[i]!=null){
                if (comp[i]==max+1){
                    for(int j=0;j<k.size();j++){
                        s.append("\t"+i+" -- "+k.get(j)+"[color = red]" + "\n");
                    }
                }
                else {
                    for(int j=0;j<k.size();j++){
                        s.append("\t"+i+" -- "+k.get(j)+ "\n");
                    }
                }
            }
        }
        System.out.println(s);
        System.out.println("}");
       // long timeSpent = System.currentTimeMillis() - startTime;

    }

}

