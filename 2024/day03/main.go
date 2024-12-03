package main

import (
	_ "embed"
	"fmt"
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
	pos := 0
	enabled := true
	for ; ; pos++ {
		for ; pos < len(input) && input[pos] != '('; pos++ {
		}
		if pos >= len(input) {
			return p1, p2
		}
		if pos > 5 && input[pos-5:pos] == "don't" && input[pos+1] == ')' {
			enabled = false
			continue
		}
		if pos > 2 && input[pos-2:pos] == "do" && input[pos+1] == ')' {
			enabled = true
			continue
		}
		if pos > 3 && input[pos-3:pos] == "mul" {
			vals := [2]int{}
			arg := 1
			for pos++; pos < len(input); pos++ {
				ch := input[pos]
				if ch == ')' {
					mul := vals[0] * vals[1]
					p1 += mul
					if enabled {
						p2 += mul
					}
					break
				}
				if ch == ',' {
					arg++
					if arg > 2 {
						break
					}
					continue
				}
				if ch >= '0' && ch <= '9' {
					vals[arg-1] *= 10
					vals[arg-1] += int(ch - '0')
					continue
				}
				break
			}
		}
	}
}

var benchmark = false
