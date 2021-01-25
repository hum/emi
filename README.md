# Emi
Interpreter for a made-up language written in Go.

### Todo:
- [x] Finish basic functions of the lexer
- [x] A REPL
- [ ] Write a parser
  - [ ] AST
- [ ] Evaluate expressions
- [ ] Add more data types and built-in functions to the interpreter (arrays, hashes, etc.)


### Syntax
```go
let five = 5;                 
let ten = 10;                 
let add = fn(x, y) {          
  x + y;                      
};                            
let result = add(five, ten);
```
