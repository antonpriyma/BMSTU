#include <stdio.h>
#include <stdlib.h>

typedef struct Prefix prefix;
typedef struct List list;

/*
struct Prefix{
    int x;
    int count;
};*/

struct List{
    int k;
    int v;
    list* next
};

void printlist(list* t){
    printf("%d %d\n", t->k, t->v);
}

list* initlist(int k, int v){
    list* elem = malloc(sizeof(list));
    elem->k=k;
    elem->v=v;
    elem->next=NULL;
}




int checkempty(list* t){
    if (t->next==NULL)
        return 1;
    else
        return 0;
}

list* listsearch(list* t,int k){
    list* x =t;
    while (x!=NULL && x->k!=k){

        x=x->next;
    }

    return x;
}

void insertbeforehead(list* t, int y){
    list* elem = malloc(sizeof(list));
    elem->k=y;
    if (y!=0)
        elem->v=0;
    else
     elem->v=1;
    elem->next=t->next;

    t->next=elem;

}

int h(int x, int n){
    int result;
    result = abs(x)%n;
    return result;
}

void inithash(int n, list **t){
    int i;
    for (i=0;i<n;i++)
    {
        t[i] = initlist(0, 0);
    }
}

list* mapsearch(list** t, int k, int n){
    list* result;
    result = listsearch(t[h(k,n)],k);
    return result;
}

void freelist(list* t){
    list* b;
    b=t;
    while (b!=NULL){
        t = t->next;
        free(b);
        b = t;
    }
}

int main(){
    /*
    int i,x,y,n;
    scanf("%d", &n);
    int a[n];
    int help=0;
    for (i=0;i<n;i++){
        scanf("%d", &a[i]);
        help ^= a[i];
    }
    printf("%d", help);
    */
    int res=0,val,n,i;
    scanf("%d", &n);
    int  a[n];// считываение кол-ва, создание массива
    list* buf;
    list* ex[n];
    inithash(n,ex);


    for (i=0;i<n;i++){
        scanf("%d", &val);
        //if (i==0 && val ==0) res++;
        if (i>0){
            a[i]=a[i-1]^val;
            if ((buf=mapsearch(ex,a[i],n))==NULL)
                insertbeforehead(ex[h(a[i],n)],a[i]);
            else
                buf->v++;
        }
        else {
            a[i] = val;
            if ((buf=mapsearch(ex,a[i],n))==NULL)
                insertbeforehead(ex[h(a[i],n)],a[i]);
            else
                buf->v++;
        }

    }

    /*for (i=0;i<n;i++){
        if ((buf =mapsearch(ex,i,n))!=NULL )
            res += buf->v*((buf->v+1)/2);

        }*/

    for (i=0;i<n;i++){
        buf = ex[i];
        while (buf!=NULL){
            res += buf->v*(buf->v+1)/2;
            buf=buf->next;
        }

    }

    for (i=0;i<n;i++)
        freelist(ex[i]);

    //for (i=0;i<n;i++)
      //  free(ex[i]);

    /*
    insertbeforehead(ex[0],1);
    printlist(ex[0]);
    list* b = listsearch(ex[0], 1);

    printf("%d %d", b->k, b->v);
    */
    printf("%d",res);






}





