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
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	allIDs := utils.GetInts(input)
	sum := 0
	pairs := [][2]int{}
	for i := 0; i < len(allIDs); i += 2 {
		pairs = append(pairs, [2]int{allIDs[i], allIDs[i+1]})
	}
	for _, pair := range pairs {
		for id := pair[0]; id <= pair[1]; id++ {
			idStr := strconv.Itoa(id)
			if len(idStr)&1 != 0 {
				continue
			}
			half := len(idStr) / 2
			if idStr[:half] == idStr[half:] {
				sum += id
			}
		}
	}
	return sum
}

func part2() int {
	allIDs := utils.GetInts(input)
	sum := 0
	pairs := [][2]int{}
	for i := 0; i < len(allIDs); i += 2 {
		pairs = append(pairs, [2]int{allIDs[i], allIDs[i+1]})
	}
	for _, pair := range pairs {
		for id := pair[0]; id <= pair[1]; id++ {
			idStr := strconv.Itoa(id)
			for reps := 2; reps <= len(idStr); reps++ {
				if len(idStr)%reps != 0 {
					continue
				}
				if idStr == strings.Repeat(idStr[:len(idStr)/reps], reps) {
					sum += id
					break
				}
			}
		}
	}
	return sum
}

var benchmark = false
