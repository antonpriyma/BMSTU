//
// Created by anton on 25.05.18.
//

//#ifndef LAB12_NUMBER_H
//#define LAB12_NUMBER_H
/*
#include <string>
#include <iterator>
#include <stack>
#include <vector>

using namespace std;

class PrefixIterator;
class StringSeq;

class Character{


friend ostream& operator<< (ostream& os, Character& n) {
    os << n.value;
    return os;
}

private:
    char value;
    StringSeq* father;//Ссылка на строку
    int numb;//Позиция в строке
    int neibour ;//Соответствующая скобка
public:
    Character(char p,int k);

    Character& operator= (char d);

    void setFather(StringSeq* x);

    void setValue(char c);

    void setNeibour(int i);

    int getNumb();

};


class StringSeq {
private:
    friend class PrefixIterator;
    friend class Character;
    vector<Character> x;
public:


    PrefixIterator begin();
    PrefixIterator end();

    StringSeq(std::string s);
    void findNeibours(std::string s);
};


class PrefixIterator:
        public std::iterator<
                std::forward_iterator_tag,
                        Character
        >
{
private:
    Character *c;
    int count;
    int length;
    bool is_default;

public:
    PrefixIterator(): is_default(true) {}
    PrefixIterator(StringSeq &n): is_default(false) {c=&n.x[0];count=0;length=n.x.size();}
    PrefixIterator(Character *n): is_default(false) {c=n;}

    bool operator == (const PrefixIterator &other) const;
    bool operator != (const PrefixIterator &other) const;
    string & operator ();
    PrefixIterator& operator++ ();
    PrefixIterator operator++ (int);
    bool is_end() const { return count>=length;}
};
#endif //LAB12_NUMBER_H
 */
