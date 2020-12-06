# Emi
Interpreter for a made-up language written in Go.

This is a sort-of pet project of mine. Both to understand Go more and to understand interpreters.

### Todo:
- [x] Finish basic functions of the lexer
- [x] A REPL
- [ ] Write a parser
  - [ ] AST
- [ ] Evaluate expressions
- [ ] Add more data types and built-in functions to the interpreter (arrays, hashes, etc.)


### Syntax of the language
```go
let five = 5;                 // LET, IDENT, ASSIGN, INT, SEMICOLON
let ten = 10;                 // LET, IDENT, ASSIGN, INT, SEMICOLON
let add = fn(x, y) {          // LET, IDENT, ASSIGN, FUNCTION, LPAREN, IDENT, COMMA, IDENT, RPAREN, LBRACE
  x + y;                      // IDENT, PLUS, IDENT, SEMICOLON
};                            // RBRACE, SEMICOLON
let result = add(five, ten);  // LET, IDENT, ASSIGN, IDENT, LPAREN, IDENT, COMMA, IDENT, RPAREN, SEMICOLON, EOF
```
