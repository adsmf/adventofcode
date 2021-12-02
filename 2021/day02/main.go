package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := followRouteFast()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

// Initial (and more readable) implementation
func followRouteInitial() (int, int) {
	x := 0
	depthP1 := 0
	depthP2 := 0
	for _, line := range utils.GetLines(input) {
		parts := strings.Split(line, " ")
		val, _ := strconv.Atoi(parts[1])
		switch parts[0] {
		case "forward":
			x += val
			depthP2 += val * depthP1
		case "down":
			depthP1 += val
		case "up":
			depthP1 -= val
		}
	}
	return x * depthP1, x * depthP2
}

// Prefering performance over readability
func followRouteFast() (int, int) {
	x := 0
	depthP1 := 0
	depthP2 := 0
	inputLen := len(input)
	for i := 0; i < inputLen; i++ {
		ch := input[i]
		switch ch {
		case 'f':
			val := int(input[i+8] - '0')
			x += val
			depthP2 += val * depthP1
			i += 9
		case 'd':
			val := int(input[i+5] - '0')
			depthP1 += val
			i += 6
		case 'u':
			val := int(input[i+3] - '0')
			depthP1 -= val
			i += 4
		}
	}
	return x * depthP1, x * depthP2
}

var benchmark = false
