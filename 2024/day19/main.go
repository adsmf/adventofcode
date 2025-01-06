package main

import (
	_ "embed"
	"fmt"
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
	p1 := 0
	p2 := 0
	towels := make([]string, 0, 500)
	s := 0
parseTowels:
	for pos, ch := range input {
		switch ch {
		case '\n':
			towels = append(towels, input[s:pos])
			break parseTowels
		case ',':
			towels = append(towels, input[s:pos])
			s = pos + 1
		case ' ':
			s++
		}
	}
	utils.EachLine(input[s-1:], func(index int, line string) (done bool) {
		possible := matches(line, towels)
		if possible > 0 {
			p1++
			p2 += possible
		}
		return
	})
	return p1, p2
}

var cache = map[string]int{}

func matches(pattern string, towels []string) int {
	if cached, found := cache[(pattern)]; found {
		return cached
	}
	count := 0
	for _, towel := range towels {
		if pattern == towel {
			count++
			continue
		}
		if strings.HasPrefix(pattern, towel) {
			count += matches(pattern[len(towel):], towels)
		}
	}
	cache[pattern] = count
	return count
}

var benchmark = false
