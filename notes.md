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
