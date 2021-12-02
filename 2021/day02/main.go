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
	p1, p2 := followRoute()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func followRoute() (int, int) {
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

type position struct {
	x, aim, depth int
}

var benchmark = false
