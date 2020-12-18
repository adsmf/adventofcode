package main

import (
	"fmt"

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
	precedence = equalPrecedence
	sum := 0
	for _, line := range expressions {
		sum += calculate([]byte(line))
	}
	return sum
}

func part2(expressions []string) int {
	precedence = highAddPrecedence
	sum := 0
	for _, line := range expressions {
		sum += calculate([]byte(line))
	}
	return sum
}

func calculate(input []byte) int {
	rpn := make([]token, 0, 50)
	opStack := make([]tokenType, 0, 50)
	tokens := tokenize(input)
	for _, tok := range tokens {
		switch tok.token {
		case tokenMultiply, tokenAdd:
			for len(opStack) > 0 && precedence[opStack[len(opStack)-1]] >= precedence[tok.token] {
				var popped tokenType
				popped, opStack = opStack[len(opStack)-1], opStack[:len(opStack)-1]
				rpn = append(rpn, token{token: popped})
			}
			opStack = append(opStack, tok.token)
		case tokenLeftParen:
			opStack = append(opStack, tok.token)
		case tokenRightParen:
			for len(opStack) > 0 {
				var popped tokenType
				popped, opStack = opStack[len(opStack)-1], opStack[:len(opStack)-1]
				if popped == tokenLeftParen {
					break
				}
				rpn = append(rpn, token{token: popped})
			}
		case tokenTypeLiteral:
			rpn = append(rpn, tok)
		}
	}
	for len(opStack) > 0 {
		rpn, opStack = append(rpn, token{token: opStack[len(opStack)-1]}), opStack[:len(opStack)-1]
	}

	return evaluateRPN(rpn)
}

func evaluateRPN(tokens []token) int {
	stack := make([]int, 0, 50)
	for _, v := range tokens {
		switch v.token {
		case tokenAdd:
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]+stack[len(stack)-1])
		case tokenMultiply:
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]*stack[len(stack)-1])
		default:
			stack = append(stack, v.value)
		}
	}
	return stack[0]
}

func tokenize(input []byte) []token {
	tokens := make([]token, 0, 50)
	acc, hasLiteral := 0, false
	for _, char := range input {
		if char >= '0' && char <= '9' {
			hasLiteral = true
			acc = acc*10 + int(char-'0')
			continue
		}
		if hasLiteral {
			tokens = append(tokens, token{value: acc})
			hasLiteral, acc = false, 0
		}
		switch char {
		case '(':
			tokens = append(tokens, token{token: tokenLeftParen})
		case ')':
			tokens = append(tokens, token{token: tokenRightParen})
		case '+':
			tokens = append(tokens, token{token: tokenAdd})
		case '*':
			tokens = append(tokens, token{token: tokenMultiply})
		}
	}
	if hasLiteral {
		tokens = append(tokens, token{value: acc})
	}
	return tokens
}

var precedence map[tokenType]int
var equalPrecedence = map[tokenType]int{tokenAdd: 1, tokenMultiply: 1}
var highAddPrecedence = map[tokenType]int{tokenAdd: 2, tokenMultiply: 1}

type token struct {
	value int
	token tokenType
}

type tokenType int

const (
	tokenTypeLiteral tokenType = iota
	tokenAdd
	tokenMultiply
	tokenLeftParen
	tokenRightParen
)

var benchmark = false
