#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int max(int a, int b)
{
    if (a>b)
        return a;
    else
        return b;
}

int* Delta1(char *s, size_t n)
{
    int *u1;
    u1 = (int *)malloc(n*sizeof(int));
    int len=strlen(s),i,a=0;
    while (a<n)
    {
        u1[a]=len;
        a++;
    }

    int j = 0;
    while (j<len)
    {
        u1[s[j]]=len-j-1;
        j++;
    }
    return u1;
}

int* Suffix(char *s)
{
    int len = strlen(s);
    int *d;
    d = (int *)malloc(len*sizeof(int));
    int i=len-2,t=len-1;
    d[t]=t;

    while (i>=0)
    {
        while (t<len-1 && s[t]!=s[i])
            t = d[t+1];
        if (s[t]==s[i])
            t--;
        d[i] = t;
        i--;
    }

    return d;

}

int* Delta2(char *s)
{
    int len = strlen(s);
    int *u2;
    u2 = (int *)malloc(len*sizeof(int));
    int *d = Suffix(s),i=0;
    int t = d[0];
    while (i<len)
    {
        while (t<i)
            t=d[t+1];
        u2[i]= t+len-i;
        i++;
    }
    i=0;
    while (i<len-1)
    {
        t = i;
        while (t<len-1)
        {
            t = d[t+1];
            if (s[i]!=s[t])
                u2[t]=len-i-1;

        }
        i++;
    }

    return u2;

}

int BMSubst(char *s,char *t,int n)
{
    int *u1,*u2;
    u1 = Delta1(s,255);
    u2 = Delta2(s);
    int len = strlen(s);
    int i=0,k = len-1;
    int lent = strlen(t);

    while (k<lent)
    {
        i = len - 1;
        while (t[k]==s[i])
        {
            if (i==0)
                return k;
            i--;
            k--;
        }
   
        k += max(u1[t[k]],u2[i]);
    }
    k = strlen(t);
    return k;


}


//int SimpleBMSubst(char *s, char *t, int n)
//{
//    int k2,i,j,k,k1;
//    int *u1;
//    u1 = Delta1(s,n);
//    k = strlen(s)-1;
//    k2 = k;
//    k1 = strlen(t);
//    while (k<k1)
//    {
//       i = k2;
//       while (t[k]==s[i])
//       {
//           if (i==0)
//           {
//                    free(u1);
//                    return k;
//           }
//            i--;
//            k--;
//       }
//       k += max(u1[t[k]],k2-i+1);
//     }
//    k = strlen(t);
//    free(u1);
//    return k;

//}

int main(int argc, char **argv)
{
    char *s,*t;
    s=(char *)malloc(1000);
    t=(char *)malloc(1000);
    strcpy(t,argv[2]);
    strcpy(s,argv[1]);
    int i,count,j=0,k=strlen(s),k1=0,k2=0;

    while ((count=strlen(t))>=k)
    {
    if (j>1)
        k2=k1+1+k2;
    else
        k2=k1;
    //k1 = SimpleBMSubst(s,t,255);
    k1 = BMSubst(s,t,255);
    if (k1==strlen(t))
        break;

    if (j==0)
        printf("%d ", k1+k2);
    else
        printf("%d ", k1+k2+1);
    j++;
    for (i=0;i<count-k1-1;i++)
        t[i]=t[i+k1+1];

    t[count-k1-1]='\0';
}
    return 0;
}


