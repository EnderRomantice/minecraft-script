package lexer

import "github.com/minecraft-script/internal/token"

type Lexer struct {
	input        string // 输入源代码 | Input source code
	position     int    // 当前字符位置 | Current character position
	readPosition int    // 下一个字符位置 | Next character position
	ch           byte   // 当前字符 | Current character
}

// New 返回一个新的 Lexer 实例，初始化读取字符 | Returns a new Lexer instance, initializes character reading
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar 读取下一个字符并更新 position 和 ch | Reads the next character and updates position and ch
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // 表示 EOF | Represents EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken 返回下一个 token | Returns the next token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = token.Token{Type: token.ASSIGN, Literal: string(l.ch)}
	case ',':
		tok = token.Token{Type: token.COMMA, Literal: string(l.ch)}
	case '(':
		tok = token.Token{Type: token.LPAREN, Literal: string(l.ch)}
	case ')':
		tok = token.Token{Type: token.RPAREN, Literal: string(l.ch)}
	case '[':
		tok = token.Token{Type: token.LBRACKET, Literal: string(l.ch)}
	case ']':
		tok = token.Token{Type: token.RBRACKET, Literal: string(l.ch)}
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
		return tok // 直接返回，因为已经在readString中读取了下一个字符 | Return directly, as the next character has already been read in readString
	case '#':
		l.skipComment()
		return l.NextToken() // 跳过注释后继续读取下一个token | Continue reading the next token after skipping comments
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.NUMBER
			return tok
		} else {
			tok = token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}
		}
	}

	l.readChar()
	return tok
}

// skipWhitespace 跳过空白字符 | Skips whitespace characters
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// skipComment 跳过注释 | Skips comments
func (l *Lexer) skipComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
}

// readIdentifier 读取标识符 | Reads an identifier
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber 读取数字 | Reads a number
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readString 读取字符串 | Reads a string
func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	str := l.input[position:l.position]
	l.readChar() // 跳过结束的引号 | Skip the closing quote
	return str
}

// isLetter 判断是否是字母 | Determines if a character is a letter
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit 判断是否是数字 | Determines if a character is a digit
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
