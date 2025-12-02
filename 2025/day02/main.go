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
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	allIDs := utils.GetInts(input)
	p1, p2 := 0, 0
	for i := 0; i < len(allIDs); i += 2 {
		for id := allIDs[i]; id <= allIDs[i+1]; id++ {
			idStr := strconv.Itoa(id)
			for reps := 2; reps <= len(idStr); reps++ {
				if len(idStr)%reps != 0 {
					continue
				}
				if idStr == strings.Repeat(idStr[:len(idStr)/reps], reps) {
					if reps == 2 {
						p1 += id
					}
					p2 += id
					break
				}
			}
		}
	}
	return p1, p2
}

var benchmark = false
