cmake_minimum_required(VERSION 3.15)
project(4_new)

set(CMAKE_CXX_STANDARD 14)


find_package(LLVM REQUIRED CONFIG)

message(STATUS "Found LLVM ${LLVM_PACKAGE_VERSION}")
message(STATUS "Using LLVMConfig.cmake in: ${LLVM_DIR}")

add_executable(4_new parser lexer irbuilder codegen compiler.cpp test)

llvm_map_components_to_libnames(llvm_libs support core irreader)

set (LIBS
        ${LLVM_LIBRARIES_PATH}/LLVMSupport.lib
        ${LLVM_LIBRARIES_PATH}/LLVMCore.lib
        ${LLVM_LIBRARIES_PATH}/LLVMMC.lib
        ${LLVM_LIBRARIES_PATH}/LLVMTarget.lib
        ${LLVM_LIBRARIES_PATH}/LLVMAnalysis.lib
        ${LLVM_LIBRARIES_PATH}/LLVMipa.lib
        ${LLVM_LIBRARIES_PATH}/LLVMTransformUtils.lib
        ${LLVM_LIBRARIES_PATH}/LLVMInstCombine.lib
        ${LLVM_LIBRARIES_PATH}/LLVMScalarOpts.lib
        ${LLVM_LIBRARIES_PATH}/LLVMCodeGen.lib
        ${LLVM_LIBRARIES_PATH}/LLVMExecutionEngine.lib
        ${LLVM_LIBRARIES_PATH}/LLVMJIT.lib
        ${LLVM_LIBRARIES_PATH}/LLVMX86Utils.lib
        ${LLVM_LIBRARIES_PATH}/LLVMX86Info.lib
        ${LLVM_LIBRARIES_PATH}/LLVMMCParser.lib
        ${LLVM_LIBRARIES_PATH}/LLVMX86AsmParser.lib
        ${LLVM_LIBRARIES_PATH}/LLVMX86AsmPrinter.lib
        ${LLVM_LIBRARIES_PATH}/LLVMAsmPrinter.lib
        ${LLVM_LIBRARIES_PATH}/LLVMSelectionDAG.lib
        ${LLVM_LIBRARIES_PATH}/LLVMX86CodeGen.lib
        ${LLVM_LIBRARIES_PATH}/LLVMX86Disassembler.lib
        ${LLVM_LIBRARIES_PATH}/LLVMInterpreter.lib
        )

# Link against LLVM libraries
target_link_libraries(4_new ${llvm_libs} ${LIBS})