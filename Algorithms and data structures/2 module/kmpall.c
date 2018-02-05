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
	for(i = 0, j = 0; i < N; i++)
	{
		while(j > 0 && p[j] != s[i])
			j = d[j - 1];
		if(p[j] == s[i])
            j++;
		if(j == M)
        {
		    free(d);
            return i - j + 1;
        }
	}
    free(d);
	return -1;
}

int main(int argc, char **argv)
{
int i;
char *s,*s1;
s=(char *)malloc(1000);

s1=(char *)malloc(1000);
strcpy(s1,argv[1]);
strcpy(s,argv[2]);
int j=0,k2=0,buf,count,k1=0,k = strlen(s1);

while ( (count=strlen(s))>=k)
{
    if (j>1)
    k2=k1+1+k2;
    else
        k2=k1;
    k1 = knut(s,s1);
    if (k1==-1)
        break;
    if (j==0)
    printf("%d ", k1+k2);
    else
        printf("%d ", k1+k2+1);
    j++;
    for (i=0;i<count-k1-1;i++)
    {
        s[i]=s[i+k1+1];

    }
    s[count-k1-1]='\0';


}
free(s1);
free(s);
return 0;

}
