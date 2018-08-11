# IDEAs

sourcecode -> lexer
lexer recognizes sourcecode using predefined unicode symbols
lexer emits raw tokens
syntax parser takes these tokens and performs tokens smashing (token minimization) into the respective tokens:
  FOR IDENT KEYWORD ITERABLE BLOCK would compose into FOR with properties for start, end, keyword, body, etc
semantic parser will then take these tokens and unravel them internally to ensure that they perform what they mean
  and make sense while doing so
the semantic also stringifies the tokens

<br>

## Token:
  - Name  string -> `forin, forof, forover, etc` ...
  - Type  TokenType -> `INT, FOR, IDENT, etc` ...
  - ID    int
  - value Value
  - String string -> `stringified token: "for", "if", [name_of_ident], etc` ...
> Setters and Getters for all values above
  - SetProp()
  - GetProp()
  - Express() -> `back-produces entire string for token in Express`
  - Translate() -> `produces entire to whatever implements Translate (LLVM, C++, Rust, etc` ...)
<br><br><br>
## Value:
  - String string
  - True interface{}
  - Type ValueType -> `int, string, var, etc` ...
  - Props:<br>&nbsp;&nbsp;`PropMap {`<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`sync.Mutex`<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`Map map[string]Token`<br>
  &nbsp;&nbsp;`}`