# Conan Compiler

## Overview:
Conan_Interpreter is a general-purpose compiler for conan language, similar to C-like syntax and rich set of features.

- Design: Lexis -> Parsing -> Abstract Syntax Tree -> Evaluation/Compilation -> VM Execution

## ✨Technology
- Go

## 🚀Features
- Data Types: Integer, Booleans, Strings, Arrays, and Hashes
- Operators: Prefix( !,- ), infix (+,-,*,/,>,<,==,!= ) and Index Operators
- Control Flow: Conditionals(if/else) and return statements.
- Bindings: Global and Local Binding support.
- Functions: First-class functions, higher order functions, closures, and built-in functions(len, puts, first, rest, push, etc)

## Architecture
1. Lexer/Tokens: The Lexer reads source code and produces tokens (token.INT, token.IDENT).
2. Parsers/AST: The Parser converts the tokens into an Abstract Syntax Tree (AST).
3. Compiler/Bytecode: The Compiler traverses the AST and translates nodes into a sequence of bytecode instructions.
4. VM/Execution: The Virtual Machine executes the bytecode via a fetch-decode-execute cycle.

## Core Compiler Data Structures
- Constant Pool: Stores objects that are determined at compile time, such as integer, string and functional literals.
- Symbol Table: Used to determine the scope of identifiers and assign them a unique index. Eg: Global, Local, Builtin, Free and Functions.
- Compilation Scopes: A stack of CompilationScope objects is used to manage instructions, allowing instructions compiled inside a function body to be isolated from instructions in the main program.
- Main Stack: A LIFO structure(size: 2048) for computation, storing object values, intermediate results, function arguments and local variables.

## 🚦Running the Project
1. Clone the repository
2. Install the `Go` lang.
3. Run `go run main.go`.
4. Write go code and execute.

## Performance
The compiler and virtual machine architecture significantly improves performance over the original tree-walking interpreter. Based on a recursive Fibonacci computation benchmark, the VM implementation is measured to be about 3.3 times faster compared to tree-walking interpreter.

## Preview
![ConanInterpreter](docs/Preview/Conan_Interpreter.mov)