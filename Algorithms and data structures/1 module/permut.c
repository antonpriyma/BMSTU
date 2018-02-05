#include <stdio.h>
#include <stdlib.h>

int main()
{
    struct h{
        int elem;
        int count;
    };
    int j,i,x,k,t,m;
    struct h a[8],b[8];
    t =0;
    m=0;

    for(i=0;i<8;i++){
        a[i].elem = 0;
        a[i].count = 0;
        b[i].elem = 0;
        b[i].count = 0;
    }

    for(i=0;i<8;i++){
        scanf("%d", &x);
        k=0;
        for(j=0;j<8;j++){
            if (a[j].elem == x){
                a[j].count++;
                k=1;
                break;
            }
        }
        if (k == 0){
            a[t].elem = x;
            a[t].count = 1;
            t++;
        }
    }

    t=0;

    for(i=0;i<8;i++){
        k=0;
        scanf("%d", &x);

        for(j=0;j<8;j++){
            if (b[j].elem == x){
                b[j].count++;
                k=1;
                break;
            }
        }
        if (k == 0){
            b[t].elem = x;
            b[t].count = 1;
            t++;
        }
    }

    for (i=0; i<8;i++){
        for(j=0;j<8;j++){
            if ((a[i].elem == b[j].elem) && (a[i].count == b[j].count)){
                    m++;
                    break;
            }
        }
    }

    if (m == 8) {
        printf("yes");
    }
    else{
        printf("no");
    }




    return 0;
}