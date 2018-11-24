//
// Created by anton on 26.05.18.
//


#include "Number.h"
#include "PrefixIterator.h"
#include <vector>


StringSeq::StringSeq(vector<string> s) {
    for (int i=0;i<s.size();i++){
        String l(s[i]);
        l.k=i;
        x.push_back(l);
    }
}
String& String::operator=(string d) {
        String s = father->findPreix(k);
        father->x[k].s.replace(0,s.s.length(),d);

}
String StringSeq::findPreix(int x) {
    String s;
    String s1=this->x[x];
    String s2=this->x[x+1];
    int len = min(s1.s.length(),s2.s.length());
    for (int i=0;i<len;i++){
        if (s1.s.at(i)==s2.s.at(i)){
            s.s+=s1.s.at(i);
        } else{
            break;
        }
    }
    return s;
}

PrefixIterator StringSeq::begin()
{
    return PrefixIterator(*this);
}
/*
PrefixIterator StringSeq::end()
{
    return PrefixIterator(&x[x.size()-1]);
}*/

bool PrefixIterator::operator==(const PrefixIterator &other) const {
    return (is_default && other.is_default) ||
           (is_default && other.is_end()) ||
           (is_end() && other.is_default) ||
           (c == other.c);
}

bool PrefixIterator::operator!=(const PrefixIterator &other) const {
    return !(*this == other);
}

String PrefixIterator::operator*() {
    int p = this->count;
    return batya->findPreix(p);
}

PrefixIterator PrefixIterator::operator++(int) {
    PrefixIterator tmp(*this);
    operator++();
    count++;
    return tmp;
}


PrefixIterator& PrefixIterator::operator++() {
    if (is_default) throw "not initialized iterator";
    if (is_end()) throw "iterator overflow";
    return *this;
}
