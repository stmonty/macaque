package token

type TokenType string

type Token struct {
    Type    TokenType
    Literal string

}

const (

    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    //Identifiers and Literals
    IDENT = "IDENT"
    INT   = "INT"

    // Operators
    ASSIGN = "="
    PLUS   = "+"


    //Delimiters
    SEMICOLON = ";"
    COMMA     = ","
    LPAREN    = "("
    RPAREN    = ")"
    LBRACE    = "{"
    RBRACE    = "}"

    //Keywords
    FUNCTION = "FUNC"
    LET      = "LET"

)

var keywords = map[string]TokenType{
    "fn": FUNCTION,
    "let": LET,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}
