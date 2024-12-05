package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"

	"github.com/adsmf/adventofcode/utils"
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
	order := [100][]int{}
	utils.EachLine(input, func(index int, line string) (done bool) {
		if line == "" {
			return false
		}
		if line[2] == '|' {
			x, _ := strconv.Atoi(line[0:2])
			y, _ := strconv.Atoi(line[3:5])
			order[x] = append(order[x], y)
			return false
		}
		pages := utils.GetInts(line)
		pagesUsed := [100]bool{}
		valid := true
	check:
		for _, page := range pages {
			for _, preReq := range order[page] {
				if pagesUsed[preReq] {
					valid = false
					break check
				}
			}
			pagesUsed[page] = true
		}
		if valid {
			p1 += pages[len(pages)/2]
			return false
		}
		slices.SortFunc(pages, func(a, b int) int {
			for _, preReq := range order[a] {
				if preReq == b {
					return -1
				}
			}
			return 0
		})
		p2 += pages[len(pages)/2]
		return false
	})
	return p1, p2
}

var benchmark = false
