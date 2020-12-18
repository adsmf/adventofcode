package main

import (
	"fmt"
	"go/scanner"
	"go/token"
	"strconv"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	expressions := utils.ReadInputLines("input.txt")
	p1 := part1(expressions)
	p2 := part2(expressions)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(expressions []string) int {
	precedence = map[token.Token]int{token.ADD: 1, token.MUL: 1}
	sum := 0
	for _, line := range expressions {
		sum += calculate([]byte(line))
	}
	return sum
}

func part2(expressions []string) int {
	precedence = map[token.Token]int{token.ADD: 2, token.MUL: 1}
	sum := 0
	for _, line := range expressions {
		sum += calculate([]byte(line))
	}
	return sum
}

var precedence map[token.Token]int

func calculate(input []byte) int {
	var s scanner.Scanner

	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(input))

	s.Init(file, input, nil, scanner.ScanComments)

	rpn := make([]rpnEntry, 0, 50)
	opStack := make([]token.Token, 0, 50)
	for {
		_, tok, literal := s.Scan()
		if tok == token.EOF {
			break
		}
		switch tok {
		case token.INT:
			value, _ := strconv.Atoi(literal)
			rpn = append(rpn, rpnEntry{value: value})
		case token.MUL, token.ADD:
			for len(opStack) > 0 && precedence[opStack[len(opStack)-1]] >= precedence[tok] {
				var popped token.Token
				popped, opStack = opStack[len(opStack)-1], opStack[:len(opStack)-1]
				rpn = append(rpn, rpnEntry{token: popped})
			}
			opStack = append(opStack, tok)
		case token.LPAREN:
			opStack = append(opStack, tok)
		case token.RPAREN:
			for len(opStack) > 0 {
				var popped token.Token
				popped, opStack = opStack[len(opStack)-1], opStack[:len(opStack)-1]
				if popped == token.LPAREN {
					break
				}
				rpn = append(rpn, rpnEntry{token: popped})
			}
		}
	}
	for len(opStack) > 0 {
		rpn, opStack = append(rpn, rpnEntry{token: opStack[len(opStack)-1]}), opStack[:len(opStack)-1]
	}

	return evaluateRPN(rpn)
}

func evaluateRPN(tokens []rpnEntry) int {
	stack := make([]int, 0, 50)
	for _, v := range tokens {
		switch v.token {
		case token.ADD:
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]+stack[len(stack)-1])
		case token.MUL:
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]*stack[len(stack)-1])
		default:
			stack = append(stack, v.value)
		}
	}
	return stack[0]
}

type rpnEntry struct {
	value int
	token token.Token
}

var benchmark = false
