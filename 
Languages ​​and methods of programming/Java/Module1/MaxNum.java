import java.io.IOException;
import java.util.Comparator;
import java.util.Scanner;

public class MaxNum {

    static class Comp implements Comparator<String>{
        @Override
        public int compare(String num1, String num2){
            if (num1==num2) return 0;
            else{
                int i=0;
                char[] arr1 = num1.toCharArray(),arr2 = num2.toCharArray();
                while(arr1[i]==arr2[i]){
                    i++;
                }
                if ((int)arr1[i]>(int)arr2[i]) return -1;
                return 1;
            }
        }
    }



    private static void Scan(String s,String[] arr, int n) throws IOException {
        char[] c = s.toCharArray();
        String s1 = "";
        Scanner in = new Scanner(System.in);
        int j=0;
        for(int i=0;i<s.length();i++){
            if (c[i]!=' ')
                s1+=c[i];
            else {
                if(s1!="") {
                    arr[j]=s1;
                    j++;
                }
                s1 = "";
            }
        }
        arr[arr.length-1]=s1;
    }

    public static void sort(int[] arr,int n){
        for (int i=0;i<n-1;i++) {
            String x = arr[i] +""+  arr[i+1];
            String y = arr[i+1] +""+  arr[i];
            if (y.compareTo(x)>0) {
                int buf=arr[i];
                arr[i]=arr[i+1];
                arr[i+1]=buf;
            }
            for (int j=i;j>0;j--) {
                String z = arr[j] + ""+arr[j-1];
                String d = arr[j-1] +""+ arr[j];
                if (z.compareTo(d)>0) {
                    int buf=arr[j];
                    arr[j]=arr[j-1];
                    arr[j-1]=buf;
                }
            }
        }
    }


    public static void main(String[] argc) throws IOException {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        in.nextLine();
        int[] arr = new int[n];

        //for(int i = 0; i<n; i++)
       //     arr.add(in.nextLine());
        String s = new String();
        for (int i =0;i<n;i++){
            arr[i]=in.nextInt();
        }
        //Scan(s,arr,n);
        //arr.sort(new Comp());
        sort(arr,n);

        s="";
        for(int i=0;i<n;i++)
            s+=arr[i];
        System.out.println(s);
    }
}
