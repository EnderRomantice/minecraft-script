package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/minecraft-script/internal/codegen"
	"github.com/minecraft-script/internal/lexer"
	"github.com/minecraft-script/internal/parser"
	"github.com/minecraft-script/internal/runtime"
)

func main() {
	// 定义子命令 | Define subcommands
	runCmd := flag.NewFlagSet("run", flag.ExitOnError)
	
	// run 子命令的参数 | Run subcommand arguments
	inputFile := runCmd.String("file", "", "Input script file")
	outputFile := runCmd.String("output", "", "Output command file")

	// 解析命令行参数 | Parse command line arguments
	if len(os.Args) < 2 {
		fmt.Println("Expected 'run' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		runCmd.Parse(os.Args[2:])
		if *inputFile == "" {
			fmt.Println("Please specify an input file with -file")
			os.Exit(1)
		}
		runScript(*inputFile, *outputFile)
	default:
		fmt.Printf("Unknown subcommand: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func runScript(inputFile, outputFile string) {
	// 读取输入文件 | Read input file
	input, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	// 创建词法分析器 | Create lexer
	l := lexer.New(string(input))
	
	// 创建语法分析器 | Create parser
	p := parser.New(l)
	
	// 解析程序 | Parse program
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
	
	// 创建代码生成器 | Create code generator
	cg := codegen.New(env)
	
	// 生成命令 | Generate commands
	commands := cg.Generate(program)
	
	// 输出命令 | Output commands
	output := strings.Join(commands, "\n")
	if outputFile != "" {
		// 确保输出目录存在 | Ensure output directory exists
		dir := filepath.Dir(outputFile)
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("Error creating output directory: %v\n", err)
			os.Exit(1)
		}
		
		// 写入输出文件 | Write output file
		if err := ioutil.WriteFile(outputFile, []byte(output), 0644); err != nil {
			fmt.Printf("Error writing output file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Commands written to %s\n", outputFile)
	} else {
		// 输出到控制台 | Output to console
		fmt.Println(output)
	}
}