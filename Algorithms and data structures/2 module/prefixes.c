#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int *prefix(char* s, int n)
{
     // Динамический массив длины М
        int *d = (int*)malloc(n * sizeof(int));
    int i,j;
	// Вычисление префикс-функции
	d[0] = 0;
	for(i = 1, j = 0; i < n; i++)
	{
		while(j > 0 && s[j] != s[i])
			j = d[j-1];
		if(s[j] == s[i])
			j++;
		d[i] = j;
	}

	return d;
}

int* result(int n, int* d)
{
    int *ans;
    ans = (int *)malloc((n+1)*sizeof(int));
    int i,j;
    for (i=0;i<n;++i)
        ans[i]=0;
    for (i=0;i<n;++i)
    {
        printf("%d ", d[i]);
        printf("\n");
        ++ans[d[i]];
        for (j=0;j<n;j++)
            printf("%d ", ans[j]);
            printf("\n");
    }


    for (i = n-1; i>0; --i)
        ans[d[i-1]] += ans[i];
    return ans;
}
int main(int argc, char** argv)
{
    int i,len = strlen(argv[1]);
    int *a = prefix(argv[1],len);
    //for (i=0;i<len;i++)
    //printf("%d ", a[i]);
    //printf("\n");

    int size=0,answer=0;
    for (i = 1; i < len; i++) {
      if ((a[i] > 0) && ((i + 1) % (i + 1 - a[i]) == 0)) {
        size = i + 1;
        answer = (i + 1)/(i + 1 - a[i]);
        printf("%d %d\n", size, answer);
      }
    }
    free(a);
    


 //   b  = result(len, a);
  //  for (i=0;i<len;i++)
   //     printf("%d ",b[i]);


    return 0;
}
