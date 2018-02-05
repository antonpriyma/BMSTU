#include <stdio.h>
#include <stdlib.h>
#include "elem.h"

struct Elem *searchlist(struct Elem *list, int k)
{
    struct Elem *k1;

    while (list->tail!= 0)
    {
        if (list->tag==INTEGER)
        {
            if (list->value.i==k)
                return list;
            else
                list = list->tail;
        }
        else
        {

            if (list->tag==FLOAT)
            {
                if (list->value.f==k)
                    return list;
                else
                list = list->tail;
            }
            else
            if (list->tag==LIST)
                if ((*list).value.list!=0)
                {
                    k1=searchlist((*list).value.list, k);
                    if (k1!=0)
                        return k1;
                    else
                        list = list->tail;
                }
                else
                    list = list->tail;
        }

    }

    //Проверка последнего элемента списка

    if (list->tag==INTEGER)
        if (list->value.i==k)
            return list;
        else
            return 0;

    if (list->tag==FLOAT)
        if (list->value.f==k)
            return list;
        else
            return 0;


    if (list->tag==LIST)
        if ((*list).value.list!=0)
        {
            k1=searchlist((*list).value.list, k);
            if (k1!=0)
                return k1;
            else
                return 0;
        }
        else
            return 0;
}