#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define M 26

typedef  struct Borvrtx bor;


void input(char *s)
{
    int i=0;
    short c;
    while ((c=getchar())!=' ' && c!='\n')
        s[i++]=c;
    getchar();
}

struct Borvrtx{
    bor* argc[M];
    short v;
    bor* parent;
    char word;
};

bor* initbor(){
    int i;
    bor* t;
    t = calloc(1, sizeof(bor));
    t->v=0;
    t->word=0;
    t->parent=NULL;
    return t;
}

int boremty(bor* t){
    int i;
    if (t->v!=0) return 0;
    for (i=0;i<M;i++)
        if (t->argc[i]!=NULL) return 0;
    return  1;
}

bor* descend(bor* t,bor* x, char* s,int len,int* f){
    bor* y;
    int k;
    int i;
    *f = 0;
    x = t;
    for (i=0;i<len;i++){
         k = s[i]-97;
         if (x->argc[k]!=0)
            y = x->argc[k] ;
          else {
          *f=i;
          return x;
          }
        x = y;
        *f=i;
    }
    *f=i;
    return  x;
}

int mapsearch(bor *t, char* s){
    bor* x;
    int len = strlen(s);
    int i;
    x = descend(t,x,s,len,&i);

    if (i==len) return  x->v;
    return  0;
}

void insert(bor* t, char*s){
    bor* x;
    bor* y;
    int k,len = strlen(s);
    int i;
    x = descend(t,x,s,len,&i);

    if ((x->word == 1) && (i==len)) return;
    while (i<len){
         y  = initbor();
         k=s[i]-97;
         x->argc[k]=y;
         y->parent = x;

         x = x->argc[k];

         i++;

    }
       x->word=1;

    while (x->parent!=NULL){
        x->v++;
        x=x->parent;

    }
}

void delete(bor* t, char* s){
    bor *y,*x;
    int j,len = strlen(s);
    int i;
    x = descend(t,x,s,len,&i);
    
    if (i!=len) return;
    x->v--;
     while (x->parent!=NULL){
              j=0;
         while (j<M && x->argc[j]==NULL)
             j++;
         if (j<M) {
             while (x->parent!=NULL){
                 x=x->parent;
                 x->v--;
             }

             break;
         }
         y=x->parent;
         i--;
         free(y->argc[s[i]-97]);
         y->argc[s[i]-97]=0;
         y->v--;
         x=y;

     }

}


void freebor(bor *t){
    int i;
    for (i=0;i<M;i++)
        if (t->argc[i]!=NULL)
        freebor(t->argc[i]);

    free(t);
}


int main() {
    int k;
    bor *t;
    char c;
    t = initbor();
    int i, n;
    char cmd[10], buf[100000];
    scanf("%d", &n);

    for (i = 0; i < n; i++) {
        	scanf("%s%s", cmd, buf);
		switch (cmd[0]) {
		case 'I':
			insert(t, buf);
			break;
		case 'D':

            

			delete(t, buf);
			break;
		case 'P':
            
			printf("%d\n", mapsearch(t, buf));
			break;
		};
	}

    freebor(t);
    return 0;
}
