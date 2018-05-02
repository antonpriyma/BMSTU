import java.util.ArrayList;
import java.util.Scanner;

public class MinDist {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        String s = in.nextLine();
        int min = s.length();
        char x = in.next().charAt(0), y = in.next().charAt(0);
        ArrayList<Integer> xlist = new ArrayList<>(), ylist = new ArrayList<>();

        for (int i =0;i<s.length();i++){
            if (s.charAt(i)==x){
                xlist.add(i);
            }
            if (s.charAt(i)==y){
                ylist.add(i);
            }
        }


        for (int i =0;i<xlist.size();i++){
            for (int j=0;j<ylist.size();j++){
                if (Math.abs(xlist.get(i)-ylist.get(j))<min){
                    min = Math.abs(xlist.get(i)-ylist.get(j));
                }
            }
        }
        System.out.println(min-1);
    }
}
