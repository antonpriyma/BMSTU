#include <stdio.h>
#include <stdlib.h>

typedef struct Elem elem;

struct Elem
{
    elem *next;
    elem *prev;
    int v;
};


int checkempty(elem *l)
{
    if (l->next = l)
        return 1;
    return 0;
}

int length(elem *l)
{
    int len =0;
    elem *x = l;
    while (x->next!=l)
    {
        len++;
        x = x->next;
    }
    return len;
}

void insertafter(elem *x, elem *y)
{
    elem *z = x->next;
    x->next=y;
    y->prev=x;
    y->next=z;
    z->prev=y;
}

void delete(elem *x)
{
    elem *y;
    y = x->prev;
    elem *z;
    z = x->next;
    y->next = z;
    z->prev = y;
    x->prev = NULL;
    x->next = NULL;
}

void insertsort(elem *a) {
	elem *b, *c;
	for (b = a->next->next; b != a; b = b->next)
    {
		for (c = b->prev; c != a && b->v < c->v; c = c->prev);
		delete(b);
		insertafter(c, b);
	}
}


elem *initlist()
{
	elem *p = (elem *) malloc(sizeof(elem));
	p->next = p;
	p->prev = p;
	return p;
}


int main()
{
        long int i, n;
        elem *list=initlist(), *y,*p;
        scanf("%ld", &n);
        for (i = 0; i < n; i++)
        {
                y=initlist();//buf
                scanf("%d", &(y->v));
                insertafter(list, y);
        }
        insertsort(list);
        y=list->next;
        while (y != list)
        {
                printf("%d ", y->v);
                p = y->next;
                free(y);
        	y=p;
        }
        free(list);
        return 0;
}
