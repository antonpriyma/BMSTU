#include <stdio.h>
#include <stdlib.h>
#include <math.h>

//Седловая точка в матрице


int main()
{
    int flag,s,t,i,j,k,max,min,imax,jmax;
    scanf("%d %d", &s , &t);
    int a[s][t], b[s],c[t];
    flag=0;

    for (i=0;i<s;i++)
        b[i]=-pow(2,31);

    for (i=0; i<t; i++)
        c[i]=pow(2,31)-1;

    for (i=0;i<s;i++){
        for (j=0;j<t;j++){
            scanf("%d", &a[i][j]);
            if (a[i][j]>b[i])   b[i]=a[i][j];
            if (a[i][j]<c[j])   c[j]=a[i][j];        }

    }

    for (i=0;i<s;i++){
        if (flag==1) break;
        for (j=0;j<t;j++)
            if (b[i]==c[j]){
                printf("%d %d", i, j);
                flag=1;
                break;
            }
    }

    if(flag==0) printf("none");

            return 0;


        }

