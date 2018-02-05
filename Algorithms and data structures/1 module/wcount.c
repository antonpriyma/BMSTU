#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

int wcount(char *s)
{
    bool flag;
    int i,count=0;

    flag = false;

    for(i=0;i<strlen(s);i++)
    {
        if ((s[i]!=' ') && (flag==false))
        {
            flag = true;
            count++;
        }
        else
        if ((s[i]==' ') && (flag == true))
        {
            flag = false;
        }
    }
    return count;
    }



int main()
{
   char s[1000];
   gets(s);
   printf("%d", wcount(s));
}