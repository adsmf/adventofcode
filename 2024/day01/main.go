package main

import (
	_ "embed"
	"fmt"
	"sort"

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
	l1, l2 := make([]int, 0, len(input)>>3), make([]int, 0, len(input)>>3)
	utils.EachInteger(input, func(index, value int) (done bool) {
		l1 = append(l1, value)
		l1, l2 = l2, l1
		return false
	})
	sort.Ints(l1)
	sort.Ints(l2)
	p1, p2 := 0, 0
	l2ptr := 0
	for i := range len(l1) {
		n1, n2 := l1[i], l2[i]
		if n1 > n2 {
			p1 += n1 - n2
		} else {
			p1 += n2 - n1
		}
		for ; l2ptr < len(l2) && l2[l2ptr] < n1; l2ptr++ {
		}
		for ; l2ptr < len(l2) && l2[l2ptr] == n1; l2ptr++ {
			p2 += n1
		}
	}
	return p1, p2
}

var benchmark = false
