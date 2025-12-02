package main

import (
	_ "embed"
	"fmt"
	"math"

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
	p1, p2 := 0, 0
	start := 0
	utils.EachInteger(input, func(index, value int) (done bool) {
		if index&1 == 0 {
			start = value
			return
		}
		end := value
		for id := start; id <= end; id++ {
			digits := int(math.Log10(float64(id))) + 1
			for reps := 2; reps <= digits; reps++ {
				if digits%reps != 0 {
					continue
				}
				chunkDigits := digits / reps
				if !digitsRepeat(id, reps, chunkDigits) {
					continue
				}

				if reps == 2 {
					p1 += id
				}
				p2 += id
				break
			}
		}
		return
	})
	return p1, p2
}

func digitsRepeat(value int, reps, chunkDigits int) bool {
	m := 1
	for range chunkDigits {
		m *= 10
	}
	search := value % m
	for i := 2; i <= reps; i++ {
		value /= m
		if value%m != search {
			return false
		}
	}
	return true
}

var benchmark = false
