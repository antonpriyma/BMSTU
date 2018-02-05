#include <stdio.h>
#include <stdlib.h>

//Суммы, образующие степени двойки.

int checkbit(const int value, const int position)
{
    return((value & (1 << position)) != 0);
}

int checkstep(const int x)
{
    int i,j=0,m,count=0,k1 =x;
    if (!(x==1) && ((x%2!=0) || (x<0))) return 0;
    while (k1/2 >= 1)
    {
        j++;
        k1 = k1/2;
    }

    for (i = 0; i<= j; i++)
    {
        m = (checkbit(x, i) );
        if (m == 1)
            count++;
        if (count>1) return 0;
    }
    if (count==1) return 1;
    else
        return 0;
}//проверка на степень двойки

int comb(int *a, int m, int n)
{
    return comb_rec(a, m ,0,n, 0);
}

int comb_rec(int *a, int m ,int c, int n,int sum)
{
    int i,buf,j,q1[n],n1=n;
    if (m==0)
        sum += checkstep(c);
    else
    {
        for (i=0;i<n;i++)
            q1[i]=a[i];
        for (i=0;i<n1;i++)
        {
            if (n<m)
                break;
            buf=q1[0];

            for(j=0;j<n-1;j++)
                q1[j]=q1[j+1];
            q1[n-1]=0;
            n--;
            sum += comb_rec(q1,m-1,c+buf,n,0);
        }
    }
    return sum;
}

int main()
{
    int s1=0;
    int n,i;
    scanf("%d", &n);
    int a[n];

    for (i=0;i<n;i++)
        scanf("%d",&a[i]);

    for (i=1;i<=n;i++)
        s1 += comb(a,i,n);

    printf("%d",s1);
    return 0;
}






