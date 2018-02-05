#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int knut (char* s, char* p)
{
        int i, j, N, M;
        N = strlen(s);
        M = strlen(p);

    // Динамический массив длины М
        int *d = (int*)malloc(M * sizeof(int));

	// Вычисление префикс-функции
	d[0] = 0;
	for(i = 1, j = 0; i < M; i++)
	{
		while(j > 0 && p[j] != p[i])
			j = d[j-1];
		if(p[j] == p[i])
			j++;
		d[i] = j;
	}

	// Поиск
	int cond =1;
	for(i = 0, j = 0; i < N; i++)
	{
		while(j > 0 && p[j] != s[i])
			j = d[j - 1];
		if(p[j] == s[i])
            j++;
		if (j==0)
            cond = 0;
        if(j == strlen(s)) j = p[j - 1];

	}
    free(d);
	return cond;
}


int main(int argc, char **argv)
{
   if (knut(argv[2], argv[1])==1) 
        printf("yes\n")
   else
        printf("no\n");
   return 0;
}
