#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *fibstr(int n)
{
        int i,count=1,count1=1,buf;


        for (i=3;i<=n;i++){
                buf=count;
                count += count1;
                count1=buf;
        }

        char sbuf[count],a[count],b[count],*fib = (char *)malloc(count+1);
        fib[0]='\0';
        sbuf[0]='\0';
        a[0]='\0';
        b[0]='\0';
        strcat(a,"a");
        strcat(b,"b");
        if (n==1)
                strcat(fib,"a");
        if (n==2)
                strcat(fib,"b");
        if (n>2)
        {
                for (i=3;i<=n;i++){
                strcpy(sbuf,b);
                strcat(b,a);
                strcpy(a,sbuf);
        }
        strcpy(fib,b);
        int len=strlen(fib);

        for(i=0;i<len/2;i++)
        {
                char t=fib[i];
                fib[i]=fib[len-i-1];
                fib[len-i-1]=t;
        }
        }
        puts(fib);
        return fib;
        }

int main()
{
    int n;
    scanf("%d", &n);
    free(fibstr(n));
    return 0;
}
