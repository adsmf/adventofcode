package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	p1, p2 := 0, 0
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	cmds := re.FindAllStringSubmatch(input, -1)
	enabled := true
	for _, cmd := range cmds {
		switch cmd[0] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			v1, _ := strconv.Atoi(cmd[1])
			v2, _ := strconv.Atoi(cmd[2])
			mult := v1 * v2
			p1 += mult
			if enabled {
				p2 += mult
			}
		}
	}
	return p1, p2
}

var benchmark = false
