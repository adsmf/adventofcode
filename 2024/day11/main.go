package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

const cacheSize = 150_000

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	p1, p2 := 0, 0
	utils.EachInteger(input, func(_, stone int) (done bool) {
		p1 += countIter(stone, 25)
		p2 += countIter(stone, 75)
		return false
	})
	return p1, p2
}

var cache [cacheSize]int

func countIter(val int, steps int8) int {
	useCache := false
	var h uint32
	if int(steps) < 1<<7 {
		h = uint32(val<<7) + uint32(steps)
		if int(h) < len(cache) {
			useCache = true
		}
	}
	if useCache && cache[h] != 0 {
		return cache[h]
	}
	if steps == 0 {
		return 1
	}
	save := func(res int) int {
		if useCache {
			cache[h] = res
		}
		return res
	}
	if val == 0 {
		return save(countIter(1, steps-1))
	}
	s1, s2, evenLen := split(val)
	if evenLen {
		return save(countIter(s1, steps-1) + countIter(s2, steps-1))
	}
	return save(countIter(val*2024, steps-1))
}

func split(val int) (int, int, bool) {
	tens := 0
	for c := val; c > 0; c /= 10 {
		tens++
	}
	if tens&1 == 1 {
		return 0, 0, false
	}
	mult := 1
	for i := 0; i < tens/2; i++ {
		mult *= 10
	}
	first := val / mult
	second := val - first*mult
	return first, second, tens&1 == 0
}

var benchmark = false
