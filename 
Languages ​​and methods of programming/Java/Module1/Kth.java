import java.util.Scanner;

public class Kth {
    private static long x;
    private static long result;
    private static long power,i,help;

    private static long pow(int x, long y){
        long result = x;
        if (y==0) return 1;
        for(int i =1;i<y;i++){
            result*=x;
        }
        return result;
    }

    private static long whichisx(long help, long i){
        long buf,result;
        if (help/i==0) buf = 1;
        else{
            buf = help/i + help%i;
        }
        if (buf==1) {
            result = pow(10, i-1);
        }
        else{
            result=pow(10,i-1)+buf-1;
        }
       return result;
    }

    public static void main(String[] argc){
        Scanner in = new Scanner(System.in);
        x = in.nextLong() + 1;
        power = 1;
        help = x-9;

        while(help>0){
            x=help;
            help -=9*pow(10,power)*(power+1);
            power++;
        }//Диапазон значений x

        result = whichisx(x,power);//Число, которому принадлежит х

        for(x= power - x%power;x>0 && x%power!=0;x--){
            result/=10;
        }
        System.out.println(result%10);
    }
}

