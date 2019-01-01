import java.util.ArrayList;
import java.util.Scanner;

public class Dividers {

    static class DividersList extends ArrayList<Long>{


        @Override
        public String toString(){
            String s = "";
            for (int i=this.size()-1;i>0;i--){
                s = s + "\t" + this.get(i)+"\n";
            }
                s = s + "\t" + this.get(0);
            return s;
        }
    }


    public static void Dividers(DividersList arr, long start, long i, long j){
        int  sqrt = (int)Math.sqrt(start);
        arr.add(j);
        if (i!=j) {
            for (j++; j <= sqrt && start%j!=0; j++);
                if (j<=sqrt){
                    Dividers(arr,start,start/j,j);
                }
                arr.add(i);


        }
    }

    public static void main(String[] argc){
        Scanner in = new Scanner(System.in);
        long size=0,start=in.nextLong(),counter=0;

        DividersList dividers = new DividersList();//Создание списка делителей start
        Dividers(dividers,start,start,1);

        size = dividers.size();
        long[][] help = new long[2][10000];
        System.out.println("graph {");

        System.out.println(dividers);

        for(int i = (int)size-1; i>=0;i--){
            for (int j = i-1;j>=0;j--){
                if (dividers.get(i)%dividers.get(j)==0){
                    boolean flag = true;
                    for (int q = i-1;q>j;q--){
                        if (dividers.get(i)%dividers.get(q)==0 && dividers.get(q)%dividers.get(j)==0){
                            flag = false;
                            break;
                        }
                    }
                    if (flag){
                        help[0][(int)counter] =dividers.get(i);
                        help[1][(int)counter++]  = dividers.get(j);
                    }
                }
            }
        }

        for(int i = 0 ; i <counter;i++)
            System.out.println("\t" + help[0][i] + "--" + help[1][i]);
        System.out.println("}");

    }
}

