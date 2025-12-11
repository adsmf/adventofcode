package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input []byte

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	cache := map[search]int{}
	mappings := map[string][]string{}
	utils.EachLine(input, func(index int, line []byte) (done bool) {
		in := line[0:3]
		outFwd := make([]string, 0, (len(line)/4)-1)
		for i := 5; i < len(line); i += 4 {
			out := line[i : i+3]
			outFwd = append(outFwd, string(out))
		}
		mappings[string(in)] = outFwd
		return
	})
	p1 := countWays(cache, mappings, "you", false, false)
	p2 := countWays(cache, mappings, "svr", true, true)
	return p1, p2
}

func countWays(cache map[search]int, mappings map[string][]string, node string, needDAC, needFFT bool) int {
	if cached, found := cache[search{node, needDAC, needFFT}]; found {
		return cached
	}
	if node == "out" {
		if needDAC || needFFT {
			return 0
		}
		return 1
	}
	if node == "dac" {
		needDAC = false
	}
	if node == "fft" {
		needFFT = false
	}
	sub := 0
	for _, next := range mappings[node] {
		sub += countWays(cache, mappings, next, needDAC, needFFT)
	}
	cache[search{node, needDAC, needFFT}] = sub
	return sub
}

type search struct {
	node    string
	needDAC bool
	needFFT bool
}

var benchmark = false
