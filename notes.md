# Lexical Analysis

The first transformation when interpreting a source code is called "lexical analysis"", or "lexing". It's done by a "lexer" (also called "tokenizer" or "scanner" - some use word or the other to denote subtle differences in behaviour).

Tokens are small, easliy categorizable data structures that are then fed to the parser, which does the second transformation, transforming the tokens into an "Abstract Syntax Tree".

The lexing result looks like this:

Source code:
```
let x = 5 + 5;
```

Tokens:
```
[
LET,
IDENTIFIER(x),
EQUAL_SIGN,
INTEGER(5),
PLUS_SIGN,
INTEGER(5),
SEMICOLON,
]
```

Note that whitespaces don't show up as tokens. In our case it's okay, but there are interpreters that do consider whitespaces as tokens and may affect the final result.

A production-ready lexer should also attach the line number, column-number and filename to a token. It's useful to print better error messages, indicating exactly where the error is.

The words that look like identifiers but aren't since they're part of the language are called keywords.

Is not the lexer's job to tell if some piece of code makes sense, works or contains errors. Its only job is to transform the source code into a sequence of tokens.

The difficulty of parsing different languages often comes down to how far you have to peek ahead or (or look backwards) in the source code to make sense of it.

# Parser

A parser turns its input into a data structure that represents the input, often some kind of Abstract Syntax Tree (AST).
JSON.parse from JavaScript is a parser, and conceptually is identical to programming languages parsers.

The "Abstract" from "Abstract Syntax Tree" is based on the fact that certain details visibile in the source code are omitted in the AST, such as semicolons, newlines, whitespaces, comments, braces, brackets and parentheses: these characters only guides the parser when constructing the tree.
There is no one, universal AST format that's used by every parser. Their implementations are all pretty similar, but they differ in details.

The process of parsing is also called "syntactic analysis".

Parsing is one of the most well-understood problems in computer science and really smart people have already invested a lot of time into the problemas of parsing.
The results of their work are CFG, BNF, EBNF, parser generators and advanced parsing techniques.

There are two main strategies when parsing a programming language: top-down and bottom-up.
A lot of slightly different forms of each strategy exist. For example, "recursive descent parsing", "early parsing" or "predictive parsing" are all variations of top-down parsing.
The difference is that the top-down starts with constructing root node of the AST and then descends, while bottom-up parsing starts with the leaves and constructs the root.

The parser of this project is a recursive descent parser. And in particular, it's "top down operator precedence" parser, sometimes called "Pratt parser", after its inventor, Vaughan Pratt. It's often recommended to begginers because it closely mirrors the way we think about AST's.

The difference between expressions and statements is that expressions produce a value, while statements don't. For example, in the expression `5 + 5`, the value is `10`, while in the statement `let x = 5`, the value is `undefined`. What is an expression or statement depends of the programming language being used.
