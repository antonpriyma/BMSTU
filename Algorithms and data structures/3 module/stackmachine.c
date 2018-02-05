#include <stdio.h>
#include <stdlib.h>


int min(int a, int b)
{
    if (a>b)
        return b;
    else
        return a;
}

int max(int a,int b)
{
      if (a>b)
        return a;
    else
        return b;
}
struct stack1
{
    int *data;
    int top;
    int cap;
};

int checkcommand(char *s)
{
    switch(s[0])
    {
        case 'C': return 1;
        case 'A': return 2;
        case 'S':
            {
                switch(s[1])
                {
                    case 'U': return 3;
                    case 'W': return 10;
                }
            }

        case 'M':
            {
                switch(s[1])
                {
                    case 'U': return 4;
                    case 'A': return 6;
                    case 'I': return 7;
                }
            }
        case 'D':
            {
                switch(s[1])
                {
                    case 'U': return 9;
                    case 'I': return 5;
                }
            }
        case 'N': return 8;


    }
}

struct stack1 *initstack()
{
    struct stack1 *st = (struct stack1 *)malloc(sizeof(struct stack1));
    st->cap=1000000;
    st->data = (int *)malloc(1000000*sizeof(int));
    st->top=-1;

    return st;

};
int checkempy(struct stack1 *s)
{
    if (s->top==-1)
        return 1;
    return 0;
}

void push(struct stack1 *s, int x)
{
    s->data[++s->top]=x;
}

int pop(struct stack1 *s)
{
    int x;
    if (checkempy(s))
    {
        printf("Stack is empy");
        return 0;
    }
    x=s->data[s->top];
    s->top--;
    return x;
}

int main()
{
    int i,m,x,x1;
    struct stack1 *s;
    s = initstack();
    scanf("%d\n", &m);


    char command[5];
    for (i=0;i<m;i++)
    {
    scanf("%s", &command);
    int command_type = checkcommand(command);

    switch (command_type)
    {
    case 1:
        {
            scanf("%d", &x);
            push(s,x);
            break;
        }
    case 2:
        {
            x = pop(s);
            x1 = pop(s);
            push(s, x+x1);
            break;
        }
    case 3:
        {
            x = pop(s);
            x1 = pop(s);
            push(s,x-x1);
            break;
        }
    case 4:
        {
            x = pop(s);
            x1 = pop(s);
            push(s,x*x1);
            break;
        }
    case 5:
        {
            x = pop(s);
            x1 = pop(s);
            push(s,x/x1);
            break;
        }
    case 6:
        {
            x = pop(s);
            x1 = pop(s);
            push(s,max(x,x1));
            break;
        }
    case 7:
        {
            x = pop(s);
            x1 = pop(s);
            push(s,min(x,x1));
            break;
        }
    case 8:
        {
            x = pop(s);
            x *= -1;
            push(s,x);
            break;
        }
    case 9:
        {
            x = pop(s);
            s->top++;
            push(s,x);
            break;
        }
    case 10:
        {
            x = pop(s);
            x1 = pop(s);
            push(s,x);
            push(s,x1);
            break;
        }

    }
    }

    printf("%d", s->data[s->top]);
    free(s);

}
