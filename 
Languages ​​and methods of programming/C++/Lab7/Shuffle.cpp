//
// Created by anton on 19.04.18.
//
#include "Shuffle.h"

Shuffle::Shuffle(int n) {
    from = new int[n];
    to = new int[n];
    this->n=n;
    for (int i=0;i<n;i++){
        from[i]=i+1;
        to[i]=i+1;
    }
}

int& Shuffle::operator[](int i) {
    return to[i];
}

int Shuffle::getIndex(int elem)const {
    for (int i=0;i<this->n;i++){
        if (this->from[i]==elem){
            return i;
        }
    }
}

bool Shuffle::equal(Shuffle s) {
    for (int i=0;i<this->n;i++){
        if (s.to[i]!=this->to[i]){
            return false;
        }
    }
    return true;
}
bool Shuffle::checkBackwards() {
    Shuffle buf(this->n);
    for (int i=0;i<this->n;i++){
        buf.from[i]=this->to[i];
        buf.to[i]=this->from[i];
    }

    Shuffle multiply = this->mult(buf);
    Shuffle help(this->n);
    return multiply.equal(help);
}

Shuffle Shuffle::mult(Shuffle x) {
    Shuffle result(this->n);
    for (int i=0;i<this->n;i++){
        result.to[i]=x.to[this->to[i]-1];
    }
    return result;
}

Shuffle Shuffle::operator*(Shuffle s) {
    return this->mult(s);
}

Shuffle::Shuffle(const Shuffle &s) {
    from=new int[s.n];
    to=new int[s.n];
    n=s.n;
    for(int i=0;i<s.n;i++){
        to[i]=s.to[i];
        from[i]=s.from[i];
    }
}

void Shuffle::set(int i, int j) {
    to[getIndex(i)]=j;
}

Shuffle& Shuffle::operator=(const Shuffle &s) {
    if ( this != & s ) {
        int *new_a = new int [ s .n ];
        std :: copy ( s .from , s . from + s .n , new_a );
        delete []from;
        n = s .n;
        from = new_a ;
        int *new_b = new int [ s .n ];
        std :: copy ( s .to, s . to + s .n , new_a );
        delete []to;
        n = s .n;
        to = new_b ;
    }
    return *this;
}

Shuffle::~Shuffle() {
    delete []from;
    delete []to;
}

void Shuffle::swap(int i, int j) {
    int buf = this->to[i-1];
    this->to[i-1]=this->to[j-1];
    this->to[j-1]=buf;
}



