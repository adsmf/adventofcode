package main

import (
	_ "embed"
	"fmt"
	"slices"
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
	order := [100][100]bool{}

	orderSec := true
	pages := make([]byte, 0, 25)
	for pos := 0; pos < len(input)-1; pos++ {
		if input[pos] == '\n' {
			orderSec = false
			continue
		}
		if orderSec {
			x := (input[pos]-'0')*10 + (input[pos+1] - '0')
			y := (input[pos+3]-'0')*10 + (input[pos+4] - '0')
			pos += 5
			order[x][y] = true
			continue
		}

		for pages = pages[0:0]; pos < len(input); pos++ {
			page := (input[pos]-'0')*10 + (input[pos+1] - '0')
			pages = append(pages, page)
			pos += 2
			if pos >= len(input) || input[pos] == '\n' {
				break
			}
		}
		sortFn := func(a, b byte) int {
			if order[a][b] {
				return -1
			}
			return 0
		}
		if slices.IsSortedFunc(pages, sortFn) {
			p1 += int(pages[len(pages)/2])
			continue
		}
		slices.SortFunc(pages, sortFn)
		p2 += int(pages[len(pages)/2])
	}
	return p1, p2
}

var benchmark = false
