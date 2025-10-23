package codegen

import (
	"fmt"

	"github.com/minecraft-script/internal/ast"
	"github.com/minecraft-script/internal/runtime"
)

type Codegen struct {
	env *runtime.Environment
}

func New(env *runtime.Environment) *Codegen {
	return &Codegen{env: env}
}

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

func (c *Codegen) generateAssignStatement(stmt *ast.AssignStatement) string {
	name := stmt.Name.Value
	value := c.generateExpression(stmt.Value)
	c.env.Set(name, value)
	return "" // 赋值语句不生成命令 | Assignment statements do not generate commands
}

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