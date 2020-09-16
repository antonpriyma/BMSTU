#include "Lexer.h"
#include "AST.h"

#ifndef PARSER_H
#define PARSER_H

class Parser{
public:
    Lexer lexer;
    Token curToken;
    void nextToken();
    Expr* parseNumber();
    Expr* parseParen();
    Expr* parseIdentifier();
    Expr* parseBinOp();
    Expr* parseBinOpRhs(Expr *lhs);
    Expr* parseParenExpr();
    Expr* parsePrimary();
    Expr* parseIf();
    Expr* parseBlock();
    Expr* parseDeclaration();
    Expr* mainFunctionParsing();
    Parser(std::string text) {this->lexer = Lexer(text); nextToken(); };
};

#endif /* PARSER_H */

