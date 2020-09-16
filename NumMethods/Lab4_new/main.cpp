#include <iostream>
using namespace std;

int main(){
    int N;
    cin>>N;
    int maxD=-1;
    int x=0,i=0,max=-1;
    int a[8];
    for(i=0;i<8;i++) {
        cin>>a[i];
    }

    for(i=8;i<N;i++){
        cin>>x;
        if(a[8%i]>max) {
            max=a[8%i];
        }

        if ((max*x) >  maxD) {
            maxD = (max*x);
        }

        a[8%i] = x;
    }


    cout<<maxD;

    return 0;
}




