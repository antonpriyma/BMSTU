#include <iostream>
using namespace std;

int main() {

    int N;
    cin>>N;
    int i=0,x=00;
    int K=16;
    int a[K];
    for(i=0;i<K;i++) {
        a[i]=0;
    }
    for(i=0;i<N;i++) {
        cin>>x;
        a[x-1]++;
    }
    for(i=0;i<16;i++) {
        int max=-1;
        int maxindex=0;
        for(int j=i;j<16;j++) {
            if(a[j]>max) {
                max=a[j];
                maxindex=j;
            }
        }
        int tmp=0;
        tmp=a[i];
        a[i]=max;
        a[maxindex]=tmp;

        if (max > 0 ) {
            cout << maxindex + 1 << " " << max << endl;
        }
    }



    return 0;
}
