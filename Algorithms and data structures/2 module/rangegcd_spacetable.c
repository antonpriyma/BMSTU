#include <stdio.h>
#include <stdlib.h>
#include <math.h>


int nod(int a, int b)
{
    int f=0,c = abs(a);
    int d = abs(b);
    while (d)
    {
        f = c%d;
        c = d;
        d = f;
    }
    return c;
}

void computelog(int m, int *log)
{
        int i, j;
	for (j = 0, i = 1; j < m;)
		if (j < (1 << i)) {
			log[j] = i - 1;
			j++;
		}	
		else
			i++;
		
}

void SparceTable_Build(int *a, int** st, int n, int m)
{
    int i=0, j =1;
    while (i<n)
    {
        st[i][0]=a[i];
        i++;
    }

    while (j<m)
    {
        i=0;
        while (i<=(n-(1<<j)))
        {
           st[i][j]= nod(st[i][j-1], st[i+(1<<(j-1))][j-1]);
           i++;
        }
        j++;
    }
}

void sparsetable_query(int n, int m, int **st, int l, int r, int *log)
{
	int j = log[r-l+1];
	printf("%d\n", nod(st[l][j], st[r-(1 << j) + 1][j]));
}

int main()
{
    int q,i,n,m,l,r;
    scanf("%d", &n);
    int a[n];
    for (i=0;i<n;i++)
        scanf("%d", &a[i]);
    int *lg;
    lg = (int *)malloc((n+1)*sizeof(int));
    computelog(n+1,lg);
    m = lg[n]+1;
    int **st;
    st = (int**)malloc(n*sizeof(int*));
    for(i = 0; i < n; i++)
		st[i] = (int*)malloc(m*sizeof(int));
    SparceTable_Build(a,st,n,m);
    scanf("%d", &q);
    for(i = 0; i < q; i++)
    {
		scanf("%d %d", &l, &r);
		sparsetable_query(n, m, st,l,r,lg);
	}

	for(i = 0; i < n; i++)
		free(st[i]);
	free(st);
	free(lg);
	return 0;
}








