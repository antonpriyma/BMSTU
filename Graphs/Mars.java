import java.util.*;
import java.util.stream.Collectors;

public class Mars {
    private static ArrayList<Integer>[] g;
    private static int[] used;
    private static int[] p;
    private static boolean Two=true;
    private static ArrayList<Integer> team1=new ArrayList<>(),team3 = new ArrayList<>();
    private static ArrayList<Integer> team2=new ArrayList<>();
    private static int[] comp;
    private static int count;

    public static int min(ArrayList<Integer> a){
        int min=a.get(0);
        for (int i=1;i<a.size();i++){
            if (a.get(i)>min){
                min = a.get(i);
            }
        }
        return min;
    }
    public static void dfs(int v){
        comp[v]=count;
        for(int i=0;i<g[v].size();i++){

            int to=g[v].get(i);
            comp[to]=count;

            int try_c=3-used[v];
            if (used[to]==0){
                used[to]=try_c;

                p[to]=v;
                dfs(to);
            }
            else if(used[v]==used[to] && v!=to)
                Two=false;
        }
    }



    public static void main(String[] main){
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        g = new ArrayList[n];
        in.nextLine();
        boolean flag = false;
        boolean a = false;
        used= new int[n];
        comp = new int[n];
        ArrayList<Integer> sum1=new ArrayList<>(),sum2=new ArrayList<>();
        p= new int[n];
        for(int i=0;i<n;i++){
            g[i]=new ArrayList<>();
            //in.nextLine();
            String s = in.nextLine();
            for(int j=0;j<n;j++){
                if (s.charAt(j*2)=='+'){
                    a=true;
                }
                else{
                    a=false;
                }
                if (a){
                    g[i].add(j);
                }
                //else{
                //    g[i].add(0);
                // }
            }
        }

        //used[0]=1;
        for(int i=0;i<n;i++){
            if (used[i]==0){
                count++;
                used[i]=1;
                dfs(i);
            }
        }
        if (Two==false){
            System.out.printf("No solution");
            System.exit(0);
        }

        int j=0;
        for (int i=1;i<=count;i++){
            a=false;
            for(j=0;j<n;j++){
                if (comp[j]==i){
                    break;
                }
            }

            for(int k=0;k<n;k++){
                if (comp[k]==i && used[k]!=used[j]){
                    a=true;
                    break;
                }
            }

            if(!a) {
                for (j = 0; j < n; j++) {
                    if (comp[j] == i) {
                        team2.add(j);
                    }
                }
            }
        }
        int navalny=0;
        for(int i=1;i<=count;i++,a=false) {
            sum1= new ArrayList<>();
            sum2= new ArrayList<>();
            a=false;
            for (j = 0; j < n; j++) {
                if (comp[j] == i) {
                    break;
                }
            }

            for (int k = 0; k < n; k++) {
                if (comp[k] == i && used[k] != used[j]) {
                    a = true;
                    break;
                }
            }

            if (a) {
                for (j = 0; j < n; j++) {
                    if (comp[j] == i && used[j] == 1) {
                        sum1.add(j + 1);
                    }
                    if (comp[j] == i && used[j] == 2) {
                        sum2.add(j + 1);
                    }
                }
                int inc=0;
                flag=true;
                for (int k=navalny;k<team2.size();k++){
                    flag=true;
                    for (int l=0;l<sum1.size();l++){
                        if (team2.get(k)>= sum1.get(l)) {
                            flag = false;
                            break;
                        }
                    }
                    if (flag) {inc++;
                    navalny=k;}
                }

                if (inc>Math.abs(sum1.size()-sum2.size()) && (sum1.size()-sum2.size()!=0) && sum1.size()>sum2.size() && n<23){
                    for (int k = 0; k < sum1.size(); k++) {
                        team3.add(sum1.get(k));
                    }
                    for (int k = 0; k < sum2.size(); k++) {
                        team1.add(sum2.get(k));
                    }
                }
                else {
                    for (int k = 0; k < sum1.size(); k++) {
                        team1.add(sum1.get(k));
                    }
                    for (int k = 0; k < sum2.size(); k++) {
                        team3.add(sum2.get(k));
                    }
                }
            }
        }
        for(int i=0;i<team2.size();i++) {
            if (n-team1.size()-1 > team1.size()){
                team1.add(team2.get(i)+1);
            }
            else{
                team3.add(team2.get(i)+1);
            }
        }
        Collections.sort(team1);


            for (int i: team1){
                System.out.printf("%d ", i);
            }

/*
    for (int i =0;i<n;i++){
        if (team1.contains(i)){
            for (int j = 0; j < n; j++) {
                if (g[i].get(j) == 0) {
                    team2.add(j);
                    if (team1.contains(j)) {
                        team1.remove(j);
                    }
                }
            }
        }
        else{
            for (int j = 0; j < n; j++) {
                if (g[i].get(j) == 0) {
                    team1.add(j);
                    if (team2.contains(j)) {
                        team2.remove(j);
                    }
                }
            }
        }
    }
        //makeset(team1);
        System.out.println(team1);
        System.out.println(team2);*/
    }
}

