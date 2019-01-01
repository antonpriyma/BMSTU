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

vector<string> makeBiogr(string str){
    vector<string> bi;
    char a[3] = { 0 };
    int size = str.size();
    if (size <= 2) {
        bi.push_back(str);
    }
    else
        for(int i = 0; i < size - 1; i++) {
            a[0] = str[i];
            a[1] = str[i+1];
            bi.push_back(a);
        }
    //for(string s : bi)
    //    cout << s << " ";
    //cout << endl;
    sort(bi.begin(), bi.end());
    return bi;
}

vector<string> checkIt(map<string,pair<vector<string>,int>> dict,string s,vector<string> &result){
    vector<string> biogr;
    biogr= makeBiogr(s);
    double max = -0.1;
    string word=s;
    int freq=0;

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
        auto help = it->second.second;
        if (sim > max || (sim == max&& it->second.second > freq) || (sim == max && help == freq && it->first < word)) {
            max = sim;
            freq = help;
            word = it->first;

        }
    }
    if (s != "")
        // result.push_back(word);
        cout<<word<<endl;
    return result;
}

const unordered_set<char> delimiters {
        '~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_',
        '+', '-', '=', '`', '{', '}', '[', ']', '|', '\\', ':', ';',
        '\"', '\'', ',', '.', '/', '<', '>', '?', ' ', '\t', '\n'
};

int main() {
    vector<string> result;
    set<string> biograms;
    string word="";
    map<string,pair<vector<string>,int>> dict;
    ifstream f("count_big.txt");
    while(f) {
        string str;
        int x;
        f >> str >> x;
        dict.emplace(str, make_pair(makeBiogr(str), x));
    }

    vector<string> in;
    do{
        cin>>word;
        if (word==""){
            break;
        }
        if (cin)
            result=checkIt(dict,word,result);
    }while (cin);
    f.close();
    //for (string s:result) cout << s << endl;

    return 0;
}
