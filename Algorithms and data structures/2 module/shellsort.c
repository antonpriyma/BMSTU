void shellsort(unsigned long nel,
        int (*compare)(unsigned long i, unsigned long j),
        void (*swap)(unsigned long i, unsigned long j))
{
         int a[nel];
    int k,f1=1,f2=1,i1=0,buf=1;
    a[1]=1;
    a[2]=1;

    for (i1=3;buf>0;i1++)
    {
     k=f2;
     f2 +=f1;
     f1 = k;
     a[i1]=f2;
     buf = nel - f2 ;
    }

    int i=0, j=0, step;
        i1--;
    for (step = a[i1]; i1 >= 1; step= a[i1--])
        for (i = step; i < nel; i++)
        {
            for (j = i; j >= step; j -= step)
            {
                if ((compare(j,j-step))==-1)
                    swap(j,j-step);
                else
                    break;
            }
        }


}

