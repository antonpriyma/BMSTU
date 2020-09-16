#include <string>
#include <vector>
#include "Parser.h"
#include <cstdlib>
#include <iostream>
#include <fstream>
#include <llvm/IR/Module.h>

int main(int argc, char** argv) {
    std::vector<Token> tokens;
    std::string s = "function main(){var d = 13 if(d+4){var a = 5} else {var a = 10} return a}";

    Parser(s).mainFunctionParsing()->codegen();
    
    return 0;
}

