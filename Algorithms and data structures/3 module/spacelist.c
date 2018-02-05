#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>
//ранги элементов через список с пропусками
typedef struct spacelist spacelist;

struct  spacelist{
    spacelist *next[13];
    char data[100];
    int k;
    int span[13];
};

spacelist *initspacelist(int l, int k, char *v) {
    spacelist *p = malloc(sizeof(spacelist));
    int i;
    for (i=0;i<13;i++) {
        p->next[i] = NULL;
        p->span[i]=0;
    }
    return p;
}

spacelist *createlist(int l) {
    spacelist *p = initspacelist(l, l, 0);
    return p;
}

spacelist* succ(spacelist* l){
    return  l->next[0];
}

int skip(spacelist *l, int key, spacelist **p, int * mr) {
    int r = 0;
    spacelist *x = l;
    int i = 12;
    while (i >= 0) {
        while (x->next[i] != NULL && x->next[i]->k < key) {
            r += x->span[i];
            x = x->next[i];
        }
        if(mr!=NULL)mr[i] = r-1;
        p[i] = x;
        i--;
    }
    return r;
}

char* lookup(spacelist *l, int k){
    spacelist **p = calloc(13, sizeof(spacelist *));
    skip(l,k,p,0);
    spacelist *ver;
    ver = succ(p[0]);
    free(p);
    return  ver->data;
}

void insert(spacelist *l, int k, char* s){
    spacelist **p = calloc(13, sizeof(spacelist *));
    int *x = calloc(13, sizeof(int)),i=0;
    skip(l,k,p,x);
    spacelist *ver =  calloc(1, sizeof(spacelist));
    ver->k=k;
    strcpy(ver->data,s);
    ver->span[0]=1;

    for (i=0;i<1;i++){
        int y = p[i]->span[i];
        ver->span[i]=1;
        p[i]->span[i]=1;
        ver->next[i]=p[i]->next[i];
        p[i]->next[i]=ver;
    }
    int r = rand();

    for(; i<13 && r % 2; i++, r/=2){
        int d = p[i]->span[i];
        ver->span[i] = d ? d - (x[0] - x[i]) : 0;
        p[i]->span[i] = x[0] - x[i] + 1;
        ver->next[i] = p[i]->next[i];
        p[i]->next[i] = ver;
    }
    for(; i<13&& p[i]->span[i]; i++){
        p[i]->span[i]++;
    }
    free(p);
    free(x);
}


int  mapempty(spacelist* l){
    if (l->next[0]==0) return 1;
    return  0;
}

int rank(spacelist *l, int k) {
    spacelist **p = calloc(13, sizeof(spacelist*));
    int span = skip(l, k, p,0);
    free(p);
    return span;
}

void Delete(spacelist *l, int key) {
    spacelist **p = calloc(13, sizeof(spacelist*));
    skip(l, key, p, NULL);
    spacelist * x = succ(p[0]);
    int i;
    for(i=0; i<13; i++){
        if(p[i]->next[i] == x){
            p[i]->next[i] = x->next[i];
            p[i]->span[i] += x->span[i] - 1;
        }else if(p[i]->span[i]){
            p[i]->span[i]--;
        }
    }
    free(p);
    free(x);
}

void Deletef(spacelist *l) {
    spacelist * x = l->next[0];
    int i;
    for(i=0; i<13 && l->next[i]; i++){
        l->next[i] = x->next[i];
    }
    free(x);
}
void freel(spacelist *l) {
    while(!mapempty(l))
        Deletef(l);
}




int main() {
    int n, i, k;
    char cmd[7], str[100];
    scanf("%d", &n);
    spacelist *l = createlist(0);
    for (i = 0; i < n; i++) {
        scanf("%s%d", cmd, &k);
        switch (cmd[0]) {
            case 'I': {
                scanf("%s", str);
                insert(l, k, str);
            }
                break;
            case 'D': {
                Delete(l, k);
            }
                break;
            case 'L': {
                printf("%s\n", lookup(l, k));
            }
                break;
            case 'R' : {
                printf("%d\n", rank(l, k));
            }
        }
    }

    freel(l);
    free(l);

    return 0;
}
