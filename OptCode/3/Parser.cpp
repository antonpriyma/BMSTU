#define IDENTIFIER -1
#define NUMBER 1
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

#include <vector>
#include <iostream>

#include "Parser.h"
#include "AST.h"

void Parser::nextToken(){
   curToken = lexer.nextToken();
   if (curToken.getType() == 1111) {
       std::cout << "lexical error\n";
       exit(2);
   }
}

Expr* Parser::parseNumber(){
    auto result = new NumberExpr(curToken.getIntValue());
    
    nextToken();
    return result;
}

Expr* Parser::parseIdentifier(){
    auto result = new Variable(curToken.getStringValue());
    
    nextToken();
    return result;
}

Expr* Parser::parseParenExpr(){
    nextToken();
    auto expr = parseBinOp();
    
    
    if(curToken.getType() != RPAREN){
        std::cout << ") expected founded: " << curToken.toString() << std::endl;
    }
    
    nextToken();
    
    return expr;
}

Expr* Parser::parsePrimary(){
    switch(curToken.getType()){
        case IDENTIFIER:
            return parseIdentifier();
            break;
        case NUMBER:
            return parseNumber();
            break;
        case LPAREN:
            return parseParenExpr();
            break;
    }
    return 0;
}

Expr* Parser::parseBinOpRhs(Expr* lhs){
    
    if(curToken.getType() != 6 && curToken.getType() != 7){
        return lhs;
    }
    
    char op = curToken.getStringValue()[0];
    nextToken();
    
    auto rhs = parsePrimary();
    
    if(curToken.getType() == 6 || curToken.getType() == 7){
        rhs = parseBinOpRhs(rhs);
    }
    
    return new BinOp(op, lhs, rhs);
}

Expr* Parser::parseBinOp(){
    auto lhs = parsePrimary();
    
   
    
    return parseBinOpRhs(lhs);
}

Expr* Parser::parseIf(){
    nextToken();
    
    auto cond = parseParenExpr();
    
    if(curToken.getType() != LBRACE){
        std::cout << "{ expected" << std::endl;
    }
    
    auto thenBlock = parseBlock();
    
    nextToken();
    
   
    
    if(curToken.getType() == ELSE){
        nextToken();
        
        if(curToken.getType() != LBRACE){
            std::cout << "{ expected" << std::endl;
        }
        
        auto elseBlock = parseBlock();
        
        nextToken();
        
        return new IfExpr(cond, thenBlock, elseBlock);
    }
    
    return new IfExpr(cond, thenBlock);
}

Expr* Parser::parseDeclaration(){
    nextToken();
    
    auto variable = parseIdentifier();
    
    
    if(curToken.getType() != EQUAL){
        std::cout << "= expected" << std::endl;
    }
    
    nextToken();
    
    auto value = parsePrimary();
    
    return new Declaration(variable, value);
}

Expr* Parser::parseBlock(){
    nextToken();
    
    bool flag = true;
    
    std::vector<Expr*> elems;
    
    while(flag){
        switch(curToken.getType()){
            case VAR:
                elems.push_back(parseDeclaration());
                break;
            case IF:
                elems.push_back(parseIf());
                break;
            default:
                flag = false;
                break;
        }
    }
    
    return new Block(elems);
}


Expr* Parser::mainFunctionParsing(){
    if(curToken.getType() != FUNCTION){
        std::cout << "function expected"<< std::endl;
    }
    
    nextToken();
    
    if(curToken.getType() != MAIN){
        std::cout << "main expected" << std::endl;
    }
    
    nextToken();
    
    if(curToken.getType() != LPAREN){
        std::cout << "( expected" << std::endl;
    }
    
    nextToken();
    
    if(curToken.getType() != RPAREN){
        std::cout << ") expected" << std::endl;
    }
    
    nextToken();
    
    if(curToken.getType() != LBRACE){
        std::cout << "{ expected " << std::endl;
    }
    
    auto body = parseBlock();
    
    if(curToken.getType() != RETURN){
        std::cout << "return expected" << std::endl;
    }
    
    nextToken();
    
    auto returnStatment = parsePrimary();
    
    return new MainFunction(body, returnStatment);
}
