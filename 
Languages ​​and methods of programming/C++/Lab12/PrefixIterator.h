//
// Created by anton on 26.05.18.
//

#ifndef LAB12_PREFIXITERATOR_H
#define LAB12_PREFIXITERATOR_H

#include <string>
#include <iterator>
#include <stack>
#include <vector>
#include "Number.h"

using namespace std;

class PrefixIterator;
class StringSeq;


class String{
    friend ostream& operator<< (ostream& os, String n) {
        os << n.s;
        return os;
    }

private:
    string s;
    StringSeq *father;
    int k;
    friend class  StringSeq;


public:
    String& operator=(string d) ;

    String(string s){
        this->s=s;
    }
    String(){
        s="";
    }


};


class StringSeq {
private:
    friend class PrefixIterator;
    friend class String;
    vector<String> x;
public:


    PrefixIterator begin();
    //PrefixIterator end();

    StringSeq(vector<string> x);
    String findPreix(int x);
};



class PrefixIterator:
        public std::iterator<
                std::forward_iterator_tag,
                String
        >
{
private:
    String *c;
    StringSeq *batya;
    int count;
    int length;
    bool is_default;

public:
    PrefixIterator(): is_default(true) {}
    PrefixIterator(StringSeq &n): is_default(false) {c=&n.x[0];count=0;length=n.x.size()-1;batya=&n;}
    //PrefixIterator(Character *n): is_default(false) {c=n;}

    bool operator == (const PrefixIterator &other) const;
    bool operator != (const PrefixIterator &other) const;
    String operator*();
    PrefixIterator& operator++ ();
    PrefixIterator operator++ (int);
    bool is_end() const { return count>=length;}
};

#endif //LAB12_PREFIXITERATOR_H
