#include <stdio.h>
#include <stdlib.h>
#include <string.h>


int checkbit(const int value, const int position)
{
    return((value & (1 << position)) != 0);
}

int strdiff(char *a, char *b)
{
    int i,res,j,min=0,k=0;
    if (strlen(a)!=strlen(b))
    {

        if (strlen(a)<=strlen(b))
            min = strlen(a);
        else
            min = strlen(b);
    }
    else
        min = strlen(b);

    for(j=0;j<=min;j++)
    {
        if ((a[j]==(res=a[j]&b[j])) && (b[j]==(res=a[j]&b[j])) )
        {
            k += 8;
            continue;
        }
        else
        {
            for (i=0;i<8;i++)
                if ((checkbit(a[j],i)!=(checkbit(res,i))) || (checkbit(b[j],i)!=(checkbit(res,i))))
                {
                    k += i;
                    break;
                }
            break;
        }   
    }
    if (j>min) 
        return -1;
    return k;
}

