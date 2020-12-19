package main

import (
	"fmt"
	"io/ioutil"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1, p2 := calc()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func calc() (int, int) {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	banks := utils.GetInts(string(inputBytes))
	states := map[string]int{}
	cycles := 0

	for {
		cycles++
		redistribute(banks)
		hash := hashBanks(banks)
		if states[hash] > 0 {
			repCount := 1
			for redistribute(banks); hashBanks(banks) != hash; repCount++ {
				redistribute(banks)
			}
			return cycles, repCount
		}
		states[hash]++
	}
}

func hashBanks(banks []int) string { return fmt.Sprintf("%v", banks) }

func redistribute(banks []int) []int {
	max, maxPos := -1, -1

	for pos, val := range banks {
		if val > max {
			max, maxPos = val, pos
		}
	}

	banks[maxPos] = 0
	allPlus := max / len(banks)
	remainder := max % len(banks)

	for i := 0; i < len(banks); i++ {
		banks[i] += allPlus
	}

	for i := 1; i <= remainder; i++ {
		banks[(maxPos+i)%len(banks)]++
	}
	return banks
}

var benchmark = false
