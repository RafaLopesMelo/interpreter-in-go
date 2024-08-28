# 💻 RM Lang

## 📄 Introduction

This project is an implementation of an interpreter written in (Go)[https://go.dev/] for a fictional programming language called "RMLang".

## 🚀 Features

- Basic syntax
  - Variables binding
  - Arithmetic expressions 
- Common data types support
  - Integer
  - Boolean
  - String
  - Array
  - Hash map
- Operators
  - Arithmetic operators: +, -, *, /
  - Comparison operators: ==, !=, <, >, <=, >=
  - Logical operators: ! (not)
- Control structures
  - If statements: Basic conditional statements
- Functions
  - First class citizens
  - High order functions
  - Anonimous functions
- Built-in functions
  - **len**: Accepts an array or string as unique argument and returns its size or length
  - **first**: Accepts an array as unique argument and returns its first element
  - **last**: Accepts an array as unique argument and returns its last element
  - **rest**: Accepts an array as unique argument and returns its elements except the first one
  - **push**: Accepts an array as first argument and a expression as second argument, creates a copy of the array adding the element at the last position and returns it
  - **puts**: Prints the arguments to the STDOUT
- REPL
