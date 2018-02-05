#include <stdio.h>
#include <stdlib.h>
#include <math.h>


int compare(const void *a, const void *b)
{
    return abs(*(int *)a) - abs(*(int *)b);
}

void swap(void *a, void *b, size_t width)
{
    void *temp = malloc(width);
    memcpy(temp, a, width);
    memcpy(a, b, width);
    memcpy(b, temp, width);

    free(temp);
}

void insertion_sort(void *base, size_t nel, size_t width,
          int (*compare)(const void *a, const void *b))
{
    int i;
    long loc;
    for (i = 0; i < nel; i++)
    {
        for (loc = i - 1; loc >= 0
                               && compare(base + loc * width, base + (loc + 1) * width) > 0; loc--)
            swap(base + loc * width, base + (loc + 1) * width, width);
    }
}

char mergesort(int *a,int *buf, int start, int end)
{
    int i1,k,i,j;
    if (start==end) return 0;
    if ((end-start)<5)
    {
        insertion_sort(a,end+1,sizeof(int),compare);
        return 0;
    }
    int mid = (start+end)/2;
    
    mergesort(a,buf,start,mid);
    mergesort(a,buf,mid+1,end);

    int left=start;
    int right = mid+1;
    int bufi=start;

    while ((left<=mid) && (right<=end))
    {
        if (abs(a[right])>=abs(a[left]))
        {
            buf[bufi]=a[left];
            left++;
        }
        else
        {
            buf[bufi]=a[right];
            right++;
        }
        bufi++;
    }

    for (i=left;i<=mid;i++)
    {
        buf[bufi]=a[i];
        bufi++;
    }
    for (i=right;i<=end;i++)
    {
        buf[bufi]=a[i];
        bufi++;
    }

        for (i=start,bufi=0;i<=end;i++,bufi++)
        a[i]=buf[i];
}

int main()
{
    int i,n;
    scanf("%d", &n);
    int *a,*b;
    a = (int *)malloc((n+1)*sizeof(int));
    b = (int *)malloc(n*sizeof(int));
    for (i=0;i<n;i++)
        scanf("%d",&a[i]);
    mergesort(a,b,0,n-1);
    for (i=0;i<n;i++)
        printf("%d ",a[i]);
    free(a);
    free(b);
    return 0;
}
