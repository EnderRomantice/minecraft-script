# Minecraft Script 项目文档 | Project Documentation

## 项目概述 | Project Overview

Minecraft Script 是一个专门为 Minecraft 游戏设计的脚本语言解释器，旨在简化 Minecraft 命令的编写和执行。该项目使用 Go 语言实现，包含完整的词法分析、语法分析、抽象语法树构建和代码生成等编译器前端组件。

*Minecraft Script is an interpreter designed specifically for Minecraft game, aiming to simplify the writing and execution of Minecraft commands. This project is implemented in Go language and includes complete compiler frontend components such as lexical analysis, syntax analysis, abstract syntax tree construction, and code generation.*

## 项目架构 | Project Architecture

项目采用标准的编译器前端架构，主要包含以下模块：

*The project adopts a standard compiler frontend architecture, mainly containing the following modules:*

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

## 处理流程 | Processing Flow

Minecraft Script 的处理流程如下：

*The processing flow of Minecraft Script is as follows:*

1. **词法分析 | Lexical Analysis**：将源代码文本转换为词法单元（Token）序列 | Convert source code text into a sequence of tokens
2. **语法分析 | Syntax Analysis**：将词法单元序列解析为抽象语法树（AST）| Parse the token sequence into an abstract syntax tree (AST)
3. **代码生成 | Code Generation**：将抽象语法树转换为 Minecraft 命令 | Convert the abstract syntax tree into Minecraft commands
4. **执行 | Execution**：在 Minecraft 环境中执行生成的命令 | Execute the generated commands in the Minecraft environment

## 语言特性 | Language Features

Minecraft Script 支持以下语言特性：

*Minecraft Script supports the following language features:*

1. **变量定义与使用 | Variable Definition and Usage**
   ```
   # 定义变量 | Define variables
   start_pos = [0, 64, 0]
   end_pos = [10, 74, 10]
   
   # 使用变量 | Use variables
   fill(start_pos, end_pos, "stone")
   ```

2. **向量字面量 | Vector Literals**
   ```
   # 使用方括号定义三维坐标 | Define 3D coordinates using square brackets
   pos = [10, 65, 20]
   ```

3. **内置命令 | Built-in Commands**
   - `fill(pos1, pos2, block)` - 填充区域 | Fill an area
   - `setblock(pos, block)` - 放置方块 | Place a block

## 使用方法 | Usage

### 安装 | Installation

```bash
# 克隆仓库 | Clone repository
git clone https://github.com/minecraft-script/minecraft-script.git
cd minecraft-script

# 构建项目 | Build project
go build -o minecraftscript cmd/minecraftscript/main.go
```

### 运行脚本 | Running Scripts

```bash
# 运行脚本并输出到文件 | Run script and output to file
./minecraftscript run -file examples/simple.mcs -output out/simple.mcfunction

# 运行脚本并打印到控制台 | Run script and print to console
./minecraftscript run -file examples/simple.mcs
```

### 示例脚本 | Example Script

```
# 定义起始和结束位置 | Define start and end positions
start_pos = [0, 64, 0]
end_pos = [10, 74, 10]

# 使用 fill 命令填充石头 | Use fill command to place stone blocks
fill(start_pos, end_pos, "stone")
```

## 开发指南 | Development Guide

### 添加新命令 | Adding New Commands

要添加新的 Minecraft 命令支持，需要修改以下文件：

*To add support for new Minecraft commands, you need to modify the following files:*

1. `internal/token/token.go` - 添加新命令的词法单元 | Add token for the new command
2. `internal/codegen/codegen.go` - 实现命令的代码生成逻辑 | Implement code generation logic for the command

### 扩展语言特性 | Extending Language Features

要扩展语言特性，可能需要修改以下模块：

*To extend language features, you may need to modify the following modules:*

1. `internal/lexer/lexer.go` - 添加新的词法规则 | Add new lexical rules
2. `internal/parser/parser.go` - 添加新的语法规则 | Add new syntax rules
3. `internal/ast/ast.go` - 添加新的 AST 节点类型 | Add new AST node types

## 文档 | Documentation

更多详细文档请参考：

*For more detailed documentation, please refer to:*

- [模块详解 | Module Details](docs/modules.md)
- [教程 | Tutorial](docs/TUTORIAL.md)
- [示例 | Examples](docs/examples.md)

## 许可证 | License

本项目采用 MIT 许可证。详情请参阅 LICENSE 文件。

*This project is licensed under the MIT License. See the LICENSE file for details.*
