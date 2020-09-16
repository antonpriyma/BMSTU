#include <vector>
#include <string>

#include "llvm/IR/Value.h"

#ifndef AST_H
#define AST_H

using namespace llvm;

class Expr {
public:
    virtual ~Expr() {}
    virtual std::string toString() { return "Expr()"; }
    virtual Value *codegen() = 0;
};

class NumberExpr : public Expr {
    int value;
public:
    NumberExpr(int value) : value(value) {}
    std::string toString() override { return "number(" + std::to_string(value) + ")"; }
    Value *codegen() override;
};

class Variable : public Expr {
public:
    std::string value;
    Variable(std::string value) : value(value) {}
    std::string toString() override { return "variable(" + value + ")"; }
    Value *codegen() override;
};

class BinOp : public Expr {
    char op;
    Expr *lhs, *rhs;
public:
    BinOp(char op, Expr *lhs, Expr *rhs) : op(op), lhs(lhs), rhs(rhs) {}
    std::string toString() override { return "BinOp(" +  lhs->toString() + op + rhs->toString() + ")"; }
    Value *codegen() override;
};

class IfExpr : public Expr {
    bool hasElse;
    Expr *cond, *thenBranch, *elseBranch;
public:
    IfExpr(Expr *cond, Expr *thenBranch, Expr *elseBranch) :
        cond(cond), thenBranch(thenBranch), elseBranch(elseBranch), hasElse(true) {};
    IfExpr(Expr *cond, Expr *thenBranch) : cond(cond), thenBranch(thenBranch), hasElse(false) {}
    std::string toString() override { return "if(" + cond->toString() + ") {" + thenBranch->toString() + "}"; }
    Value *codegen() override;
};

class MainFunction : public Expr {
    Expr *body, *returnExpr;
public:
    MainFunction(Expr *body, Expr *returnExpr) : body(body), returnExpr(returnExpr) {}
    
    std::string toString() override { return "MainFunction(" + body->toString() + " return:" + returnExpr->toString() + ")"; }
    Value *codegen() override;
    Module* getModule();
};


class Block : public Expr {
    std::vector<Expr*> exprs;
public:
    Block(std::vector<Expr*> exprs) : exprs(exprs) {}
    std::string toString() override;
    Value *codegen() override;
};

class Declaration : public Expr {
    Expr *variable, *value;
public:
    Declaration(Expr *variable, Expr *value) : variable(variable), value(value) {};
    std::string toString() override {return "Var(" + variable->toString() + "  " + value->toString() + ")"; };
    Value *codegen() override;
};

#endif /* AST_H */

