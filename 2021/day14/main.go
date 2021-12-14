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
	initial, rules := load(input)
	p1, p2 := expandPolymers(initial, rules)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func expandPolymers(polymerString string, rules map[pairVal]byte) (int, int) {
	pairs := make(map[pairVal]int)
	firstCh, lastCh := byte(polymerString[0]), byte(polymerString[len(polymerString)-1])
	last := polymerString[0]
	for i := 1; i < len(polymerString); i++ {
		pair := stringToPairVal(string(last) + string(polymerString[i]))
		pairs[pair]++
		last = polymerString[i]
	}
	p1 := 0
	for i := 0; i < 40; i++ {
		if i == 10 {
			p1 = diffElements(pairs, firstCh, lastCh)
		}
		nextPairs := make(map[pairVal]int, len(pairs))
		for pair, count := range pairs {
			ch1, ch2 := pair.chars()
			replace := rules[pair]
			rep1 := bytesToPairVal(ch1, replace)
			rep2 := bytesToPairVal(replace, ch2)
			nextPairs[rep1] += count
			nextPairs[rep2] += count
		}

		pairs = nextPairs
	}
	return p1, diffElements(pairs, firstCh, lastCh)
}

func diffElements(pairs map[pairVal]int, first, last byte) int {
	counts := make(map[byte]int, 26)
	counts[first]++
	counts[last]++
	for pair, count := range pairs {
		p1, p2 := pair.chars()
		counts[p1] += count
		counts[p2] += count
	}
	min, max := utils.MaxInt, 0
	for _, count := range counts {
		if min > count {
			min = count
		}
		if max < count {
			max = count
		}
	}
	return (max - min) / 2
}

func load(in string) (string, map[pairVal]byte) {
	polymer := strings.Builder{}
	pos := 0
	for ; in[pos] != '\n'; pos++ {
		polymer.WriteByte(in[pos])
	}
	pos += 2
	rules := make(map[pairVal]byte, (len(in)-pos)/8)
	for ; pos < len(in); pos += 8 {
		rules[stringToPairVal(in[pos:pos+2])] = in[pos+6]
	}
	return polymer.String(), rules
}

type pairVal uint16

func stringToPairVal(pair string) pairVal { return pairVal(pair[0])<<8 + pairVal(pair[1]) }
func bytesToPairVal(p1, p2 byte) pairVal  { return pairVal(p1)<<8 + pairVal(p2) }
func (p pairVal) chars() (byte, byte)     { return byte(p >> 8), byte(p & 0xff) }

var benchmark = false
