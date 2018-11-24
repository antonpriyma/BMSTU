#include <iostream>
#include <vector>
#include <string>
#include "OwnIterator.h"
#include "Number.h"
#include "PrefixIterator.h"

using namespace std;

typedef PrefixIterator iter;

int main() {
    /*
    std:string in;
    cin>>in;
    String s(in);
    PrefixIterator iter=s.begin();
    *iter='{';
    *iter++;
    *iter='{';

    for (iter=s.begin();!iter.is_end();iter++) {
        cout << *iter;
    }*/

    vector<string> sample(3);
    for(int i=0;i<3;i++){
        string s;
        cin>>s;
        sample[i]=s;
    }

    StringSeq seq(sample);
    PrefixIterator prefixIterator=seq.begin();
   // cout << *prefixIterator;
    *prefixIterator="Jo"а;
    for(prefixIterator=seq.begin();!prefixIterator.is_end();prefixIterator++){
        cout<<*prefixIterator;
    }


    return 0;
}