#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

void Kadane(float* a, int n)
{
        int l,r;
    l=n-1;
    r=n-1;

    float maxsum = a[0];
    int start = n-1, i=n-1;
    float sum=0,sum1=1;
    while (i>=0)
    {
        sum1=sum;
        sum += a[i];
        if (sum>=maxsum)
        {
            maxsum=sum;
            l=start;
            r=i;
        }
        i--;
        if (sum<0)
        {
            sum = 0;
            sum1= 1;
            start = i;
        }
    }

    printf("%d %d", r ,l);


}

int main()
{
    int n;
    scanf("%d", &n);

    int j=0,i;
    float a=0,b=0,cur,*t;
    t = (float *)malloc(n*sizeof(float));
    char c=0,s[1000],s1[1000];



    for (i=0;i<n;i++)
    {
       scanf("%f/%f", &a, &b);

        t[i]=log(a/b);



    }

    Kadane(t,n);
    free(t);
    return 0;



}
