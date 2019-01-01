#include <iostream>
#include <fstream>
#include <sstream>
#include <cstdlib>
#include <algorithm>
#include <vector>
#include <set>
#include <map>
#include <string>
#include <unordered_set>


using namespace std;

set<string> makeBiogr(set<string> &v,string s){
    int len =s.length();
    if (s.size()<=2)
        v.insert(s);
    for(auto it=s.begin();it!=s.end()-1;it++) {
        string t(it,it+2);
        v.emplace(t);
    }
    return v;
}

void checkIt(map<string,pair<set<string>,int>> dict,string s){
    set<string> biogr;
   biogr= makeBiogr(biogr,s);
    double max = -0.1;
    string word=s;
    int freq;

    for (auto it = dict.begin();it!=dict.end();it++){
        set<string> tmp;
        set_intersection((*it).second.first.begin(),(*it).second.first.end(),biogr.begin(),biogr.end(),inserter(tmp,tmp.begin()));
        double inter_ = tmp.size();
        if (inter_==0)
            continue;
        else tmp.clear();
        set_union((*it).second.first.begin(),(*it).second.first.end(),biogr.begin(),biogr.end(),inserter(tmp,tmp.begin()));
        double union_ = tmp.size();
        double sim=  inter_ / union_;
        if (sim>max){
            max=sim;
            word=(*it).first;
            freq = (*it).second.second;
        }else if ((sim==max && freq < (*it).second.second) || (sim==max && freq == (*it).second.second && (*it).first < word)) {
            word =(*it).first;
            freq=(*it).second.second;
        }
    }
    if (s != "")
        cout << word << endl;
}

const unordered_set<char> delimiters {
        '~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_',
        '+', '-', '=', '`', '{', '}', '[', ']', '|', '\\', ':', ';',
        '\"', '\'', ',', '.', '/', '<', '>', '?', ' ', '\t', '\n'
};

int main() {
    set<string> biograms;
    string word;
    map<string,pair<set<string>,int>> dict;
     ifstream f("/home/anton/CLionProjects/spellchecker/count_big.txt");
    string s;
     string line;
    while (getline(f,line)) {
        string a="";
        string numb="";
        vector<string> freq_string;
        bool isWord=true;
        for (auto it = line.begin();it!=line.end();it++){
            if (delimiters.find(*it)==delimiters.end() && isWord){
                a+=*it;
            } else{
                isWord= false;
                numb+=*it;
            }
        }
        int n = atoi(numb.c_str());
        set<string> buf;
        buf=makeBiogr(buf,a);
        pair<set<string>,int> elem(buf,n);
        dict.emplace(a,elem);
    }//составили словарь


    do{
        cin>>word;
        checkIt(dict,word);
    }while (getline(cin,word));
    return 0;
}
