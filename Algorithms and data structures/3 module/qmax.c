#include <stdio.h>
#include <stdlib.h>

#define EMPTY 1
#define ENQ 2
#define DEQ 3
#define MAX 4

typedef struct stack1 Stack1;



char checkcommand(char *s)
{
    switch(s[0])
    {
    case 'D':
        {
            return DEQ;
            break;
        }
    case 'E':
        switch(s[1])
        {
            case 'N':
                {
                    return ENQ;
                    break;
                }
            case 'M':
                {
                    return EMPTY;
                    break;
                }
        }
    case 'M':
        {
          return MAX;
          break;
        }
    }
}

struct stack1
{
    int *data;
    int top;
    int top1;
    int cap;
    int max ;
    int maxindex;
};

void printquene1(Stack1 q)
{
    printf("  %d\n",q.cap);
    printf("  %d\n",q.top);
    printf("  %d\n",q.top1);
    printf("  %d\n",q.max);
    printf("  %d\n",q.maxindex);

}

struct stack1 *initstack(int n)
{
    struct stack1 *st = (struct stack1 *)malloc(sizeof(struct stack1));
    st->cap=n;
    st->data = (int *)malloc(n*sizeof(int));
    st->top=0;
    st->top1=0;
    st->max=-2000000000;
    st->maxindex=0;
    return st;

};

int checkempty(Stack1* s)
{
    if (s->top==s->top1)
        return 1;
    return 0;
}

int pop(Stack1 *s)
{
    int i;
    if (!(checkempty(s)))
    {
        int x;
        //printquene1(*s);
        if (s->top1 == s->maxindex)
        {
            s->max = -2000000000;
            //printquene1(*s);
            for (i=s->top1+1; i<s->top; i++)
            {
                if (s->data[i]>=s->max)
                {
                  s->max = s->data[i];
                  s->maxindex =i;
                }
            }
        }
        x = s->data[s->top1++];
        return x;
    }
    else
    {
        printf("Empty");
        return 0;
    }
}

void push(Stack1 *s, int x)
{
    s->data[s->top++]=x;
    if (x>=s->max)
    {
        s->maxindex=s->top - 1;
        s->max = x;
    }
}

int main()
{
    long flag=0,command_type=0,i=0,m=0;
    int x=0;
    Stack1 *q;
    q = initstack(1000000);
    char command[6];
    scanf("%d", &m);
    getchar();
    for (i=0;i<m;i++)
    {
        scanf("%s", &command);
            command_type=checkcommand(command);
        switch(command_type)
        {
        case 1:
            {
              flag = checkempty(q);
              if (flag==1)
                    puts("true");
              else
                    puts("false");
              break;
            }
        case 2:
            {
                scanf("%d", &x);
                push(q,x);
                break;
            }
        case 3:
            {
                printf("%d\n", pop(q));
                break;
            }
        case 4:
            {
                printf("%d\n", q->max);
                break;
            }
        }

    }

    free(q->data);
    free(q);
    return 0;
}





