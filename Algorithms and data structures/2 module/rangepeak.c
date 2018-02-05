#include <stdio.h>
#include <stdlib.h>
#include  <string.h>

int max(int a, int b)
{
    if (a>b)
        return a;
    else
        return b;
}

int min(int a, int b)
{
    if (a<b)
        return a;
    else
        return b;
}

void update(int *tree, int *a,int idx, int val, int v, int tl, int tr) {

    if (idx <= tl && tr <= idx) {
        a[idx] = val;
        tree[v] = val;
        return;
    }


    if (tr < idx || idx < tl) {
        return;
    }


    int tm = (tl + tr) / 2;
    update(tree, a ,idx, val, (v * 2)+1,     tl,     tm);
    update(tree, a,idx, val, v * 2 + 2, tm + 1, tr);
    tree[v] = tree[v * 2 + 2] + tree[v * 2 + 1];
}

void build_tree(int v, int* t, int tl, int tr,int* tree)
{
    int tm;
    if (tl==tr)
     tree[v]=t[tl];
    else
    {
        tm = (tr+tl)/2;
        build_tree((v*2)+1,t,tl,tm,tree);
        build_tree((v*2) + 2,t, tm+1,tr,tree);
        tree[v] = tree[(v*2)+2] + tree[(v*2) + 1];
    }


}

int get_sum(int* tree,int l, int r, int v, int tl, int tr) {

    if (l <= tl && tr <= r) 
        return tree[v];
    
   if (tr < l || r < tl) 
        return 0;
    
    int tm = (tl + tr) / 2;
    return get_sum(tree,l, r, (v * 2)+1,     tl,     tm) + get_sum(tree,l, r, (v * 2) + 2, tm + 1, tr);
}

//int get_max(int* tree,int l, int r, int v, int tl, int tr) {
//    if (l <= tl && tr <= r)  {
//        return tree[v];
//    }
//    if (tr < l || r < tl) {
//        return 0;
//    }
//    int tm = (tl + tr) / 2;
//    if (tm>=r)
//        return get_max(tree,l, r, (v * 2)+1,     tl,     tm);
//    else if (l>tm)
//        return get_max(tree,l, r, (v * 2) + 2, tm + 1, tr);
//    return max(get_max(tree,l, r, (v * 2)+1,     tl,     tm),
//         get_max(tree,l, r, (v * 2) + 2, tm + 1, tr));
//}

int main()
{
    char c;
    int p;
    char s[4];
    int nop,n,l,r,i,j;
    scanf("%d", &n);
    int *tree,*t,*buf;
    t = (int *)malloc(n*sizeof(int));
    tree = (int *)malloc(4*n*sizeof(int));
    buf = (int *)malloc(n*sizeof(int));

    for (i=0;i<n;i++)
    {
        scanf("%d", &t[i]);
        buf[i]=0;
    }
    
    if (n==1)
        buf[0]=1;
    else
    {
        if (t[0]>=t[1])
            buf[0]=1;
        if (t[n-1]>=t[n-2])
            buf[n-1]=1;
        for (i=1;i<n-1;i++)
            if (t[i]>=t[i-1] && t[i]>=t[i+1])
                buf[i]=1;
    }
    
    scanf("%d", &nop);
    i=0;
    getchar();
    build_tree(0,buf,0,n-1,tree);
    
    while (i<nop)
    {

        for (j=0;j<3;j++)
        {
            c = getchar();
            s[j]=c;
        }
        
        if (s[0]=='P')
        {
            c = getchar();
            s[4]=c;
        }
        
        scanf("%d%d", &l , &r);
        if (s[0]=='P')
        {
            if (n==1) 
                p = 1;
            else
                p=get_sum(tree,l,r,0,0,n-1);
            printf("%d\n", p);
        }
        
        if (s[0]=='U')
        {
            t[l]=r;
            if (n==1)
                buf[0]=0;
            else
            {
                for (j=l-1;j<l+2;j++)
                {
                    if (j>-1 && j<n)
                    {
                        if (j==0)
                    {
                    if (t[j]>=t[j+1])
                        update(tree,buf,0,1,0,0,n-1);
                    else
                        update(tree,buf,0,0,0,0,n-1);
                    }

                    if (j==n-1)
                    {
                        if (t[j]>=t[j-1])
                            update(tree,buf,n-1,1,0,0,n-1);
                        else
                            update(tree,buf,n-1,0,0,0,n-1);
                    }

                    if (j!=n-1 && j!=0)
                        if (t[j]>=t[j-1] && t[j]>=t[j+1])
                            update(tree,buf,j,1,0,0,n-1);
                        else
                            update(tree,buf,j,0,0,0,n-1);
                    }
                }
            }
        }
        getchar();
        i++;
    }

    free(t);
    free(tree);
    free(buf);
    return 0;
}






