#include <stdio.h>
#include <stdlib.h>

typedef struct heap heap1;
typedef struct heapquene quenep;

struct heap
{
    int index;
    int v;
};

struct heapquene
{
  int cap;
  int count;
  heap1 *Heap;
};

int compare(int a, int b)
{
    if (a>b) return 1;
    return 0;
}

void swap(quenep *q, int i ,int  j)
{
    heap1 *buf;
    buf = malloc(sizeof(heap1));
    buf->index=q->Heap[i].index;
    buf->v=q->Heap[i].v;
    q->Heap[i].index=q->Heap[j].index;
    q->Heap[i].v=q->Heap[j].v;
    q->Heap[j].v=buf->v;
    q->Heap[j].index=buf->index;
    free(buf);
}


void initpriorityq(quenep *q, int n)
{
    q->Heap = malloc(n*sizeof(heap1));
    q->cap=n-1;
    q->count=0;
}


void insert(quenep* y,heap1 b)
{
        int i;
        i=y->count;
        y->count+=1;
        y->Heap[i]=b;
        while ((i > 0) && compare(y->Heap[(i - 1)/2 ].v ,y->Heap[i].v))
        {
                swap(y,(i - 1) / 2, i);
                i = (i - 1) / 2;
        }
}


   void heapify (int i, int n, quenep* y) {
        int l,r,j;
        while (1) {
                l=2*i+1;
                r=l+1;
                j=i;
                if ((l < n) && (compare(y->Heap[i].v,y->Heap[l].v)))
                        i = l;
                if ((r < n) && (compare(y->Heap[i].v,y->Heap[r].v)))
                        i=r;
                if (i==j)
                        break;
                swap(y,i,j);
        }
}


//}//упорядочивание дереева

heap1 extract(quenep *q)
{
    if (q->count==0)
    {
        printf("Error ");
        return ;
    }
    heap1 buf = q->Heap[0];
    q->count--;
    if (q->count>0)
    {
        q->Heap[0]=q->Heap[q->count];
        heapify(0,q->count,q);
    }
    return buf;
}


int main()
{
    int result_size=0;
    int i,j,n,size;
    int **arr;

    scanf("%d", &n);
    int siz[n];
    int a[n];
    arr = (int **)malloc((n)*sizeof(int *));
    quenep *q ;
    q = malloc(sizeof(quenep));
    initpriorityq(q,n+1);
    for (i=0;i<n;i++)
    {
        scanf("%d", &siz[i]);
        result_size += siz[i];
        arr[i]=(int *)malloc((siz[i]+1)*sizeof(int));
        a[i]=0;
    }


    for (i=0;i<n;i++)
        for (j=0;j<siz[i];j++)
        scanf("%d", &arr[i][j]);


    for (i=0;i<n;i++)
    {
        if (siz[i]!=0)
        {
            heap1 e;
            e.v=arr[i][0];
            e.index=i;
            insert(q,e);
            a[i]++;
        }
    }

    for (i=0;i<result_size;i++)
    {
        heap1 r;
        r = extract(q);
        printf("%d ",r.v);
         if (a[r.index]!=siz[r.index]) {
                        heap1 x;
                        x.v=arr[r.index][a[r.index]];
                        x.index=r.index;
                        insert(q,x);
                        a[r.index]++;
                }
        }

        free(q->Heap);
        for (i=0;i<n;i++)
                free(arr[i]);
        free(arr);
        free(q);

        return 0;
}











