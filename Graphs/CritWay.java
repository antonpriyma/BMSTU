import java.util.*;

public class Cpm {

    static class vertex{
        private String name;
        private int time;

        public vertex(String s){
            s = s.replace("("," ");
            s = s.replace(")"," ");
            String[] help = s.split(" ");
            try {
                this.time = Integer.parseInt(help[1]);
            }catch (Exception e){
                this.time=counter.get(help[0]);
            }
            this.name=help[0];
        }

        public int getTime() {
            return time;
        }

        public String getName() {
            return name;
        }
    }


    static class graph{
        public HashMap<String,ArrayList<vertex>> data;

        public graph(){
            data = new HashMap<>();
        }

        public  HashMap<String, ArrayList<vertex>> getData() {
            return data;
        }
    }

    static class edge{
        private String a;
        private String b;
        private int cost;

        public edge(String a,String b, int cost){
            this.a=a;
            this.b=b;
            this.cost=cost;
        }

        public String getA(){
            return this.a;
        }

        public String getB(){
            return this.b;
        }
        public int getcost(){
            return this.cost;
        }

        public void setA(String a){
            this.a=a;
        }


        public void setCost(int cost){
            this.cost=cost;
        }
    }

    private static graph graph1;
    private static String[] keyset;
    private static ArrayList<edge> vert;
    private static HashSet<String> needCheck= new HashSet<>();
    private static String[] keys;
    private static ArrayList<String> checkEd= new ArrayList<>();
    private static ArrayList<String> comp= new ArrayList<>();
    private static HashMap<String,Integer> color= new HashMap<>();
    private static HashMap<String,Integer> color1= new HashMap<>();
    private static HashMap<String,Integer> counter= new HashMap<>();
    private static HashMap<String,Integer> d=new HashMap<>();
    private static HashSet<String> cycle=new HashSet<>();
    private static HashMap<String,HashSet<String>> p=new HashMap<>();
    private static int n;
    private static int m;
    private static int min=0;
    private static final int INF = 1000000;

    public static void dfc(String v){
        color1.put(v,0);
        if (!graph1.getData().containsKey(v)) {
            color1.put(v,1);
            return;
        }
        for (vertex i: graph1.getData().get(v)){
            if (color1.get(i.name)==-1){
                dfc(i.name);
            }else if(color1.get(i.name)==0){
                cycle.add(v);
                cycle.add(i.name);
            }
        }
        color1.put(v,1);
    }

    public static void setcolor(String s){
        if (!p.containsKey(s))
            return;
        for (String i: p.get(s)){
            color.put(i,1);
            setcolor(i);
        }
    }

    public static void setcolor1(String s){
        if (!graph1.getData().containsKey(s))
            return;
        for (vertex i: graph1.getData().get(s)){
            color.put(i.name,-1);
            if (!cycle.contains(i.name))
            setcolor1(i.name);
        }
    }

    public static  void solve() {

        ArrayList<Integer> t = new ArrayList<>();
        String val="";keys = new String[d.size()];
        d.keySet().toArray(keys);

        for (String i: keys){
            if (d.get(i)==0)
                val=i;
        }
        for (;;) {
            boolean any = false;
            for (int j = 0; j < m; ++j) {
                if (color.get(vert.get(j).a)==-1){
                    continue;
                }
                if (d.get(vert.get(j).a) < INF && !checkEd.contains(vert.get(j).a)) {
                    if (d.get(vert.get(j).b) > d.get(vert.get(j).a) + vert.get(j).cost) {
                        d.put(vert.get(j).b, d.get(vert.get(j).a) + vert.get(j).cost);
                        if (p.containsKey(vert.get(j).b)) {
                            p.get(vert.get(j).b).clear();
                            p.get(vert.get(j).b).add(vert.get(j).a);
                        } else {
                            p.put(vert.get(j).b, new HashSet<>());
                            p.get(vert.get(j).b).add(vert.get(j).a);
                        }
                        any = true;
                    } else if (d.get(vert.get(j).b) == d.get(vert.get(j).a) + vert.get(j).cost) {
                        if (p.containsKey(vert.get(j).b)) {
                            ;
                            p.get(vert.get(j).b).add(vert.get(j).a);
                        } else {
                            p.put(vert.get(j).b, new HashSet<>());
                            p.get(vert.get(j).b).add(vert.get(j).a);
                        }
                    }
                }
            }
            if (!any) break;
        }



        for (int i = 0; i < d.size(); i++) {
            if (m==0){
                if (-counter.get(keys[i]) < min  && color.get(keys[i])!=-1) {
                    min = -counter.get(keys[i]);
                    t.clear();
                    t.add(i);
                }else if (-counter.get(keys[i])==min && color.get(keys[i])!=-1){
                    t.add(i);
                }
            }else {
                if (d.get(keys[i])  < min && color.get(keys[i]) != -1) {
                    min = d.get(keys[i]);
                    for (String j: keys){
                        if (color.get(j)==1)
                        color.put(j,0);
                    }
                    t.clear();
                    t.add(i);
                } else if (d.get(keys[i]) == min && color.get(keys[i]) != -1) {
                    t.add(i);
                }
            }
        }
        if (m==0){
            for (int i: t) {
                color.put(keys[i], 1);
            }
        }else {
            for (int i: t) {
                color.put(keys[i],1);
                setcolor(keys[i]);
            }

        }




    }

    public static void main(String[] argc){

        graph1= new graph();
        vert= new ArrayList<>();
        Scanner in = new Scanner(System.in);
        StringBuilder s = new StringBuilder();
        while (in.hasNextLine()){
            s.append(in.nextLine());
        }

        String[] help = s.toString().split(";");

        for (String i: help) {
            if (i.equals("")) i=s.toString();
            String[] helpBuild = i.split("<");
            vertex buf2 = new vertex(helpBuild[0]);
            if (!graph1.getData().containsKey(buf2.name)) {
                graph1.getData().put(buf2.name, new ArrayList<>());
            }
            //  graph1.getData().get(buf2.name).add(buf2);
            color.put(buf2.name,0);
            d.put(buf2.name,-buf2.time);
            color1.put(buf2.name,-1);

                counter.put(buf2.name,buf2.time);
            needCheck.add(buf2.name);


            for (int j = 0; j < helpBuild.length - 1; j++) {
                helpBuild[j]=helpBuild[j].replace(" ","");
                helpBuild[j+1]=helpBuild[j+1].replace(" ","");

                vertex buf = new vertex(helpBuild[j]);
                if (graph1.getData().containsKey(buf.name)) {
                    graph1.getData().get(buf.name).add(new vertex(helpBuild[j + 1]));
                } else {
                    graph1.getData().put(buf.name, new ArrayList<>());
                    graph1.getData().get(buf.name).add(new vertex(helpBuild[j+1]));
                    m++;
                }
                if (buf.time != -1) {
                    counter.put(buf.name,buf.time);
                }

                d.put(buf.name,INF);
                vertex buf1 = new vertex(helpBuild[j+1]);
                color.put(buf.name,0);
                if (buf1.time != -1) {
                    counter.put(buf1.name,buf1.time);
                }
                needCheck.add(buf1.name);
                color1.put(buf1.name,-1);
                d.put(buf1.name,INF);
                color.put(buf1.name,0);
                boolean flag = false;
                for (edge k: vert){
                    if (buf.name.equals(k.a) && buf1.name.equals(k.b)) {
                        flag=true;
                    }
                }
                if (!flag)
                vert.add(new edge(buf.name,buf1.name,-counter.get(buf1.name)));
            }
        }
        n=counter.size();
        keyset=new String[n];
        graph1.getData().keySet().toArray(keyset);


        m=vert.size();
        for (String i: needCheck)
            if (color1.get(i)!=1) {
                comp.add(i);
                dfc(i);
            }

        for (String i:cycle){
            color.put(i,-1);
            setcolor1(i);
        }
        System.out.println("digraph {");
        for (String x: comp){
             d.put(x,-counter.get(x));
            solve();
            checkEd.add(x);
        }
        for (int i=0;i<n;i++){
            if (color.get(keys[i])==0) {
                System.out.println("\t" + keys[i] + " [label = \"" + keys[i] + "(" + counter.get(keys[i]) + ")\"]");
            }
            else if (color.get(keys[i])==1){
                System.out.println("\t" + keys[i] + " [label = \"" + keys[i] + "(" + counter.get(keys[i]) + ")\", color = red]");
            }else if (color.get(keys[i])==-1){
                System.out.println("\t" + keys[i] + " [label = \"" + keys[i] + "(" + counter.get(keys[i]) + ")\", color = blue]");
            }
        }
        for (int i=0;i<m;i++){
            if (color.get(vert.get(i).a)==1 && color.get(vert.get(i).b)==1 && p.get(vert.get(i).b).contains(vert.get(i).a)){
                System.out.println("\t"+vert.get(i).a+" -> "+vert.get(i).b+" [color = red]");
            }else  if (color.get(vert.get(i).a)==-1 && color.get(vert.get(i).b)==-1){
                System.out.println("\t"+vert.get(i).a+" -> "+vert.get(i).b+" [color=blue]");
            }else{
                System.out.println("\t"+vert.get(i).a+" -> "+vert.get(i).b);
            }
        }
        System.out.println("}");
    }




}

