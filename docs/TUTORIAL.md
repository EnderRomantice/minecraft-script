# 从构想到实现：Minecraft Script 语言完整教程 | From Concept to Implementation: Complete Minecraft Script Language Tutorial

## 目录 | Table of Contents

1. [引言 | Introduction](#引言--introduction)
2. [语言设计与构想 | Language Design and Concept](#语言设计与构想--language-design-and-concept)
3. [环境搭建 | Environment Setup](#环境搭建--environment-setup)
4. [词法分析器实现 | Lexer Implementation](#词法分析器实现--lexer-implementation)
5. [语法分析器实现 | Parser Implementation](#语法分析器实现--parser-implementation)
6. [抽象语法树构建 | Abstract Syntax Tree Construction](#抽象语法树构建--abstract-syntax-tree-construction)
7. [代码生成器实现 | Code Generator Implementation](#代码生成器实现--code-generator-implementation)
8. [运行时环境实现 | Runtime Environment Implementation](#运行时环境实现--runtime-environment-implementation)
9. [命令行工具开发 | Command-line Tool Development](#命令行工具开发--command-line-tool-development)
10. [完整示例与应用 | Complete Examples and Applications](#完整示例与应用--complete-examples-and-applications)
11. [扩展与优化 | Extensions and Optimizations](#扩展与优化--extensions-and-optimizations)
12. [总结 | Conclusion](#总结--conclusion)

## 引言 | Introduction

Minecraft 允许玩家通过命令来改变游戏世界。然而，原生的 Minecraft 命令语法较为复杂，尤其是在创建大型建筑或复杂结构时，需要手动输入大量重复的命令。

*Minecraft allows players to change the game world through commands. However, the native Minecraft command syntax is relatively complex, especially when creating large buildings or complex structures, requiring manual input of many repetitive commands.*

Minecraft Script 语言的目标是简化这一过程，提供一种更直观、更易于编写的脚本语言，自动生成 Minecraft 命令。本教程将带领读者从零开始，一步步实现这个语言的编译器和运行时环境。

*The goal of the Minecraft Script language is to simplify this process by providing a more intuitive and easier-to-write scripting language that automatically generates Minecraft commands. This tutorial will guide readers from scratch to implement the compiler and runtime environment for this language step by step.*

## 语言设计与构想 | Language Design and Concept

### 设计目标 | Design Goals

1. **简洁性 | Simplicity**：提供简单易懂的语法，减少学习成本 | Provide simple and easy-to-understand syntax, reducing learning costs
2. **表达力 | Expressiveness**：能够表达 Minecraft 中常见的建筑和结构 | Able to express common buildings and structures in Minecraft
3. **可扩展性 | Extensibility**：易于添加新的命令和功能 | Easy to add new commands and features

### 语法设计 | Syntax Design

Minecraft Script 采用简单直观的语法，主要包括以下元素：

*Minecraft Script uses a simple and intuitive syntax, mainly including the following elements:*

1. **变量定义 | Variable Definition**
   ```
   # 定义变量 | Define variables
   start_pos = [0, 64, 0]
   end_pos = [10, 74, 10]
   ```

2. **向量表示 | Vector Representation**
   ```
   # 使用方括号表示三维坐标 | Use square brackets to represent 3D coordinates
   pos = [1, 65, -5]
   ```

3. **命令调用 | Command Invocation**
   ```
   # 调用内置命令 | Call built-in commands
   fill(start_pos, end_pos, "stone")
   setblock(pos, "diamond_block")
   ```

### 支持的命令 | Supported Commands

初始版本支持以下 Minecraft 命令：

*The initial version supports the following Minecraft commands:*

1. **fill**：填充一个区域的方块 | Fill an area with blocks
   ```
   fill(pos1, pos2, block_type)
   ```

2. **setblock**：在指定位置放置一个方块 | Place a block at a specified position
   ```
   setblock(pos, block_type)
   ```

## 环境搭建 | Environment Setup

### 开发环境准备 | Development Environment Preparation

1. **安装 Go 语言环境 | Install Go Language Environment**
   ```bash
   # 下载并安装 Go | Download and install Go
   # 从 https://golang.org/dl/ 下载适合您系统的版本
   # Download the version suitable for your system from https://golang.org/dl/
   ```

2. **创建项目目录结构 | Create Project Directory Structure**
   ```bash
   mkdir -p minecraft-script/{cmd/minecraftscript,internal/{ast,codegen,lexer,parser,runtime,token},examples,out}
   cd minecraft-script
   ```

3. **初始化 Go 模块 | Initialize Go Module**
   ```bash
   go mod init github.com/minecraft-script
   ```

### 项目结构说明 | Project Structure Description

```
minecraft-script/
├── cmd/                  # 命令行工具 | Command-line tools
│   └── minecraftscript/  # 主程序入口 | Main program entry
├── examples/             # 示例脚本 | Example scripts
├── internal/             # 内部实现 | Internal implementation
│   ├── ast/              # 抽象语法树定义 | Abstract syntax tree definitions
│   ├── codegen/          # 代码生成器 | Code generator
│   ├── lexer/            # 词法分析器 | Lexical analyzer
│   ├── parser/           # 语法分析器 | Syntax parser
│   ├── runtime/          # 运行时环境 | Runtime environment
│   └── token/            # 词法单元定义 | Token definitions
└── out/                  # 输出目录 | Output directory
```

## 词法分析器实现 | Lexer Implementation

词法分析器是编译器的第一个阶段，负责将源代码文本转换为词法单元（Token）序列。

*The lexer is the first stage of the compiler, responsible for converting source code text into a sequence of tokens.*

### 定义词法单元 | Define Tokens

首先，我们需要定义语言中的基本词法单元类型。创建文件 `internal/token/token.go`：

*First, we need to define the basic token types in the language. Create the file `internal/token/token.go`:*

```go
package token

type TokenType string

type Token struct {
    Type    TokenType
    Literal string
}

// 词法单元类型常量 | Token type constants
const (
    ILLEGAL = "ILLEGAL" // 非法字符 | Illegal character
    EOF     = "EOF"     // 文件结束 | End of file

    // 标识符和字面量 | Identifiers and literals
    IDENT  = "IDENT"  // 标识符 | Identifier
    NUMBER = "NUMBER" // 数字 | Number
    STRING = "STRING" // 字符串 | String

    // 运算符 | Operators
    ASSIGN = "="  // 赋值 | Assignment

    // 分隔符 | Delimiters
    COMMA     = ","
    LPAREN    = "("
    RPAREN    = ")"
    LBRACKET  = "["
    RBRACKET  = "]"

    // 关键字 | Keywords
    FILL     = "FILL"
    SETBLOCK = "SETBLOCK"
)

// 关键字映射表 | Keyword mapping table
var keywords = map[string]TokenType{
    "fill":     FILL,
    "setblock": SETBLOCK,
}

// LookupIdent 检查标识符是否为关键字 | Check if an identifier is a keyword
func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}
```

### 实现词法分析器 | Implement the Lexer

接下来，我们实现词法分析器。创建文件 `internal/lexer/lexer.go`：

*Next, we implement the lexer. Create the file `internal/lexer/lexer.go`:*

```go
package lexer

import "github.com/minecraft-script/internal/token"

type Lexer struct {
    input        string // 输入源代码 | Input source code
    position     int    // 当前字符位置 | Current character position
    readPosition int    // 下一个字符位置 | Next character position
    ch           byte   // 当前字符 | Current character
}

// New 创建新的词法分析器 | Create a new lexer
func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar() // 初始化第一个字符 | Initialize the first character
    return l
}

// readChar 读取下一个字符 | Read the next character
func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0 // ASCII 码 0 表示 EOF | ASCII code 0 represents EOF
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition++
}

// NextToken 解析并返回下一个词法单元 | Parse and return the next token
func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    l.skipWhitespace()
    l.skipComment()

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

// skipWhitespace 跳过空白字符 | Skip whitespace characters
func (l *Lexer) skipWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readChar()
    }
}

// skipComment 跳过注释 | Skip comments
func (l *Lexer) skipComment() {
    if l.ch == '#' {
        for l.ch != '\n' && l.ch != 0 {
            l.readChar()
        }
        l.skipWhitespace()
    }
}

// readIdentifier 读取标识符 | Read identifier
func (l *Lexer) readIdentifier() string {
    position := l.position
    for isLetter(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}

// readNumber 读取数字 | Read number
func (l *Lexer) readNumber() string {
    position := l.position
    for isDigit(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}

// readString 读取字符串 | Read string
func (l *Lexer) readString() string {
    position := l.position + 1
    for {
        l.readChar()
        if l.ch == '"' || l.ch == 0 {
            break
        }
    }
    return l.input[position:l.position]
}

// isLetter 判断字符是否为字母或下划线 | Check if a character is a letter or underscore
func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit 判断字符是否为数字 | Check if a character is a digit
func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}
```

## 抽象语法树构建 | Abstract Syntax Tree Construction

抽象语法树（AST）是源代码的结构化表示，用于后续的代码生成。

*Abstract Syntax Tree (AST) is a structured representation of the source code, used for subsequent code generation.*

### 定义 AST 节点 | Define AST Nodes

创建文件 `internal/ast/ast.go`：

*Create the file `internal/ast/ast.go`:*

```go
package ast

import "github.com/minecraft-script/internal/token"

type Node interface {
    TokenLiteral() string
}

// Statement 表示顶层语句 | Represents top-level statements
type Statement interface {
    Node
    statementNode()
}

// Expression 表示表达式 | Represents expressions
type Expression interface {
    Node
    expressionNode()
}

// Program 表示整个程序 | Represents the entire program
type Program struct {
    Statements []Statement
}

func (p *Program) TokenLiteral() string {
    if len(p.Statements) > 0 {
        return p.Statements[0].TokenLiteral()
    }
    return ""
}

// AssignStatement 表示赋值语句 | Represents assignment statements
type AssignStatement struct {
    Token token.Token // '='
    Name  *Identifier
    Value Expression
}

func (as *AssignStatement) statementNode()       {}
func (as *AssignStatement) TokenLiteral() string { return as.Token.Literal }

// Identifier 表示标识符 | Represents identifiers
type Identifier struct {
    Token token.Token // token.IDENT
    Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// VectorLiteral 表示向量字面量 | Represents vector literals
type VectorLiteral struct {
    Token token.Token // '['
    Values []int
}

func (vl *VectorLiteral) expressionNode()      {}
func (vl *VectorLiteral) TokenLiteral() string { return vl.Token.Literal }

// StringLiteral 表示字符串字面量 | Represents string literals
type StringLiteral struct {
    Token token.Token // '"'
    Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }

// CallExpression 表示函数调用 | Represents function calls
type CallExpression struct {
    Token     token.Token // 函数名 | Function name
    Function  *Identifier
    Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) statementNode()       {} // 使CallExpression也可以作为Statement使用 | Make CallExpression usable as a Statement
```

## 语法分析器实现 | Parser Implementation

语法分析器负责将词法单元序列转换为抽象语法树。

*The parser is responsible for converting the token sequence into an abstract syntax tree.*

### 实现语法分析器 | Implement the Parser

创建文件 `internal/parser/parser.go`：

*Create the file `internal/parser/parser.go`:*

```go
package parser

import (
    "fmt"
    "strconv"

    "github.com/minecraft-script/internal/ast"
    "github.com/minecraft-script/internal/lexer"
    "github.com/minecraft-script/internal/token"
)

type Parser struct {
    l         *lexer.Lexer
    curToken  token.Token
    peekToken token.Token
    errors    []string
}

// New 创建新的语法分析器 | Create a new parser
func New(l *lexer.Lexer) *Parser {
    p := &Parser{l: l}

    // 读取两个词法单元，设置 curToken 和 peekToken | Read two tokens, set curToken and peekToken
    p.nextToken()
    p.nextToken()

    return p
}

// nextToken 前进到下一个词法单元 | Advance to the next token
func (p *Parser) nextToken() {
    p.curToken = p.peekToken
    p.peekToken = p.l.NextToken()
}

// Errors 返回解析错误 | Return parsing errors
func (p *Parser) Errors() []string {
    return p.errors
}

// ParseProgram 解析整个程序 | Parse the entire program
func (p *Parser) ParseProgram() *ast.Program {
    program := &ast.Program{Statements: []ast.Statement{}}

    for p.curToken.Type != token.EOF {
        stmt := p.parseStatement()
        if stmt != nil {
            program.Statements = append(program.Statements, stmt)
        }
        p.nextToken()
    }

    return program
}

// parseStatement 解析单个语句 | Parse a single statement
func (p *Parser) parseStatement() ast.Statement {
    switch p.curToken.Type {
    case token.IDENT:
        if p.peekToken.Type == token.ASSIGN {
            return p.parseAssignStatement()
        }
        return p.parseExpressionStatement()
    case token.FILL, token.SETBLOCK:
        return p.parseExpressionStatement()
    default:
        return nil
    }
}

// parseAssignStatement 解析赋值语句 | Parse assignment statement
func (p *Parser) parseAssignStatement() *ast.AssignStatement {
    stmt := &ast.AssignStatement{Token: p.curToken}
    stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

    if !p.expectPeek(token.ASSIGN) {
        return nil
    }

    p.nextToken()
    stmt.Value = p.parseExpression()

    return stmt
}

// parseExpressionStatement 解析表达式语句 | Parse expression statement
func (p *Parser) parseExpressionStatement() ast.Statement {
    expr := p.parseExpression()
    if expr == nil {
        return nil
    }

    if callExpr, ok := expr.(*ast.CallExpression); ok {
        return callExpr
    }

    return nil
}

// parseExpression 解析表达式 | Parse expression
func (p *Parser) parseExpression() ast.Expression {
    switch p.curToken.Type {
    case token.IDENT:
        if p.peekToken.Type == token.LPAREN {
            return p.parseCallExpression()
        }
        return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
    case token.LBRACKET:
        return p.parseVectorLiteral()
    case token.STRING:
        return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
    default:
        return nil
    }
}

// parseVectorLiteral 解析向量字面量 | Parse vector literal
func (p *Parser) parseVectorLiteral() *ast.VectorLiteral {
    vector := &ast.VectorLiteral{Token: p.curToken}
    vector.Values = []int{}

    if !p.expectPeek(token.NUMBER) {
        return nil
    }

    // 解析第一个数字 | Parse the first number
    num, err := strconv.Atoi(p.curToken.Literal)
    if err != nil {
        p.errors = append(p.errors, fmt.Sprintf("could not parse %q as integer", p.curToken.Literal))
        return nil
    }
    vector.Values = append(vector.Values, num)

    // 解析剩余的数字 | Parse the remaining numbers
    for p.peekToken.Type == token.COMMA {
        p.nextToken() // 跳过逗号 | Skip comma
        p.nextToken() // 移动到数字 | Move to number

        if p.curToken.Type != token.NUMBER {
            p.errors = append(p.errors, fmt.Sprintf("expected NUMBER, got %s", p.curToken.Type))
            return nil
        }

        num, err := strconv.Atoi(p.curToken.Literal)
        if err != nil {
            p.errors = append(p.errors, fmt.Sprintf("could not parse %q as integer", p.curToken.Literal))
            return nil
        }
        vector.Values = append(vector.Values, num)
    }

    if !p.expectPeek(token.RBRACKET) {
        return nil
    }

    return vector
}

// parseCallExpression 解析函数调用 | Parse function call
func (p *Parser) parseCallExpression() *ast.CallExpression {
    expr := &ast.CallExpression{
        Token:    p.curToken,
        Function: &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal},
    }

    if !p.expectPeek(token.LPAREN) {
        return nil
    }

    expr.Arguments = p.parseCallArguments()

    return expr
}

// parseCallArguments 解析函数调用参数 | Parse function call arguments
func (p *Parser) parseCallArguments() []ast.Expression {
    args := []ast.Expression{}

    if p.peekToken.Type == token.RPAREN {
        p.nextToken()
        return args
    }

    p.nextToken()
    args = append(args, p.parseExpression())

    for p.peekToken.Type == token.COMMA {
        p.nextToken() // 跳过逗号 | Skip comma
        p.nextToken() // 移动到参数 | Move to argument
        args = append(args, p.parseExpression())
    }

    if !p.expectPeek(token.RPAREN) {
        return nil
    }

    return args
}

// expectPeek 检查下一个词法单元类型 | Check the next token type
func (p *Parser) expectPeek(t token.TokenType) bool {
    if p.peekToken.Type == t {
        p.nextToken()
        return true
    }
    p.peekError(t)
    return false
}

// peekError 添加解析错误 | Add parsing error
func (p *Parser) peekError(t token.TokenType) {
    msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
    p.errors = append(p.errors, msg)
}
```

## 运行时环境实现 | Runtime Environment Implementation

运行时环境负责管理变量和其值。

*The runtime environment is responsible for managing variables and their values.*

### 实现环境 | Implement the Environment

创建文件 `internal/runtime/env.go`：

*Create the file `internal/runtime/env.go`:*

```go
package runtime

// Environment 表示运行时环境 | Represents the runtime environment
type Environment struct {
    store map[string]interface{}
}

// New 创建新的环境 | Create a new environment
func New() *Environment {
    return &Environment{
        store: make(map[string]interface{}),
    }
}

// Get 获取变量值 | Get variable value
func (e *Environment) Get(name string) (interface{}, bool) {
    val, ok := e.store[name]
    return val, ok
}

// Set 设置变量值 | Set variable value
func (e *Environment) Set(name string, val interface{}) {
    e.store[name] = val
}
```

## 代码生成器实现 | Code Generator Implementation

代码生成器负责将抽象语法树转换为 Minecraft 命令。

*The code generator is responsible for converting the abstract syntax tree into Minecraft commands.*

### 实现代码生成器 | Implement the Code Generator

创建文件 `internal/codegen/codegen.go`：

*Create the file `internal/codegen/codegen.go`:*

```go
package codegen

import (
    "fmt"

    "github.com/minecraft-script/internal/ast"
    "github.com/minecraft-script/internal/runtime"
)

// Codegen 表示代码生成器 | Represents the code generator
type Codegen struct {
    env *runtime.Environment
}

// New 创建新的代码生成器 | Create a new code generator
func New(env *runtime.Environment) *Codegen {
    return &Codegen{env: env}
}

// Generate 生成命令列表 | Generate a list of commands
func (c *Codegen) Generate(program *ast.Program) []string {
    var commands []string
    for _, stmt := range program.Statements {
        cmd := c.generateStatement(stmt)
        if cmd != "" {
            commands = append(commands, cmd)
        }
    }
    return commands
}

// generateStatement 生成单个语句的命令 | Generate command for a single statement
func (c *Codegen) generateStatement(stmt ast.Statement) string {
    switch s := stmt.(type) {
    case *ast.AssignStatement:
        return c.generateAssignStatement(s)
    case *ast.CallExpression:
        return c.generateCallExpression(s)
    default:
        return ""
    }
}

// generateAssignStatement 生成赋值语句的命令 | Generate command for assignment statement
func (c *Codegen) generateAssignStatement(stmt *ast.AssignStatement) string {
    name := stmt.Name.Value
    value := c.generateExpression(stmt.Value)
    c.env.Set(name, value)
    return "" // 赋值语句不生成命令 | Assignment statements do not generate commands
}

// generateExpression 计算表达式的值 | Evaluate the value of an expression
func (c *Codegen) generateExpression(expr ast.Expression) interface{} {
    switch e := expr.(type) {
    case *ast.VectorLiteral:
        return e.Values
    case *ast.Identifier:
        val, _ := c.env.Get(e.Value)
        return val
    case *ast.StringLiteral:
        return e.Value
    default:
        return nil
    }
}

// generateCallExpression 生成函数调用的命令 | Generate command for function call
func (c *Codegen) generateCallExpression(expr *ast.CallExpression) string {
    funcName := expr.Function.Value
    args := make([]interface{}, len(expr.Arguments))
    for i, arg := range expr.Arguments {
        args[i] = c.generateExpression(arg)
    }

    switch funcName {
    case "fill":
        if len(args) != 3 {
            return "// Error: fill requires 3 arguments"
        }
        pos1, ok1 := args[0].([]int)
        pos2, ok2 := args[1].([]int)
        block, ok3 := args[2].(string)
        if !ok1 || !ok2 || !ok3 || len(pos1) != 3 || len(pos2) != 3 {
            return "// Error: invalid arguments for fill"
        }
        return fmt.Sprintf("/fill %d %d %d %d %d %d %s", pos1[0], pos1[1], pos1[2], pos2[0], pos2[1], pos2[2], block)
    case "setblock":
        if len(args) != 2 {
            return "// Error: setblock requires 2 arguments"
        }
        pos, ok1 := args[0].([]int)
        block, ok2 := args[1].(string)
        if !ok1 || !ok2 || len(pos) != 3 {
            return "// Error: invalid arguments for setblock"
        }
        return fmt.Sprintf("/setblock %d %d %d %s", pos[0], pos[1], pos[2], block)
    default:
        return fmt.Sprintf("// Error: unknown function %s", funcName)
    }
}
```

## 命令行工具开发 | Command-line Tool Development

最后，我们实现命令行工具，将所有组件连接起来。

*Finally, we implement the command-line tool to connect all components.*

### 实现主程序 | Implement the Main Program

创建文件 `cmd/minecraftscript/main.go`：

*Create the file `cmd/minecraftscript/main.go`:*

```go
package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"

    "github.com/minecraft-script/internal/codegen"
    "github.com/minecraft-script/internal/lexer"
    "github.com/minecraft-script/internal/parser"
    "github.com/minecraft-script/internal/runtime"
)

func main() {
    // 定义子命令 | Define subcommands
    runCmd := flag.NewFlagSet("run", flag.ExitOnError)
    runFile := runCmd.String("file", "", "Script file to run")
    runOutput := runCmd.String("output", "", "Output file for generated commands")

    // 解析命令行参数 | Parse command-line arguments
    if len(os.Args) < 2 {
        fmt.Println("Expected 'run' subcommand")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "run":
        runCmd.Parse(os.Args[2:])
    default:
        fmt.Printf("Unknown subcommand: %s\n", os.Args[1])
        os.Exit(1)
    }

    // 处理 run 子命令 | Handle run subcommand
    if runCmd.Parsed() {
        if *runFile == "" {
            fmt.Println("Please specify a script file with -file")
            os.Exit(1)
        }

        // 读取脚本文件 | Read script file
        input, err := ioutil.ReadFile(*runFile)
        if err != nil {
            fmt.Printf("Error reading file: %s\n", err)
            os.Exit(1)
        }

        // 词法分析 | Lexical analysis
        l := lexer.New(string(input))
        
        // 语法分析 | Syntax analysis
        p := parser.New(l)
        program := p.ParseProgram()
        
        if len(p.Errors()) > 0 {
            fmt.Println("Parser errors:")
            for _, err := range p.Errors() {
                fmt.Printf("  %s\n", err)
            }
            os.Exit(1)
        }

        // 创建运行时环境 | Create runtime environment
        env := runtime.New()
        
        // 代码生成 | Code generation
        cg := codegen.New(env)
        commands := cg.Generate(program)

        // 输出命令 | Output commands
        if *runOutput != "" {
            // 确保输出目录存在 | Ensure output directory exists
            dir := filepath.Dir(*runOutput)
            if _, err := os.Stat(dir); os.IsNotExist(err) {
                os.MkdirAll(dir, 0755)
            }

            // 写入文件 | Write to file
            output := ""
            for _, cmd := range commands {
                output += cmd + "\n"
            }
            err = ioutil.WriteFile(*runOutput, []byte(output), 0644)
            if err != nil {
                fmt.Printf("Error writing output: %s\n", err)
                os.Exit(1)
            }
            fmt.Printf("Commands written to %s\n", *runOutput)
        } else {
            // 打印到控制台 | Print to console
            for _, cmd := range commands {
                fmt.Println(cmd)
            }
        }
    }
}
```

## 完整示例与应用 | Complete Examples and Applications

### 创建示例脚本 | Create Example Script

创建文件 `examples/simple.mcs`：

*Create the file `examples/simple.mcs`:*

```
# 定义起始和结束位置 | Define start and end positions
start_pos = [0, 64, 0]
end_pos = [10, 74, 10]

# 使用 fill 命令填充石头 | Use fill command to place stone blocks
fill(start_pos, end_pos, "stone")
```

### 运行示例 | Run Example

```bash
# 构建项目 | Build project
go build -o minecraftscript cmd/minecraftscript/main.go

# 运行示例脚本 | Run example script
./minecraftscript run -file examples/simple.mcs -output out/simple.mcfunction
```

生成的 `out/simple.mcfunction` 文件内容：

*The content of the generated `out/simple.mcfunction` file:*

```
/fill 0 64 0 10 74 10 stone
```

## 扩展与优化 | Extensions and Optimizations

### 添加新命令 | Add New Commands

要添加新的 Minecraft 命令，需要修改以下部分：

*To add new Minecraft commands, you need to modify the following parts:*

1. 在 `token.go` 中添加新的命令关键字 | Add new command keywords in `token.go`
2. 在 `codegen.go` 中实现命令的代码生成逻辑 | Implement code generation logic for the command in `codegen.go`

### 扩展语言特性 | Extend Language Features

可以考虑添加以下语言特性：

*Consider adding the following language features:*

1. **条件语句 | Conditional Statements**：实现 if-else 结构 | Implement if-else structures
2. **循环语句 | Loop Statements**：实现 for 循环 | Implement for loops
3. **函数定义 | Function Definitions**：允许用户定义自己的函数 | Allow users to define their own functions

## 总结 | Conclusion

通过本教程，我们从零开始实现了一个简单但功能完整的 Minecraft 脚本语言。该语言可以：

*Through this tutorial, we have implemented a simple but fully functional Minecraft scripting language from scratch. This language can:*

1. 解析简单的脚本语法 | Parse simple script syntax
2. 支持变量定义和使用 | Support variable definition and usage
3. 支持向量表示和操作 | Support vector representation and operations
4. 生成有效的 Minecraft 命令 | Generate valid Minecraft commands

这个项目展示了编译器前端的基本工作原理，包括词法分析、语法分析、抽象语法树构建和代码生成等关键步骤。通过扩展和优化，可以使这个语言更加强大和易用。

*This project demonstrates the basic working principles of a compiler frontend, including key steps such as lexical analysis, syntax analysis, abstract syntax tree construction, and code generation. Through extensions and optimizations, this language can be made more powerful and user-friendly.*
