#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void heapify(int i, char **list,char heapSize,size_t width)
{
    int rightChild,largestChild,leftChild;
    char* block= malloc(width+1);
    block[0]='\0';

    for (; ; )
    {
        leftChild = 2 * i + 1;
        rightChild = 2 * i + 2;
        largestChild = i;

        if (leftChild < heapSize && compare(list[largestChild] , list[leftChild]))
        {
            largestChild = leftChild;
        }

        if (rightChild < heapSize && compare(list[largestChild],list[rightChild]))
        {
            largestChild = rightChild;
        }

        if (largestChild == i)
        {
            break;
        }

        memcpy(block,list[i],width+1);
        memcpy(list[i],list[largestChild],width+1);
        memcpy(list[largestChild],block,width+1);
        i=largestChild;
    }
    free(block);

}//упорядочивание дереева



void hsort(void *base, size_t nel, size_t width,int (*compare)(const void *a, const void *b))
{
    char **b = (char *)base;
    int i,n1=nel;

    for(i=nel/2; i>=0; i--)
        heapify(i,b,nel, width);

    while( nel > 1 ) //Пока в пирамиде больше одного элемента
    {

        char *firstelement = malloc(width+1);
        firstelement[0]='\0';
        //Меняю местами корневой элемент и отделённый:
        memcpy(firstelement,b[0],width+1);
        memcpy(b[0],b[nel-1],width+1);
        memcpy(b[nel-1],firstelement,width+1);
        nel -=1;

        //Просеиваю новый корневой элемент:
        heapify(0, b, nel, width);
        free(firstelement);
    }

    for (i=0;i<n1;i++)
    {
        puts(b[i]);
    }

}


int compare(const char *a, const char *b)
{
    char i,*q=a,*w=b;
    int e=0,r=0;

    for(i=0;i<strlen(q);i++)
    if(q[i]=='a')
        e++;

    for(i=0;i<strlen(w);i++)
    if(w[i]=='a')
        r++;
    if(e<r)
        return 1;
    else
        return 0;
}

int main()
{
    int j,max1=0,n,i ;
    scanf("%d", &n);
    getchar();
    char *s[n],c,*buff=NULL;

    for (j=0;j<n;j++){
        i=0;
        buff=(char *)malloc(100);
        while ((c=getchar())!='\n')
        {
            i++;
            *(buff+i-1)=c;//считывание строки
        }

        s[j]=(char *)malloc(100);
        *(buff+i) ='\0';
        strcpy(s[j],buff);//присваиванию элементу массива строк
        if (strlen(s[j])>max1)
            max1 = strlen(s[j]);
        if (j==0)
            s[j][max1+1]='\0';

        free(buff);
    }//считывание строк

    hsort(s,n,max1,*compare);

    for (i=0;i<n;i++)
        free(s[i]);
    return 0;
}

