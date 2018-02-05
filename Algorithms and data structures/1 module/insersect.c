#include <stdio.h>
#include <stdlib.h>
#include <math.h>

int checkbit(const int value, const int position)
{
        return((value & (1 << position)) != 0);
}

int main()
{
    int n,i,m,j,x;
    unsigned int z,y,k,k1;

    y =0;
    z =0;
    j = 0;

    scanf("%d", &n);
    for (i = 0; i<n; i++)
    {
        scanf("%d", &x);
        y = y + pow(2,x) ;
    }

    scanf("%d", &n);
    for (i = 0; i<n; i++)
    {
        scanf("%d", &x);
        z = z + pow(2,x) ;
    }

    k = (y & z);
    k1 =k;

    while (k1/2 >= 1)
    {
        j++;
        k1 = k1/2;
    }

    for (i = 0; i<= j; i++)
    {
        m = (checkbit(k, i) );
        if (m == 1)
                printf("%d ", i);
    }
    return 0;

}
