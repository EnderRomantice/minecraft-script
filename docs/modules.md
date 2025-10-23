# Minecraft Script 模块详解 | Module Details

## 1. 词法分析器 (Lexer)

词法分析器负责将源代码文本转换为词法单元（Token）序列，是编译过程的第一步。

*The lexer is responsible for converting source code text into a sequence of tokens, which is the first step in the compilation process.*

### 核心数据结构 | Core Data Structures

```go
type Lexer struct {
    input        string // 输入源代码 | Input source code
    position     int    // 当前字符位置 | Current character position
    readPosition int    // 下一个字符位置 | Next character position
    ch           byte   // 当前字符 | Current character
}
```

### 主要方法 | Main Methods

- `New(input string) *Lexer`：创建新的词法分析器实例并初始化 | Create and initialize a new lexer instance
- `readChar()`：读取下一个字符并更新位置指针 | Read the next character and update position pointers
- `NextToken()`：解析并返回下一个词法单元 | Parse and return the next token

### 工作流程 | Workflow

1. 初始化词法分析器，设置输入源代码 | Initialize the lexer with input source code
2. 逐字符读取源代码 | Read the source code character by character
3. 根据当前字符确定词法单元类型 | Determine token type based on current character
4. 生成词法单元并返回 | Generate and return the token
5. 重复步骤2-4直到源代码结束 | Repeat steps 2-4 until the end of source code

## 2. 词法单元 (Token)

词法单元模块定义了语言中的基本词法单元类型和结构。

*The token module defines the basic token types and structures in the language.*

### 核心数据结构 | Core Data Structures

```go
type TokenType string

type Token struct {
    Type    TokenType
    Literal string
}
```

### 主要常量 | Main Constants

- `ILLEGAL`：非法字符 | Illegal character
- `EOF`：文件结束 | End of file
- `IDENT`：标识符 | Identifier
- `NUMBER`：数字 | Number
- `STRING`：字符串 | String
- `ASSIGN`：赋值符号 `=` | Assignment operator
- `COMMA`：逗号 `,` | Comma
- `LPAREN`/`RPAREN`：左右括号 `()` | Left/right parentheses
- `LBRACKET`/`RBRACKET`：左右方括号 `[]` | Left/right brackets
- `FILL`/`SETBLOCK`：Minecraft 命令关键字 | Minecraft command keywords

## 3. 抽象语法树 (AST)

抽象语法树模块定义了语言的语法结构和节点类型。

*The abstract syntax tree module defines the syntactic structure and node types of the language.*

### 核心接口 | Core Interfaces

```go
type Node interface {
    TokenLiteral() string
}

type Statement interface {
    Node
    statementNode()
}

type Expression interface {
    Node
    expressionNode()
}
```

### 主要节点类型 | Main Node Types

- `Program`：表示整个程序，包含多个语句 | Represents the entire program, containing multiple statements
- `AssignStatement`：变量赋值语句，如 `pos = [1, 2, 3]` | Variable assignment statement
- `Identifier`：标识符表达式，如变量名 | Identifier expression, such as variable names
- `VectorLiteral`：向量字面量表达式，如 `[1, 2, 3]` | Vector literal expression
- `StringLiteral`：字符串字面量，如 `"stone"` | String literal
- `CallExpression`：函数调用表达式，如 `fill(pos1, pos2, "stone")` | Function call expression

### 节点关系 | Node Relationships

- `Program` 包含多个 `Statement` | Program contains multiple Statements
- `AssignStatement` 包含一个 `Identifier`（左值）和一个 `Expression`（右值）| AssignStatement contains an Identifier (left value) and an Expression (right value)
- `CallExpression` 包含一个 `Identifier`（函数名）和多个 `Expression`（参数）| CallExpression contains an Identifier (function name) and multiple Expressions (parameters)

## 4. 语法分析器 (Parser)

语法分析器负责将词法单元序列转换为抽象语法树，实现了递归下降解析算法。

*The parser is responsible for converting a sequence of tokens into an abstract syntax tree, implementing a recursive descent parsing algorithm.*

### 核心数据结构 | Core Data Structures

```go
type Parser struct {
    l         *lexer.Lexer
    curToken  token.Token
    peekToken token.Token
    errors    []string
}
```

### 主要方法 | Main Methods

- `New(l *lexer.Lexer) *Parser`：创建新的语法分析器实例 | Create a new parser instance
- `nextToken()`：前进到下一个词法单元 | Advance to the next token
- `ParseProgram() *ast.Program`：解析整个程序 | Parse the entire program
- `parseStatement() ast.Statement`：解析单个语句 | Parse a single statement
- `parseExpression() ast.Expression`：解析表达式 | Parse an expression

### 解析流程 | Parsing Process

1. 初始化语法分析器，设置词法分析器 | Initialize the parser with a lexer
2. 解析程序，创建 `Program` 节点 | Parse the program, create a Program node
3. 循环解析语句，添加到 `Program.Statements` | Loop to parse statements, add to Program.Statements
4. 根据当前词法单元类型选择相应的解析函数 | Choose appropriate parsing function based on current token type
5. 递归解析表达式和子表达式 | Recursively parse expressions and sub-expressions

## 5. 代码生成器 (Codegen)

代码生成器负责将抽象语法树转换为 Minecraft 命令。

*The code generator is responsible for converting the abstract syntax tree into Minecraft commands.*

### 核心数据结构 | Core Data Structures

```go
type Codegen struct {
    env *runtime.Environment
}
```

### 主要方法 | Main Methods

- `New(env *runtime.Environment) *Codegen`：创建新的代码生成器实例 | Create a new code generator instance
- `Generate(program *ast.Program) []string`：生成命令列表 | Generate a list of commands
- `generateStatement(stmt ast.Statement) string`：生成单个语句的命令 | Generate command for a single statement
- `generateExpression(expr ast.Expression) interface{}`：计算表达式的值 | Evaluate the value of an expression

### 生成流程 | Generation Process

1. 初始化代码生成器，设置环境 | Initialize the code generator with an environment
2. 遍历 AST 中的每个语句 | Traverse each statement in the AST
3. 根据语句类型生成相应的命令 | Generate appropriate commands based on statement type
4. 处理变量赋值和函数调用 | Handle variable assignments and function calls
5. 返回生成的命令列表 | Return the generated command list

## 6. 运行时环境 (Runtime Environment)

运行时环境负责管理变量和其值。

*The runtime environment is responsible for managing variables and their values.*

### 核心数据结构 | Core Data Structures

```go
type Environment struct {
    store map[string]interface{}
}
```

### 主要方法 | Main Methods

- `New() *Environment`：创建新的环境实例 | Create a new environment instance
- `Get(name string) (interface{}, bool)`：获取变量值 | Get variable value
- `Set(name string, val interface{})`：设置变量值 | Set variable value

### 使用方式 | Usage

1. 创建环境实例 | Create an environment instance
2. 在代码生成过程中使用环境存储变量值 | Use the environment to store variable values during code generation
3. 在生成命令时查询变量值 | Query variable values when generating commands