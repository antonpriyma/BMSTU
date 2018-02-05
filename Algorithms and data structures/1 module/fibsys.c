#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include <malloc.h>
#include <string.h>
#include <stdint.h>

int main()
{
    int i,j,c;
    char *s;
    unsigned long n,x,k,f,f1;
    


    n=0;
    c=0;
    f = 1;
    f1 = 1;
    x=0;
    scanf("%lu", &x);


    if (x!=0){
        while (x>0){
            c++;
            
            while (x>=f){
                k=f;
                f += f1;
                f1=k;
                n++;
            }
            
            if (c==1){
                s = (char*)malloc((n+1)*sizeof(char));
                for (j=0;j<n;j++)
                    *(s+j)='0';
            }
            
            *(s+j) = 0;
            *(s+j-n)='1';
            
            x = x - f1;
            f1=1;
            f=1;
            n=0;

        }   
    i=0;
    while (s[i]!=0){
        putchar(s[i]);
        i++;
    }
    free(s);
    }
    else
        printf("0");
    return 0;
}
