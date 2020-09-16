#include <string>
#include <iostream>

#include "llvm/IR/IRBuilder.h"
#include "llvm/IR/LLVMContext.h"
#include "llvm/IR/Module.h"
#include "llvm/Support/raw_ostream.h"
#include "llvm/Support/raw_os_ostream.h"

#include <fstream>

#include "AST.h"

static LLVMContext TheContext;
static IRBuilder<> Builder(TheContext);
static Module* TheModule;
static std::map<std::string, Value *> NamedValues;


static AllocaInst *CreateEntryBlockAlloca(Function *TheFunction,
                                          const std::string &VarName) {
    IRBuilder<> TmpB(&TheFunction->getEntryBlock(),
                 TheFunction->getEntryBlock().begin());
    return TmpB.CreateAlloca(Type::getInt32Ty(TheContext), 0,
                           VarName.c_str());
}


 
std::string Block::toString(){
    std::string result("block( ");
    for(int i=0; i<exprs.size(); i++){
        result += exprs.at(i)->toString() + " ";
    }
    result += ")";
    return result;
}

Value* NumberExpr::codegen(){
    return ConstantInt::get(TheContext, APInt(32, value, true));
}

Value* Variable::codegen(){
    Value *V = NamedValues[value];
    if (!V)
        std::cout << "unknown variable" << std::endl;
  return Builder.CreateLoad(V, value.c_str());
}

Value* BinOp::codegen(){
    Value *L = lhs->codegen();
    Value *R = rhs->codegen();
    
    switch(op){
        case '+':
            return Builder.CreateAdd(L, R, "addtmp");
        case '*':
            return Builder.CreateMul(L, R, "multmp");
        default:
            std::cout << "unknown operation" << std::endl;
            return 0;
    }
}

Value* Block::codegen(){
    for(auto it = exprs.begin(); it != exprs.end(); it++){
        (*it)->codegen();

    }
    return Constant::getNullValue(Type::getDoubleTy(TheContext));
}

Value* Declaration::codegen(){
    Function *TheFunction = Builder.GetInsertBlock()->getParent();
   
    Variable *v = dynamic_cast<Variable*>(variable);
    
    AllocaInst *Alloca = CreateEntryBlockAlloca(TheFunction, v->value);
    
    Value *val = value->codegen();
    
    Builder.CreateStore(val, Alloca);

    NamedValues[v->value] = Alloca;
    
    return val;
}

Value* MainFunction::codegen(){
    TheModule = new Module("all", TheContext);
    
    
    FunctionType *funcType =
		FunctionType::get(Type::getInt32Ty(TheContext), false);
    Function *func =
		Function::Create(funcType, Function::ExternalLinkage, "main", TheModule);
    BasicBlock* block = BasicBlock::Create(TheContext, "begin", func);
    
    Builder.SetInsertPoint(block);
    

    
    body->codegen();
    Value *forReturn = returnExpr->codegen();
    
    Builder.CreateRet(forReturn);
    
    std::ofstream StdOutputFile("lul");
	
    raw_os_ostream OutputFile(StdOutputFile);
    
    TheModule->print(OutputFile, nullptr);
    
    return 0;
}

Value* IfExpr::codegen(){
    Value *CondV = cond->codegen();

    CondV = Builder.CreateICmpNE(
        CondV, ConstantInt::get(TheContext, APInt(32, 0, true)), "ifcond");

    Function *TheFunction = Builder.GetInsertBlock()->getParent();

    BasicBlock *ThenBB = BasicBlock::Create(TheContext, "true", TheFunction);
    
    
    
    if(hasElse){
        BasicBlock *ElseBB = BasicBlock::Create(TheContext, "false");
        BasicBlock *MergeBB = BasicBlock::Create(TheContext, "after");
        Builder.CreateCondBr(CondV, ThenBB, ElseBB);
        
        Builder.SetInsertPoint(ThenBB);

        Value *ThenV = thenBranch->codegen();
    
        Builder.CreateBr(MergeBB);
        ThenBB = Builder.GetInsertBlock();

        TheFunction->getBasicBlockList().push_back(ElseBB);
        Builder.SetInsertPoint(ElseBB);

        Value *ElseV = elseBranch->codegen();
        
        Builder.CreateBr(MergeBB);

        ElseBB = Builder.GetInsertBlock();


        TheFunction->getBasicBlockList().push_back(MergeBB);
        Builder.SetInsertPoint(MergeBB);

        return Constant::getNullValue(Type::getDoubleTy(TheContext));
    } else {
        BasicBlock *after = BasicBlock::Create(TheContext, "after");
        Builder.CreateCondBr(CondV, ThenBB, after);
        Builder.SetInsertPoint(ThenBB);
        
        Value *ThenV = thenBranch->codegen();
        
        Builder.CreateBr(after);
        
        TheFunction->getBasicBlockList().push_back(after);
        
        Builder.SetInsertPoint(after);
        
        return Constant::getNullValue(Type::getDoubleTy(TheContext));
    }
}