#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

typedef struct Elem elem;



int compare(char *a, char *b)
{
    int len1=strlen(a),len2=strlen(b);
    if (len1>len2) return 1;
    return 0;
}

struct Elem
{
    elem *next;
    char *v;
};



void insertafter(elem *x, elem *y)
{
    elem *z;
    while (x!=NULL)
    {
        z=x;
        x=x->next;
    }
    z->next=y;
}


elem* initlist(char* str)
{
        elem* tmp = calloc(1,sizeof(elem));
        tmp->v=calloc(100,sizeof(char));
        strcpy(tmp->v,str);
        tmp->next=NULL;
        return tmp;
}

void bsort(elem *start)
{
    elem* p_list;
    elem* pp_list;
    for (p_list = start->next; p_list != NULL; p_list = p_list->next)
        for (pp_list = start->next; pp_list->next != NULL; pp_list = pp_list->next)
            if (compare(pp_list->v,pp_list->next->v)==1)
                {
                    char* tmp;
                    tmp = pp_list->v;
                    pp_list->v = pp_list->next->v;
                    pp_list->next->v = tmp;
                }
}


int main()
{
        long int i, n,len,count=0;
        char *s,buf[1000];
        bool flag = false;
        s= (char *)malloc(10000);

        gets(s);
        len = strlen(s);

        elem *list=initlist(""), *y,*p;
        p = list;
        for (i=0;i<len;i++)
            if (flag == false && s[i]==' ')
                continue;
            else if (flag == false && s[i]!=' ')
            {
                count=0;
                flag = true;
                buf[count]=s[i];
            }
            else if (flag == true && s[i]!=' ')
                buf[++count]=s[i];
            else if (flag == true && s[i]==' ')
            {
                buf[++count]='\0';
                flag = false;
                y = initlist(buf);
                buf[0]='\0';
                insertafter(list, y);
            }

        buf[++count]='\0';
        y = initlist(buf);
        insertafter(list,y);

        /*
        for (i = 0; i < n; i++)
        {
                y=initlist();//buf
                scanf("%d", &(y->v));

        }
        */

        bsort(p);
        y=p->next;
        while (y != NULL)
        {
            printf("%s ", y->v);
            p = y->next;
            free(y);
            y=p;
        }
        free(list->v);
        free(list);
        return 0;
}
