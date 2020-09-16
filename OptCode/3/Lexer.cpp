#define IF 3
#define INT 4
#define RETURN 5
#define PLUS 6
#define MUL 7
#define LBRACE 8
#define RBRACE 9
#define LPAREN 10
#define RPAREN 11
#define ELSE 12
#define EQUAL 13
#define VAR 14
#define FUNCTION 15
#define MAIN 16

#include <string>
#include <iostream>

#include "Lexer.h"

char Lexer::nextChar(){
    
    if(text.at(pos) != '$'){
        pos++;
        curChar = text.at(pos);
        return curChar;
    } else {
        return '$';
    }
}

Token Lexer::nextToken(){
    while(std::isspace(curChar)){
        nextChar();
    }
    
    if(std::isdigit(curChar)){
        int value(0);
        
        while(std::isdigit(curChar)){
            value = 10*value + (curChar - '0');
            nextChar();
        }
        
        return Token(1, value);
    } else if(std::isalpha(curChar)){
        std::string value("");
        
        while(std::isalnum(curChar)){
            value += curChar;
            nextChar();
        }
        
        if(value.compare("if") == 0){
            return Token(IF, value);
        }
        
        if(value.compare("int") == 0){
            return Token(INT, value);
        }
        
        if(value.compare("return") == 0){
            return Token(RETURN, value);
        }
        
        if(value.compare("else") == 0){
            return Token(ELSE, value);
        }
        
        if(value.compare("var") == 0){
            return Token(VAR, value);
        }
        
        if(value.compare("function") == 0){
            return Token(FUNCTION, value);
        }
        
        if(value.compare("main") == 0){
            return Token(MAIN, value);
        }
        
        return Token(-1, value);
    } else {
        char cur = curChar;
        nextChar();
        switch(cur){
            case '+':
                return Token(PLUS, "+");
            case '*':
                return Token(MUL, "*");
            case '{':
                return Token(LBRACE, "{");
            case '}':
                return Token(RBRACE, "}");
            case '(':
                return Token(LPAREN, "(");
            case ')':
                return Token(RPAREN, ")");
            case '=':
                return Token(EQUAL, "=");
            case '$':
                return Token(999, "EOF");
        }
        
        return Token(1111, "Char not found");
    }
}


