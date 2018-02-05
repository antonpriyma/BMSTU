#include <stdio.h>
#include <stdlib.h>
void selectionSort(int *num, int start, int end)
{
  int i,j,min, temp; 
  for (i = start; i < end ; i++)
  {
    min = i; 
    
    for (j = i + 1; j <= end; j++)  
    {
      if (num[j] < num[min]) 
        min = j;       
    }
    temp = num[i];     
    num[i] = num[min];
    num[min] = temp;
  }
}

void swap(int *a, int c, int b)
{
    int k;
    k = a[c];
    a[c] = a[b];
    a[b] = k;
}

void quick(int *a, int start, int end, int m)
{
        
    int l,e,i,x,mid;
    l=start;
    e=end;
    mid = l/2 + e/2;
    x = a[mid];
    if ((end-start)<5)
        selectionSort(a,start,end);
    else
    {

    while (l<=e)
    {
        while (a[l]<x)
            l++;
        while (a[e]>x)
            e--;
        if (l<=e)
            swap(a, l++,e--);
    }

    if (l<end)
            quick(a, l, end, m);
    if (e>start)
            quick(a, start, e, m);
    }
}

void sort(int *a, int n, int m)
{
    quick(a, 0 ,n-1, m);
}


int main()
{
    int i,n,m;
    scanf("%d\n%d", &n , &m);
    int a[n];

    for (i=0;i<n;i++)
        scanf("%d", &a[i]);
        sort(a, n,m);

     for (i=0;i<n;i++)
        printf("%d ", a[i]);

}
