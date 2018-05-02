import java.util.stream.IntStream;

public class Factorial {
    public static void main(String[] argc){
        if (argc.length==0){
            System.out.println("usage: java Factorial x");
        }
        else{
            int n = Integer.parseInt(argc[0]);
            int f = IntStream.range(1,n+1).reduce(1,(r,x)->r*x);
            System.out.println(f);
        }
    }
}
