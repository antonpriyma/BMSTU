//
// Created by anton on 11.05.18.
//

#ifndef UNTITLED4_POLARCURVE_H
#define UNTITLED4_POLARCURVE_H
#include <vector>
#include <iostream>

using namespace std;



template <class T>
class Coord{
private:
    T length;
    T a;
public:
    T getA(){
        return a;
    }

    T getlength(){
        return length;
    }

    void setA(T elem){
        a=elem;
    }

    void setlength(T elem){
        length=elem;
    }

    template <typename T1> friend ostream& operator<<(ostream &out,Coord<T1> &x);
};

template <typename T>
ostream& operator<<(ostream &out,Coord<T> &x){
    out<<"length: "<< x.getlength()<< " alpha: "<<x.getA();
    return out;
}




template <typename T>
class PolarCurve{
    template <typename T1> friend const PolarCurve<T1> operator *(const T1 m,const PolarCurve<T1> o);
private:
    std::vector<T (*)(T)> funcs;
    T multiplier;
public:

    //Конструктор от 1 функции
    PolarCurve(T (*g)(T));

    //Конструктор от несольких фукнций
    PolarCurve(vector<T (*)(T)> arg);


    PolarCurve(vector<T (*)(T)> arg, T k);

    //Cумма
    PolarCurve operator +(PolarCurve o);

    //Умножение функций на число
    PolarCurve operator *(T k);

    //Вычисление значения
    Coord<T> operator () (T arg);
};

template <typename T>  const PolarCurve<T> operator *(const T m,const PolarCurve<T> o){
    return PolarCurve<T>(o.funcs,m);
}

template <typename T>
PolarCurve<T>::PolarCurve(T (*g)(T)) {
    funcs.push_back(g);
    multiplier=1;
}

template <typename T>
PolarCurve<T>::PolarCurve(vector<T (*)(T)> arg) {
    funcs = arg;
    multiplier=1;
}

template <typename T>
PolarCurve<T>::PolarCurve(vector<T (*)(T)> arg, T k) {
    funcs=arg;
    multiplier=k;
}

template <typename T>
PolarCurve<T> PolarCurve<T>::operator*(T k) {
    multiplier=k;
    return PolarCurve(this->funcs,multiplier);
}

template <typename T>
Coord<T> PolarCurve<T>::operator()(T arg) {
    {
        T a = arg;
        T l=0;
        for (int i=0;i<funcs.size();i++){
            l+=funcs[i](arg)*multiplier;
        }
        Coord<T> result;
        result.setlength(l);
        result.setA(a);
        return result;
    };
}

template <typename T>
PolarCurve<T> PolarCurve<T>::operator+(PolarCurve o) {
    vector<T (*)(T)> sum;
    for (int i=0;i<this->funcs.size();i++){
        sum.push_back(this->funcs[i]);
    }

    for (int i=0;i<o.funcs.size();i++){
        sum.push_back(o.funcs[i]);
    }

    return PolarCurve(sum);
}




#endif //UNTITLED4_POLARCURVE_H
