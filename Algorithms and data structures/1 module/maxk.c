#include <stdio.h>
#include <stdlib.h>
//Максимальная сумма подряд идущих элемнтов
int main()
{

    int n,i,j,k;
    float  max,cur;
    scanf("%d",&n);
    cur=0;
    max = 0;
    float a[n];
    
    for(i=0;i<n;i++){
        scanf("%f",&a[i]);
    }
    scanf("\n%d", &k);

    for(i=0;i<n;i++){
        cur += a[i];
        if (i>=k-1){
            if (cur>max){
                 max=cur;
                }
                cur -= a[i-k+1];
            }
    }
    printf("%.0f", max); 
    return 0;
}