import java.util.Arrays;
import java.util.Scanner;

public class Gauss {

    public static void quickSort(Ratioanal[] a,int n) {
        int startIndex = 0;
        int endIndex = n-1;
        doSort(startIndex, endIndex,a);
    }

    private static void doSort(int start, int end, Ratioanal[] a) {
        if (start >= end)
            return;
        int i = start, j = end;
        int cur = i - (i - j) / 2;
        while (i < j) {
            while (i < cur && (Math.abs(a[i].numerator) >= Math.abs(a[cur].numerator))) {
                i++;
            }
            while (j > cur && (Math.abs(a[cur].numerator) >= Math.abs(a[j].numerator))) {
                j--;
            }
            if (i < j) {
                int temp = a[i].numerator;
                a[i].numerator = a[j].numerator;
                a[j].numerator = temp;
                if (i == cur) {
                    cur = j;
                }
                else if (j == cur) {
                    cur = i;
                }
            }
        }
        doSort(start, cur,a);
        doSort(cur+1, end,a);
    }//Сортировка

    static class Ratioanal{
        private int numerator;
        private int denomenator;

        public int nod(int a, int b) {
            int c;
            a = Math.abs(a);
            b = Math.abs(b);
            while (b != 0) {
                c = a % b;
                a = b;
                b = c;
            }
            if (a == 0) {
                 System.out.println("No solution");
                 System.exit(0);
            }
            return a;
        }

        public void cut(Ratioanal r){
            int a = r.nod(r.numerator, r.denomenator);
            r.numerator /= a;
            r.denomenator /= a;
            if (r.denomenator<0 && r.numerator<0){
                r.denomenator*=-1;
                r.numerator*=-1;
            }

            if (r.denomenator<0 && r.numerator>=0){
                r.denomenator*=-1;
                r.numerator*=-1;
            }

        }

        public Ratioanal(int x){
            this.numerator=x;
            this.denomenator=1;
        }

        public Ratioanal(boolean t){
            this.numerator=1;
            this.denomenator=0;
        }


        public Ratioanal(){
            this.numerator=0;
            this.denomenator=1;
        }

        private int gcd(int a, int b){
            if (b==0) return a;
            return gcd(a,a%b);
        }

        public Ratioanal add(Ratioanal x){
            int gcd = gcd(this.denomenator,x.denomenator);
            Ratioanal result = new Ratioanal();
            result.numerator=this.numerator*(x.denomenator)+x.numerator*(this.denomenator);
            result.denomenator=x.denomenator*this.denomenator;
            cut(result);
            return result;
        }

        public Ratioanal sub(Ratioanal x){
            int gcd = gcd(this.denomenator,x.denomenator);
            Ratioanal result = new Ratioanal();
            result.numerator=this.numerator*(x.denomenator)-x.numerator*(this.denomenator);
            result.denomenator=x.denomenator*this.denomenator;
            cut(result);
            return result;
        }

        public Ratioanal mult(Ratioanal x){
            Ratioanal result = new Ratioanal();
            result.numerator=this.numerator*x.numerator;
            result.denomenator=this.denomenator*x.denomenator;
            cut(result);
            return result;
        }

        public Ratioanal div(Ratioanal x){
            Ratioanal result = new Ratioanal();
            result.numerator=this.numerator*x.denomenator;
            result.denomenator=this.denomenator*x.numerator;
            cut(result);
            return result;
        }

        @Override
        public String toString(){
            String s = new String();
            s=this.numerator+"/"+this.denomenator;
            return s;
        }


    }

    public static void pivoting(Ratioanal matr[][],int i,int n,int[] where,int p){
        double max = 0;
        int maxi=p;

        for (int j=p;j<n;++j){
            Ratioanal o = matr[j][i];
            double a= Math.abs((float)o.numerator/o.denomenator);
            double b = Math.abs((float)matr[maxi][i].numerator/matr[maxi][i].denomenator);
            if (a>b){
                maxi=j;
                //max=Math.abs(matr[j][i].numerator/matr[j][i].denomenator);
            }
        }

        for (int q = i; q < n+1 ; ++q) {
            Ratioanal temp = matr[p][q];
            matr[p][q] = matr[maxi][q];
            matr[maxi][q] = temp;
        }
        where[i]=p;
    }




    public static void toStairsView(Ratioanal[][] matr,int n,int[] where){

         Ratioanal coefficient = new Ratioanal();
         for(int k=0, p=0;k<=n && p<n;++k) {
             if (p==3 && k==3){
                 p = p;
             }
             pivoting(matr,k,n,where,p);
                 for (int i = 0; i < n; i++) {
                     if (i!=p) {
                         if (matr[p][k].numerator==0)
                             continue;
                         coefficient = matr[i][k].div(matr[p][k]);
                         for (int j = k; j < n + 1; j++) {
                             matr[i][j] = matr[i][j].sub(matr[p][j].mult(coefficient));
                         }
                     }

                 }
                 p++;
             }
             }



    public static void Solve(Ratioanal[][] matr, int n,Ratioanal[] solvearray,int[] where) {
        Ratioanal x= new Ratioanal();

        for (int i = n-1;i>=0;i--){
            x=new Ratioanal();
            for(int j=0;j<n;j++){
                if (solvearray[j].denomenator!=0) {
                    x = x.add(solvearray[j].mult(matr[i][j]));
                }
            }
            for(int j=0;j<n;j++) {
                if (solvearray[j].denomenator == 0 && matr[i][j].numerator!=0) {
                    solvearray[j] = matr[i][n].sub(x).div(matr[i][j]);
                    break;
                }
            }
        }

        for (int i =0;i<n;i++){
            if (solvearray[i].denomenator == 0) {
                System.out.println("No solution");
                System.exit(0);
            }
        }
        /*
        for(int i =0;i<=n;++i){
            if (where[i]!=-1)
                solvearray[i] = matr[where[i]][n].div(matr[where[i]][i]);

        }

        for (int i=0; i<n; ++i) {
            Ratioanal sum = new Ratioanal();
            for (int j=0; j<n; ++j)
                sum = sum.add(solvearray[j].mult(matr[i][j]));
        }
        */

        for(int i=0;i<n;i++)
            System.out.println(solvearray[i]);
    }


    public static void main(String[] argc){
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();

        Ratioanal matr[][]= new Ratioanal[n][n+1];
        Ratioanal solvearray[] = new Ratioanal[n];

        for (int i=0;i<n;i++){
            solvearray[i]=new Ratioanal(true);
        }
        for(int i = 0;i<n;i++){
            for(int j=0;j<n+1;j++)
                matr[i][j]=new Ratioanal(in.nextInt());
        }


        int where[] = new int[n+1];
        for(int i=0;i<n;i++){
            where[0]=-1;
        }
        toStairsView(matr,n,where);
        Solve(matr,n,solvearray,where);





    }
}

