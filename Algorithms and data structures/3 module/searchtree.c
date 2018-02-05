//correct new version

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
//бинарное дерево поиска, ранги вершин
typedef struct searchtree searchtree;

struct searchtree{
    searchtree* parent;
    searchtree* right;
    searchtree* left;
    int k;
    int count;
    char v[100];
};

searchtree* inittree(int k, char* v){
    searchtree *t = malloc(sizeof(searchtree));
    t->parent =0;
    t->right =0;
    t->left=0;
    t->count=0;
    t->k=k;
    strcpy(t->v,v);
    return t;

}

int mapempty(searchtree* t){
    if (t==NULL) return 1;
    return 0;
}

searchtree* minimum(searchtree *t){
    searchtree* x;
    if (mapempty(t)) return NULL;
    else{
        x=t;
        while (!mapempty(x->left)) {
            x->count--;
            x = x->left;
        }
    }
    return  x;
}

searchtree* succ(searchtree* t){
    searchtree* y;
    if (!mapempty(t->right))
        return minimum(t->right);
    else{
            y = t->parent;
        while (y!=NULL && (t==y->right)){
            t = y;
            y=y->parent;
        }

    }
    return  y;
}

searchtree* descend(searchtree* t, int k){
    searchtree* x = t;
    while (x!=NULL && x->k != k){
        if (k<x->k)
            x=x->left;
        else
            x=x->right;
    }
    return  x;
}

char* mapsearch(searchtree *t, int k){
    searchtree* x = descend(t, k);
    if (x==NULL) return NULL;
    return x->v;
}

searchtree* searchrank(searchtree *t, int x) {
    searchtree *y = minimum(t);
    int count;
    count = 0;
    while (count++ != x) {
        y = succ(y);
    }
    return y;
}

searchtree* insert(searchtree *t, int k, char *v){
    searchtree* y;
    searchtree *x,*d;
    y = inittree(k,v);



    if (mapempty(t))
        t=y;
    else{
        x=t;

        while(1){
            d=x;
            if (x->k==k) return t;
            if (k<x->k){
                if (mapempty(x->left)){
                    x->left=y;
                    y->parent=x;
                    d=x;
                    while (d!=NULL){
                        d->count=0;
                        if (!mapempty(d->right))
                            d->count += d->right->count+1;
                        if (!mapempty(d->left))
                            d->count +=d->left->count+1;
                        d=d->parent;
                    }
                    break;
                }
                x=x->left;
            }
            else{
                if (mapempty(x->right)){
                    x->right=y;
                    y->parent=x;
                    d=x;
                    while (d!=NULL){
                        d->count=0;
                        if (!mapempty(d->right))
                            d->count += d->right->count+1;
                        if (!mapempty(d->left))
                            d->count +=d->left->count+1;
                        d=d->parent;
                    }
                    break;
                }
                x=x->right;
            }


        }
    }
    return  t;
}


void freet(searchtree* t){
    if (!mapempty(t)){
        freet(t->right);
        freet(t->left);
        free(t);
    }
}

searchtree* ReplaceNode(searchtree* t, searchtree*  x, searchtree*  y) {
    if (x == t) {
        t = y;
        if (y != NULL)
            y->parent = NULL;
    }
    else {
        searchtree *p = x->parent,*d;
        d=x;
        if (y != NULL) {
            y->parent = p;

        }
        if (p->left == x)
            p->left = y;
        else
            p->right = y;

        /*while (d!=NULL){
            if (mapempty(d->left) && mapempty(d->right))
                d->count=0;
            else if (mapempty(d->right))
                d->count = d->left->count+1;
            else if (mapempty(d->left))
                d->count=d->right->count+1;
            else
                d->count=d->left->count+ d->right->count +2;
            d=d->parent;
        }*/
    }


    return  t;
}

searchtree* delete(searchtree *t, int k) {
    searchtree* x = descend(t, k);
    searchtree* d;
    if (x == NULL)
        return;

    if (mapempty(x->left) && mapempty(x->right))
        t = ReplaceNode(t, x, NULL);

    else if (mapempty(x->left))
        t= ReplaceNode(t, x, x->right);

    else if (mapempty(x->right))
        t = ReplaceNode(t, x, x->left);

    else {
        searchtree * y = succ(x);

        t = ReplaceNode(t, y, y->right);
        x->left->parent = y;
        y->left = x->left;

        if (x->right != NULL)
            x->right->parent = y;

        y->right = x->right;
        t = ReplaceNode(t, x, y);
        d=x;

        while (d!=NULL){
            d->count=0;
            if (!mapempty(d->right))
                d->count += d->right->count+1;
            if (!mapempty(d->left))
                d->count +=d->left->count+1;
            d=d->parent;
        }

    }

    free(x);

    return  t;
}




int main() {

    char c[10],s[10];
    searchtree *t;
    t = NULL;

    int n, i = 0, x;


    scanf("%d", &n);
    while (i < n) {
        scanf("%s", c);

        switch (c[0]) {
            case 'I':
                scanf("%d%s", &x, s);
                t = insert(t, x, s);
                break;

            case 'S':
                scanf("%d", &x);
                puts(searchrank(t, x)->v);
                break;

            case 'D':
                scanf("%d", &x);
                t = delete(t, x);
                break;

            case 'L':
                scanf("%d", &x);
                puts(mapsearch(t, x));
                break;
        }

        i++;
    }
    freet(t);
    return 0;

}
