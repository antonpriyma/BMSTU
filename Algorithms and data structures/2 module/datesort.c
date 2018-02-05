#include <stdio.h>
#include <stdlib.h>

typedef struct
{
    int day;
    int month;
    int year;
} date;

int* directionsort(date *a, unsigned long n, int count1)
    {
        int k,k1;
        switch(count1)
        {
            case 1: k=32; break;
            case 2: k = 13; break;
            case 3: k=62; break;
        }
        k1=k;
        long i,j,count[k];
        for (i=0;i<k;i++)
            count[i]=0;
        date b[n];

        for (i=0;i<n;i++)
        {
            if (count1 == 3)
            k = a[i].year-1970;
            else if (count1==2)
            {
                k= a[i].month;

            }

            else
                k = a[i].day;
                count[(int)k]++;
        }

        for (i=1;i<=k1;i++)
            count[i]=count[i]+count[i-1];
        i = n - 1;
        while (i>=0)
        {
            if (count1 == 3)
            k = a[i].year-1970;
            else if (count1==2)
                k= a[i].month;
            else
                k = a[i].day;

            j = count[k] - 1;
            count[k] = j;

             b[j]=a[i];

            i--;
        }

        if (count1<3)
        {

            directionsort(b, n, count1+1);
        }
        else
        {
            for (i=0;i<n;i++)
    {
        printf("%d %d %d\n", b[i].year, b[i].month, b[i].day);
    }
            return b;
        }


    }

int main()
{

    int n;
    scanf("%d", &n);
    date a[n];
    int i;
    for (i=0;i<n;i++)
    {
        scanf("%d %d %d", &a[i].year, &a[i].month, &a[i].day);
    }
    directionsort(a ,n , 1);

}
