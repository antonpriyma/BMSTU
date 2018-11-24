#include <stdio.h>
#include <string.h>

int main() {
    char s[201];
    char c;
    int k=21,i=0, word=0, n=0;
    while ((c=getchar())!='.'){
        s[i++]=c;
        if (c!=' ' && word==0){
            word=1;
            n++;
        } else if (c!=' ' && word==1){
            n++;
            continue;
        }  if (c==' ' && word==0){
            continue;
        }  if (c==' ' && word==1){
            if (n<k){
                k=n;
            }
            word=0;
            n=0;
        }
    }

    int len=strlen(s);
    for (int i=0; i<len; i++){
        if (s[i]>='a' && s[i]<='z'){
            s[i]=s[i]-k;
            if (s[i]<'a'){
                s[i]+=26;
            }
        }
        if (s[i]>='A' && s[i]<='Z'){
            s[i]=s[i]-k;
            if (s[i]<'A'){
                s[i]+=26;
            }
        }
    }
    printf("%s",s);
    return 0;
}