#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef  struct  node node;
лексический анализ чеерез avlдерево
struct node // структура для представления узлов дерева
{
    char key[100];
    int v;
    unsigned char height;
    node* left;
    node* right;
};

node* initavl(char* k,int v){
    node* t = malloc(sizeof(node));
    strcpy(t->key,k);
    t->height = 1;
    t->left=0;
    t->right=0;
    t->v=v;
}

unsigned char height(node *t){ //работа с высотой через функцию
    if (!t) return 0;
    return  (t->height);
}

unsigned char bfactor(node *t){
    return (height(t->right)-height(t->left));
}

void fixheight(node* t)
{

    unsigned char hl = height(t->left);
    unsigned char hr = height(t->right);
    t->height = (hl>hr?hl:hr)+1;
}

node* rotateright(node* t) // правый поворот вокруг p
{
    node* q = t->left;
    //if (q->right==0)
    //   return  t;
    t->left = q->right;
    q->right = t;
    fixheight(t);
    fixheight(q);
    return q;
}

node* rotateleft(node* t) // левый поворот вокруг p
{
    node* q = t->right;
    //if (q->left==0)
    //  return  t;
    t->right = q->left;
    q->left = t;
    fixheight(q);
    fixheight(t);
    return q;
}

node* balance(node* t) // балансировка узла p
{
    fixheight(t);
    if( bfactor(t)==2 )
    {
        if( bfactor(t->right) < 0 )
            t->right = rotateright(t->right);
        return rotateleft(t);
    }
    if( bfactor(t)==-2 )
    {
        if( bfactor(t->left) > 0  )
            t->left = rotateleft(t->left);
        return rotateright(t);
    }
    return t; // балансировка не нужна
}

node* insert(node* t, char* k, int v) // вставка ключа k в дерево с корнем p
{
    if( !t )
        return initavl(k,v);

    if (v<t->v                                                                                                                                        )
        t->left = insert(t->left,k,v);
    else
        t->right = insert(t->right,k,v);
    return balance(t);
}

node* removemin(node* t) // удаление узла с минимальным ключом из дерева p
{
    if( t->left==0 )
        return t->right;
    t->left = removemin(t->left);
    return balance(t);
}

int  find(node* t, char* k,int v) { //поиск элемента
    if (!t)
        return -1;
    if (strcmp(k,t->key)==0)
        return t->v;
    int x,y;
    /*if (v <= t->v)
        return find(t->left, k,v);
    else
        return find(t->right, k,v);*/
    x = find(t->left,k,v);
    y = find(t->right,k,v);
    if (x==-1) return y;
    return x;

}

int spec(char s){
    if (s!=' ' && s!='(' && s!=')' && s!='+' && s!='-' && s!='*' && s!='/')
        return 0;
    return 1;
}

int digit(char *s){
    if (s<='9' && s>='0') return 1;
    return 0;
}

int word(char *s){
    if (s>=97) return 1;
    return 0;
}

void freeavl(node *t){
    if (!t) return;
    freeavl(t->right);
    freeavl(t->left);
    free(t);
    t=0;
}

int main() {
    int cur=0,k=0,i,n,count=0;
    scanf("%d\n", &n);
    node *t = 0;
    char s[n],c=0,buf[100],flag=0;

    for (i=0;i<n&&c!='\n';i++){

        if (!flag)
            scanf("%c", &c);
        s[i]=c;
        flag = 0;
        switch(c){
            case ' ':
            case '\n':
                break;
            case '(':
                printf("SPEC 4\n");
                break;
            case ')':
                printf("SPEC 5\n");
                break;
            case '+':
                printf("SPEC 0\n");
                break;
            case '-':
                printf("SPEC 1\n");
                break;
            case '*':
                printf("SPEC 2\n");
                break;
            case '/':
                printf("SPEC 3\n");
                break;
            case '0': case '1':
            case '2': case '3':
            case '4': case '5':
            case '6': case '7':
            case '8': case '9':
                if (i>0)
                    if (!word(s[i-1]))
                        printf("CONST ");
                if (i==0) printf("CONST ");
                while (digit(c)) {
                    printf("%c", c);
                    s[i++]=c;
                    c=getchar();
                }
                i--;
                printf("\n");
                flag=1;
                break;
            default:
                flag=1;
                count=0;
                s[i]=c;
                buf[count]=c;
                while ((c=getchar())>47 || digit(c)){
                    s[++i]=c;
                    buf[++count]=c;
                }
                buf[count+1]='\0';
                if ((cur=find(t,buf,k))==-1){
                    t = insert(t,buf,k);
                    printf("IDENT %d\n", k++);
                }
                else
                    printf("IDENT %d\n", cur);
                break;

        }

    }

    freeavl(t);
    t=0;


}
