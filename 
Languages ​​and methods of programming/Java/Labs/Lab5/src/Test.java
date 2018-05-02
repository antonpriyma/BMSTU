import java.util.ArrayList;

public class Test {
    public static void main(String[] argc){
        ArrayList<ArrayList<Integer>> testarray = new ArrayList<ArrayList<Integer>>();
        ArrayList<Integer> test1 = new ArrayList<Integer>();
        test1.add(10);
        test1.add(20);
        ArrayList<Integer> test2 = new ArrayList<Integer>();
        test2.add(30);
        test2.add(40);
        testarray.add(test1);
        testarray.add(test2);
        System.out.println(Stramtest.max(testarray));


    }
}
