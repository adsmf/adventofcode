package main

import (
	"container/ring"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	p1, p2 := runSim()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func runSim() (int, int) {
	fishCounts := countInitialArray()
	p1 := 0
	offset := 0
	for day := 0; day < 80; day++ {
		toUpdate := (offset + 7)
		if toUpdate > 8 {
			toUpdate -= 9
		}
		fishCounts[toUpdate] += fishCounts[offset]
		offset++
		if offset > 8 {
			offset = 0
		}
	}
	p1 = sumArray(fishCounts)
	for day := 80; day < 256; day++ {
		toUpdate := (offset + 7)
		if toUpdate > 8 {
			toUpdate -= 9
		}
		fishCounts[toUpdate] += fishCounts[offset]
		offset++
		if offset > 8 {
			offset = 0
		}
	}
	return p1, sumArray(fishCounts)
}

func countInitialArray() [9]int {
	counts := [9]int{}
	for _, ch := range input {
		if ch >= '0' {
			counts[ch-'0']++
		}
	}
	return counts
}

func sumArray(counts [9]int) int {
	total := 0
	for _, count := range counts {
		total += count
	}
	return total
}

//
// Alternative (slower) implementations for benchmark comparison
///
func runSimSlice() (int, int) {
	fishCounts := countInitial()
	p1 := 0
	for day := 0; day < 256; day++ {
		if day == 80 {
			p1 = sum(fishCounts)
		}
		nextDay := append(fishCounts[1:], fishCounts[0])
		nextDay[6] += fishCounts[0]
		fishCounts = nextDay
	}
	return p1, sum(fishCounts)
}

func countInitial() []int {
	counts := make([]int, 9, 9+256) // Pre-allocate space for 256 iterations
	for _, ch := range input {
		if ch >= '0' {
			counts[ch-'0']++
		}
	}
	return counts
}

func sum(counts []int) int {
	total := 0
	for _, count := range counts {
		total += count
	}
	return total
}
func runSimNoPreallocate() (int, int) {
	fishCounts := countInitialNoPreallocate()
	p1 := 0
	for day := 0; day < 256; day++ {
		if day == 80 {
			p1 = sum(fishCounts)
		}
		nextDay := append(fishCounts[1:], fishCounts[0])
		nextDay[6] += fishCounts[0]
		fishCounts = nextDay
	}
	return p1, sum(fishCounts)
}

func countInitialNoPreallocate() []int {
	counts := make([]int, 9, 9) // Pre-allocate space for 256 iterations
	for _, ch := range input {
		if ch >= '0' {
			counts[ch-'0']++
		}
	}
	return counts
}

func runSimNoAppend() (int, int) {
	f := countInitialArray()
	p1 := 0
	for day := 0; day < 256; day++ {
		if day == 80 {
			p1 = sumArray(f)
		}
		// More assignment operations are slower than the subslice and append
		f[0], f[1], f[2], f[3], f[4], f[5], f[6], f[7], f[8] = f[1], f[2], f[3], f[4], f[5], f[6], f[7]+f[0], f[8], f[0]
	}
	return p1, sumArray(f)
}

func runSimRing() (int, int) {
	fishCounts := countInitialNoPreallocate()
	countRing := ring.New(9)
	for i, v := 0, countRing; i < 9; i, v = i+1, v.Next() {
		v.Value = &fishCounts[i]
	}
	p1 := 0
	for day := 0; day < 256; day++ {
		if day == 80 {
			p1 = sumRing(countRing)
		}
		*(countRing.Prev().Prev().Value).(*int) += *countRing.Value.(*int)
		countRing = countRing.Next()
	}
	return p1, sumRing(countRing)
}

func sumRing(counts *ring.Ring) int {
	total := 0
	for i, v := 0, counts; i < counts.Len(); i, v = i+1, v.Next() {
		total += *(v.Value.(*int))
	}
	return total
}

var benchmark = false
