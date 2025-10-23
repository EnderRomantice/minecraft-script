package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// 所有token类型 | All token types
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT  = "IDENT"
	NUMBER = "NUMBER"
	STRING = "STRING"

	ASSIGN = "="
	COMMA  = ","
	LPAREN = "("
	RPAREN = ")"
	LBRACKET = "["
	RBRACKET = "]"

	FILL     = "FILL"
	SETBLOCK = "SETBLOCK"
)

// 关键字映射表 | Keyword mapping table
var keywords = map[string]TokenType{
	"fill":     FILL,
	"setblock": SETBLOCK,
}

// LookupIdent 检查标识符是否是关键字 | Checks if an identifier is a keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
