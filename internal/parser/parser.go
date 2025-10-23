package parser

import (
	"strconv"

	"github.com/minecraft-script/internal/ast"
	"github.com/minecraft-script/internal/lexer"
	"github.com/minecraft-script/internal/token"
)

type Parser struct {
	l       *lexer.Lexer
	curTok  token.Token
	peekTok token.Token
	errors  []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curTok = p.peekTok
	p.peekTok = p.l.NextToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	for p.curTok.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

// ---------------- Statements ----------------

func (p *Parser) parseStatement() ast.Statement {
	if p.peekTok.Type == token.ASSIGN {
		return p.parseAssignStatement()
	}
	if p.curTok.Type == token.FILL || p.curTok.Type == token.SETBLOCK {
		return p.parseCallExpression()
	}
	return nil
}

func (p *Parser) parseAssignStatement() *ast.AssignStatement {
	stmt := &ast.AssignStatement{
		Token: p.curTok,
		Name: &ast.Identifier{
			Token: p.curTok,
			Value: p.curTok.Literal,
		},
	}

	p.nextToken() // 跳过标识符 | Skip identifier
	p.nextToken() // 跳过 = | Skip =

	stmt.Value = p.parseExpression()
	return stmt
}

// ---------------- Expressions ----------------

func (p *Parser) parseExpression() ast.Expression {
	switch p.curTok.Type {
	case token.IDENT:
		return &ast.Identifier{
			Token: p.curTok,
			Value: p.curTok.Literal,
		}
	case token.LBRACKET:
		return p.parseVectorLiteral()
	case token.STRING:
		return &ast.StringLiteral{
			Token: p.curTok,
			Value: p.curTok.Literal,
		}
	}
	return nil
}

func (p *Parser) parseVectorLiteral() *ast.VectorLiteral {
	vec := &ast.VectorLiteral{
		Token: p.curTok,
	}

	p.nextToken() // 跳过 [ | Skip [
	
	// 解析第一个数字 | Parse the first number
	if p.curTok.Type != token.NUMBER {
		p.errors = append(p.errors, "expected number in vector")
		return nil
	}
	
	num, _ := strconv.Atoi(p.curTok.Literal)
	vec.Values = append(vec.Values, num)
	
	p.nextToken() // 跳过数字 | Skip number
	
	// 解析剩余的数字 | Parse the remaining numbers
	for p.curTok.Type == token.COMMA {
		p.nextToken() // 跳过 , | Skip ,
		
		if p.curTok.Type != token.NUMBER {
			p.errors = append(p.errors, "expected number after comma in vector")
			return nil
		}
		
		num, _ := strconv.Atoi(p.curTok.Literal)
		vec.Values = append(vec.Values, num)
		
		p.nextToken() // 跳过数字 | Skip number
	}
	
	if p.curTok.Type != token.RBRACKET {
		p.errors = append(p.errors, "expected ']' at end of vector")
		return nil
	}
	
	return vec
}

func (p *Parser) parseCallExpression() *ast.CallExpression {
	expr := &ast.CallExpression{
		Token: p.curTok,
		Function: &ast.Identifier{
			Token: p.curTok,
			Value: p.curTok.Literal,
		},
	}
	
	p.nextToken() // 跳过函数名 | Skip function name
	
	if p.curTok.Type != token.LPAREN {
		p.errors = append(p.errors, "expected '(' after function name")
		return nil
	}
	
	p.nextToken() // 跳过 ( | Skip (
	
	// 解析参数 | Parse arguments
	for p.curTok.Type != token.RPAREN {
		arg := p.parseExpression()
		if arg != nil {
			expr.Arguments = append(expr.Arguments, arg)
		}
		
		p.nextToken() // 跳过参数 | Skip argument
		
		if p.curTok.Type == token.COMMA {
			p.nextToken() // 跳过 , | Skip ,
		} else if p.curTok.Type != token.RPAREN {
			p.errors = append(p.errors, "expected ',' or ')' after argument")
			return nil
		}
	}
	
	return expr
}
