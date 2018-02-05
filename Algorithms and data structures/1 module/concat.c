#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <malloc.h>

char *concat(char **s, int n){

    int k,i,len = 0;

    for (i=0;i<n;i++)
    {
        len += strlen(*(s+i));
    }

    char *result = (char *)malloc(len+1);
    result[0]='\0';
    k=0;
    for (i=0;i<n;i++)
    {
        strcat(result,*(s+i));
        k += strlen(*(s+i));

    }
    printf("%s",result);

    return result;
}

int main()
{

    int n,j,n1,i;
    j=0;
    n1=0;
    n=1;
    i=0;
    scanf("%d", &n);
    getchar();
    char *s[n],*buff = NULL;//ћассив указателей, буферна€ строка
    char c;

    for (j=0;j<n;j++){
      n1=100;
      i=0;
      buff=(char *)malloc(100);
      while ((c=getchar())!='\n'){
        if (i++>=n1){
            buff = (char *)realloc(buff,n1+1);
            n1 ++;
        }
        *(buff+i-1)=c;//считывание строки

      }
      *(buff+i) ='\0';
      s[j]=buff;//присваиванию элементу массива строк
    }

        free(concat(s,n));//вывод результута, освобождение пам€ти

        for (i=0;i<n;i++)
            free(s[i]);//освобождение пам€ти

    return 0;
}








