#include "llvm/IR/Module.h"
#include "llvm/IR/DerivedTypes.h"
#include "llvm/IR/LLVMContext.h"
#include "llvm/IR/Function.h"
#include "llvm/IR/PassManager.h"
#include "llvm/IR/Verifier.h"
#include "llvm/IR/IRPrintingPasses.h"
#include "llvm/IR/IRBuilder.h"
#include "llvm/Support/raw_ostream.h"
#include "llvm/Support/raw_os_ostream.h"

#include <fstream>

using namespace llvm;

static LLVMContext TheContext;

int main(int argc, char **argv) {
    Module *mod = new Module("all", TheContext);

    FunctionType *funcType =
            FunctionType::get(Type::getInt32Ty(TheContext), false);

    Function *func =
            Function::Create(funcType, Function::ExternalLinkage, "main", mod);


    BasicBlock *block = BasicBlock::Create(TheContext, "entry", func);

    IRBuilder<> builder(TheContext);

    builder.SetInsertPoint(block);

    auto *L = ConstantInt::get(Type::getInt32Ty(TheContext), 353);
    auto *R = ConstantInt::get(Type::getInt32Ty(TheContext), 48);

    auto result = builder.CreateAdd(L, R);

    builder.CreateRet(result);

    std::ofstream StdOutputFile("lel");

    raw_os_ostream OutputFile(StdOutputFile);

    mod->print(OutputFile, nullptr);

    delete mod;
    return 0;
}