//
// Created by anton on 19.04.18.
//

#ifndef LAB7_DECLARATION_H
#define LAB7_DECLARATION_H

#include <iostream>
#include <iosfwd>
using namespace std;

class Shuffle{
    friend ostream& operator << (ostream& os, const Shuffle& shuffle){
        for(int i=0;i<shuffle.n;i++){
            os << shuffle.from[i] <<"->"<< shuffle.to[i]<<" ";
        }
        return os;
    }
public:
    int& operator[](int i);
    Shuffle operator *(Shuffle s);
    Shuffle(const Shuffle& s);
    Shuffle& operator=(const Shuffle& s);
    bool checkBackwards();
    Shuffle mult(Shuffle);
    void swap(int i,int j);
    void set(int i,int j);
    bool equal(Shuffle s);
    virtual ~Shuffle();
    int getIndex(int elem)const;

    Shuffle(int n);
private:
    int *from;
    int  *to;
    int n;
};
#endif //LAB7_DECLARATION_H
