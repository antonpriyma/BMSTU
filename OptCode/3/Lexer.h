#include <string>

#ifndef LEXER_H
#define LEXER_H

class Token {
public:
    Token(int type, int intValue) :
            type(type), intValue(intValue), stringValue("") {};
    Token(int type, std::string stringValue) :
            type(type), intValue(-1), stringValue(stringValue) {};
    Token(): type(-1), intValue(0), stringValue("") {};
    int getType() { return type; };
    std::string getStringValue() { return stringValue; };
    int getIntValue() { return intValue; };
    std::string toString() { return stringValue + std::to_string(intValue); };
private:
    int type;
    std::string stringValue;
    int intValue;
};

class Lexer{
public:
    Lexer(std::string text) : text(text + "$"), pos(0), curChar(text.at(0)) {};
    Lexer(): text(" "), pos(0), curChar('$') {};
    Token nextToken();
private:
    std::string text;
    int pos;
    char curChar;
    char nextChar();
};

#endif /* LEXER_H */

