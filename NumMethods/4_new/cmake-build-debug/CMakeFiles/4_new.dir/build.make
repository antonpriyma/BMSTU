# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.15

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Remove some rules from gmake that .SUFFIXES does not remove.
SUFFIXES =

.SUFFIXES: .hpux_make_needs_suffix_list


# Suppress display of executed commands.
$(VERBOSE).SILENT:


# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /Applications/CLion.app/Contents/bin/cmake/mac/bin/cmake

# The command to remove a file.
RM = /Applications/CLion.app/Contents/bin/cmake/mac/bin/cmake -E remove -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /Users/a.priyma/Programs/BMSTU/NumMethods/4_new

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /Users/a.priyma/Programs/BMSTU/NumMethods/4_new/cmake-build-debug

# Include any dependencies generated for this target.
include CMakeFiles/4_new.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/4_new.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/4_new.dir/flags.make

CMakeFiles/4_new.dir/compiler.cpp.o: CMakeFiles/4_new.dir/flags.make
CMakeFiles/4_new.dir/compiler.cpp.o: ../compiler.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/Users/a.priyma/Programs/BMSTU/NumMethods/4_new/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building CXX object CMakeFiles/4_new.dir/compiler.cpp.o"
	/Library/Developer/CommandLineTools/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/4_new.dir/compiler.cpp.o -c /Users/a.priyma/Programs/BMSTU/NumMethods/4_new/compiler.cpp

CMakeFiles/4_new.dir/compiler.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/4_new.dir/compiler.cpp.i"
	/Library/Developer/CommandLineTools/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /Users/a.priyma/Programs/BMSTU/NumMethods/4_new/compiler.cpp > CMakeFiles/4_new.dir/compiler.cpp.i

CMakeFiles/4_new.dir/compiler.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/4_new.dir/compiler.cpp.s"
	/Library/Developer/CommandLineTools/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /Users/a.priyma/Programs/BMSTU/NumMethods/4_new/compiler.cpp -o CMakeFiles/4_new.dir/compiler.cpp.s

# Object files for target 4_new
4_new_OBJECTS = \
"CMakeFiles/4_new.dir/compiler.cpp.o"

# External object files for target 4_new
4_new_EXTERNAL_OBJECTS =

4_new: CMakeFiles/4_new.dir/compiler.cpp.o
4_new: CMakeFiles/4_new.dir/build.make
4_new: /usr/local/opt/llvm/lib/libLLVMSupport.a
4_new: /usr/local/opt/llvm/lib/libLLVMCore.a
4_new: /usr/local/opt/llvm/lib/libLLVMIRReader.a
4_new: /LLVMSupport.lib
4_new: /LLVMCore.lib
4_new: /LLVMMC.lib
4_new: /LLVMTarget.lib
4_new: /LLVMAnalysis.lib
4_new: /LLVMipa.lib
4_new: /LLVMTransformUtils.lib
4_new: /LLVMInstCombine.lib
4_new: /LLVMScalarOpts.lib
4_new: /LLVMCodeGen.lib
4_new: /LLVMExecutionEngine.lib
4_new: /LLVMJIT.lib
4_new: /LLVMX86Utils.lib
4_new: /LLVMX86Info.lib
4_new: /LLVMMCParser.lib
4_new: /LLVMX86AsmParser.lib
4_new: /LLVMX86AsmPrinter.lib
4_new: /LLVMAsmPrinter.lib
4_new: /LLVMSelectionDAG.lib
4_new: /LLVMX86CodeGen.lib
4_new: /LLVMX86Disassembler.lib
4_new: /LLVMInterpreter.lib
4_new: /usr/local/opt/llvm/lib/libLLVMAsmParser.a
4_new: /usr/local/opt/llvm/lib/libLLVMBitReader.a
4_new: /usr/local/opt/llvm/lib/libLLVMCore.a
4_new: /usr/local/opt/llvm/lib/libLLVMBinaryFormat.a
4_new: /usr/local/opt/llvm/lib/libLLVMRemarks.a
4_new: /usr/local/opt/llvm/lib/libLLVMBitstreamReader.a
4_new: /usr/local/opt/llvm/lib/libLLVMSupport.a
4_new: /usr/local/opt/llvm/lib/libLLVMDemangle.a
4_new: CMakeFiles/4_new.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/Users/a.priyma/Programs/BMSTU/NumMethods/4_new/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Linking CXX executable 4_new"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/4_new.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/4_new.dir/build: 4_new

.PHONY : CMakeFiles/4_new.dir/build

CMakeFiles/4_new.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/4_new.dir/cmake_clean.cmake
.PHONY : CMakeFiles/4_new.dir/clean

CMakeFiles/4_new.dir/depend:
	cd /Users/a.priyma/Programs/BMSTU/NumMethods/4_new/cmake-build-debug && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /Users/a.priyma/Programs/BMSTU/NumMethods/4_new /Users/a.priyma/Programs/BMSTU/NumMethods/4_new /Users/a.priyma/Programs/BMSTU/NumMethods/4_new/cmake-build-debug /Users/a.priyma/Programs/BMSTU/NumMethods/4_new/cmake-build-debug /Users/a.priyma/Programs/BMSTU/NumMethods/4_new/cmake-build-debug/CMakeFiles/4_new.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/4_new.dir/depend

