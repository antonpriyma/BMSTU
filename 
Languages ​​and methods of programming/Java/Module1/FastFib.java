import java.math.BigInteger;
import java.util.Scanner;

public class FastFib {
    public static BigInteger[][] prod(BigInteger[][] a, BigInteger[][] b) {
        Integer[][] result = new BigInteger[2][2];
        result[0][0]=a[0][0].multiply(b[0][0]).add(a[0][1].multiply(b[1][0]));
        result[0][1]=a[0][0].multiply(b[0][1]).add(a[0][1].multiply(b[1][1]));;
        result[1][0]=a[1][0].multiply(b[0][0]).add(a[1][1].multiply(b[1][0]));
        result[1][1]=a[1][0].multiply(b[0][1]).add(a[1][1].multiply(b[1][1]));

        return result;
    }

    public static BigInteger[][] pow(BigInteger[][] a, int n) {

        BigInteger[][] result = new BigInteger[a.length][a.length];
        BigInteger[][] result1 = new BigInteger[a.length][a.length];
        BigInteger[][] result2 = new BigInteger[a.length][a.length];
        BigInteger[][] result3 = new BigInteger[a.length][a.length];
        BigInteger[][] help = new BigInteger[2][2];
        help[0][0]=BigInteger.ONE;
        help[0][1]=BigInteger.ZERO;
        help[1][0]=BigInteger.ZERO;
        help[1][1]=BigInteger.ONE;

        result = a;
        int i = 1,j=0;
        if (n==0) return help;
        else if (n%2==1) return prod(a,pow(a,n-1));
        else{
            BigInteger[][] b = pow(a,n/2);
            return prod(b,b);
        }

    }


    public static void main(String[] argc){
        long startTime = System.currentTimeMillis();
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        BigInteger[][] matr = new BigInteger[2][2];
        matr[0][0]=BigInteger.ONE;
        matr[0][1]=BigInteger.ONE;
        matr[1][0]=BigInteger.ONE;
        matr[1][1]=BigInteger.ZERO;

        BigInteger[][] result = pow(matr,n-1);

        BigInteger[][] result1 = new BigInteger[2][1];
        result1[1][0]=result[1][0].add(result[1][1]);
        System.out.println(result[0][0]);
        //long timeSpent = System.currentTimeMillis() - startTime;
        //System.out.println("программа выполнялась " + timeSpent + " миллисекунд");

        }

    }


