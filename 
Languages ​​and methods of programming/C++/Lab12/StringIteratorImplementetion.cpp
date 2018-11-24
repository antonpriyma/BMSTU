//
// Created by anton on 26.05.18.
//


#include "Number.h"
#include <vector>
/*
Character::Character(char p, int k) {
    value=p;
    numb=k;
}

Character& Character::operator=(char d) {
    father->x[numb].setValue(d);
    if (value=='{'){
        father->x[neibour].setValue('}');
    }else if (value=='['){
        father->x[neibour].setValue(']');
    }else if (value=='('){
        father->x[neibour].setValue(')');
    }

    if (value=='}'){
        father->x[neibour].setValue('{');
    }else if (value==']'){
        father->x[neibour].setValue('[');
    }else if (value==')'){
        father->x[neibour].setValue('(');
    }
}

void Character::setNeibour(int i) {
    this->neibour=i;
}

void Character::setValue(char c){
    this->value=c;
}

void Character::setFather(StringSeq *x) {
        father=x;
}

int Character::getNumb() {
        return numb;
}

StringSeq::StringSeq(std::string s) {
    for (int i=0;i<s.length();i++){
        Character p(s.at(i),i);
        p.setFather(this);
        x.push_back(p);
    }
    findNeibours(s);
}

void StringSeq::findNeibours(std::string s) {
    stack<Character, vector<Character>> arr;
    int count = 0;
    for (int i = 0; i < s.length(); i++) {
        if (s.at(i) == '(') {
            arr.push(x[i]);
        }
        if (s.at(i) == ')') {
            x[arr.top().getNumb()].setNeibour(i);
            x[i].setNeibour(arr.top().getNumb());
            arr.pop();
        }
    }

    for (int i = 0; i < s.length(); i++) {
        if (s.at(i) == '{') {
            arr.push(x[i]);
        }
        if (s.at(i) == '}') {
            x[arr.top().getNumb()].setNeibour(i);
            x[i].setNeibour(arr.top().getNumb());
            arr.pop();

        }
    }

    for (int i = 0; i < s.length(); i++) {
        if (s.at(i) == '[') {
            arr.push(x[i]);
        }
        if (s.at(i) == ']') {
            x[arr.top().getNumb()].setNeibour(i);
            x[i].setNeibour(arr.top().getNumb());
            arr.pop();

        }
    }
}

PrefixIterator StringSeq::begin()
{
    return PrefixIterator(*this);
}

PrefixIterator StringSeq::end()
{
    return PrefixIterator(&x[x.size()-1]);
}

bool PrefixIterator::operator==(const PrefixIterator &other) const {
    return (is_default && other.is_default) ||
           (is_default && other.is_end()) ||
           (is_end() && other.is_default) ||
           (c == other.c);
}

bool PrefixIterator::operator!=(const PrefixIterator &other) const {
    return !(*this == other);
}

string & PrefixIterator::operator () {
        if (is_default || is_end()) {
            throw "not initialized iterator";
        }
        return *c;
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
    c++;
    return *this;
}
 */
