#include <stdio.h>
#include <stdlib.h>
#include <math.h>
//Наибольший простой делитель
int main()
{
    int x2,x1,i,s,j,x;
    scanf("%d", &x);
    char     a[(int)sqrt((abs(x)))+1];
    x1=(int)sqrt((abs(x)))+1;

    for(i=0; i<x1; i++){
           a[i] = 1;
       }
       //поскольку 1 не простое число, обнулим ячейку с этим числом
       a[1]=0;
    for(s=2; s<x1; s++){
           if(a[s]!=0){
               for(j=s*2; j<x1; j+=s){
                   a[j]=0;
               }
           }
       }
    x2=0;
     while (abs(x)>1){
     if (x2==x){
        i=abs(x);
        break;
     }
        x2=x;
        for (i=1; i<x1; i++)
            if (a[i]!=0)
                if (x%i==0)
                {
                    x /= i;
                    break;
                }
     }
                printf("%d", i);
    return 0;
}

