package ast

import "github.com/minecraft-script/internal/token"

type Node interface {
	TokenLiteral() string
}

// ---------- Statements ----------

// 整个程序由多条语句组成 | The entire program consists of multiple statements
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// Statement 表示顶层语句（赋值语句等） | Represents top-level statements (assignment statements, etc.)
type Statement interface {
	Node
	statementNode()
}

// 变量赋值语句: name = (1, 2, 3) | Variable assignment statement: name = (1, 2, 3)
type AssignStatement struct {
	Token token.Token // '='
	Name  *Identifier
	Value Expression
}

func (as *AssignStatement) statementNode()       {}
func (as *AssignStatement) TokenLiteral() string { return as.Token.Literal }

// ---------- Expressions ----------

type Expression interface {
	Node
	expressionNode()
}

type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// 向量字面量: [1, 2, 3] | Vector literal: [1, 2, 3]
type VectorLiteral struct {
	Token token.Token // '['
	Values []int
}

func (vl *VectorLiteral) expressionNode()      {}
func (vl *VectorLiteral) TokenLiteral() string { return vl.Token.Literal }

// 字符串字面量: "stone" | String literal: "stone"
type StringLiteral struct {
	Token token.Token // '"'
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }

// 函数调用表达式: fill(pos1, pos2, "stone") | Function call expression: fill(pos1, pos2, "stone")
type CallExpression struct {
	Token     token.Token // 函数名 | Function name
	Function  *Identifier
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) statementNode()       {} // 添加这个方法使CallExpression也可以作为Statement使用 | Add this method to make CallExpression usable as a Statement
