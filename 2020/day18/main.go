package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	sum := 0
	for _, line := range utils.ReadInputLines("input.txt") {
		sum += calculate(line, false)
	}
	return sum
}

func part2() int {
	sum := 0
	for _, line := range utils.ReadInputLines("input.txt") {
		sum += calculate(line, true)
	}
	return sum
}

func calculate(input string, part2 bool) int {
	fakeInput := strings.ReplaceAll(input, "*", "-")
	if part2 {
		fakeInput = strings.ReplaceAll(fakeInput, "+", "*")
	}
	expr, _ := parser.ParseExpr(fakeInput)
	return evalExpression(expr)
}

func evalExpression(expr ast.Expr) int {
	switch node := expr.(type) {
	case *ast.BasicLit:
		val, _ := strconv.Atoi(node.Value)
		return val
	case *ast.ParenExpr:
		return evalExpression(node.X)
	case *ast.BinaryExpr:
		switch node.Op {
		case token.ADD:
			return evalExpression(node.X) + evalExpression(node.Y)
		case token.SUB:
			return evalExpression(node.X) * evalExpression(node.Y)
		case token.MUL:
			return evalExpression(node.X) + evalExpression(node.Y)
		default:
			panic(fmt.Errorf("Unsupported operation: %v", node.Op))
		}
	default:
		panic(fmt.Errorf("Unsupported expression: %v", reflect.TypeOf(node)))
	}
}

var benchmark = false
