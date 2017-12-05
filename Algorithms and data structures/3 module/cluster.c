#include <stdio.h>
#include <stdlib.h>

typedef struct heap heap1;
typedef struct heapquene quenep;

int compare(int a, int b)
{
    if (a>b) return 1;
    return 0;
}


struct heap
{
    int v1;
    int v2;
    int sum;
};

struct heapquene
{
  int cap;
  int count;
  heap1 *Heap;
};

void swap(quenep *q, int i ,int  j)
{
    heap1 *buf;
    buf = malloc(sizeof(heap1));
    buf->sum=q->Heap[i].sum;
    buf->v1=q->Heap[i].v1;
    buf->v2=q->Heap[i].v2;
    q->Heap[i].sum=q->Heap[j].sum;
    q->Heap[i].v1=q->Heap[j].v1;
    q->Heap[i].v2=q->Heap[j].v2;
    q->Heap[j].v1=buf->v1;
    q->Heap[j].sum=buf->sum;
    q->Heap[j].v2=buf->v2;
    free(buf);
}

void initpriorityq(quenep *q, int n)
{
    q->Heap = malloc(n*sizeof(heap1));
    q->cap=n;
    q->count=0;
}



void insert(quenep* y,heap1 b)
{
    int i;
    i=y->count;
    y->count+=1;
    y->Heap[i]=b;
    while ((i > 0) && compare(y->Heap[(i - 1)/2 ].sum ,y->Heap[i].sum))
    {
        swap(y,(i - 1) / 2, i);
        i = (i - 1) / 2;
    }
}

void heapify (int i, int n, quenep* y)
{
    int l,r,j;
    while (1)
    {
        l=2*i+1;
        r=l+1;
        j=i;
        if ((l < n) && (compare(y->Heap[i].sum,y->Heap[l].sum)))
            i = l;
        if ((r < n) && (compare(y->Heap[i].sum,y->Heap[r].sum)))
            i=r;
        if (i==j)
            break;
        swap(y,i,j);
    }
}

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
    heap1 *cur;
    quenep *q;
    q = (quenep *)malloc(sizeof(quenep));
    heap1 buf;
    int i,n,k;
    scanf("%d", &k);
    scanf("%d", &n);
    initpriorityq(q,n);
    cur = (heap1 *)malloc(n*sizeof(heap1));

    for (i=0;i<n;i++)
    {
        scanf("%d %d", &cur[i].v1, &cur[i].v2);
    }

    for (i=0;i<k;++i)
    {
        cur[i].sum = cur[i].v1 + cur[i].v2;
        insert(q, cur[i]);
    }
    //инициализация

    i=k;
    int max;
    while (q->count)
    {
        buf=extract(q);
        if (i<n)
        {
            if (cur[i].v1>buf.sum)
            max=cur[i].v1;
        else
            max=buf.sum;
            cur[i].sum=max+cur[i].v2;
            insert(q, cur[i]);
            i++;
        }
    }
        printf("%d", buf.sum);
        free(cur);
        free(q->Heap);
        free(q);
        return 0;
}
