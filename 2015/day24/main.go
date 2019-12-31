package main

import (
	"fmt"
	"github.com/adsmf/adventofcode/utils"
	"io/ioutil"
	"sort"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	p := loadInput("input.txt")
	total := totalWeight(p)
	target := total / 3
	fmt.Printf("Packages: %v\nTotal: %d (3*%d)\n", p, total, target)
	return findSmallest3(p, target)
}

func part2() int {
	p := loadInput("input.txt")
	total := totalWeight(p)
	target := total / 4
	fmt.Printf("Packages: %v\nTotal: %d (4*%d)\n", p, total, target)
	return findSmallest4(p, target)
}

func findSmallest3(packages []int, targetWeight int) int {
	validGroupBits := map[int]bool{}
	validByBitCount := map[int][]int{}

	fmt.Println("Enumerating potentially valid compartments...")
	for i := 0; i < 1<<len(packages); i++ {
		weight := 0
		numBits := 0
		for bit := 0; bit < len(packages); bit++ {
			if (i & (1 << bit)) > 0 {
				weight += packages[bit]
				numBits++
			}
			if weight > targetWeight {
				break
			}
		}
		if weight == targetWeight {
			validGroupBits[i] = true
			if val, found := validByBitCount[numBits]; found {
				validByBitCount[numBits] = append(val, i)
			} else {
				validByBitCount[numBits] = []int{i}
			}
		}
	}

	sizes := []int{}
	for size := range validByBitCount {
		sizes = append(sizes, size)
	}
	sort.Ints(sizes)

	fmt.Println("Checking groups by size...")
	var opts []int
	allBits := (1 << len(packages)) - 1
	for _, size := range sizes {
		validOfSize := []int{}

		for _, try := range validByBitCount[size] {
		nextTry:
			for g2 := range validGroupBits {
				if g2&try > 0 {
					continue
				}
				for g3 := range validGroupBits {
					if try|g2|g3 == allBits {
						validOfSize = append(validOfSize, try)
						break nextTry
					}
				}
			}
		}
		if len(validOfSize) > 0 {
			opts = validOfSize
			break
		}
	}

	fmt.Println("Finding best QE...")

	lowestQE := -1
	for _, group := range opts {
		qe := 1
		for bit := 0; bit < len(packages); bit++ {
			if (group & (1 << bit)) > 0 {
				qe *= packages[bit]
			}
		}
		if lowestQE == -1 || lowestQE > qe {
			lowestQE = qe
		}
	}
	return lowestQE
}

func findSmallest4(packages []int, targetWeight int) int {
	validGroupBits := map[int]bool{}
	validByBitCount := map[int][]int{}

	fmt.Println("Enumerating potentially valid compartments...")
	for i := 0; i < 1<<len(packages); i++ {
		weight := 0
		numBits := 0
		for bit := 0; bit < len(packages); bit++ {
			if (i & (1 << bit)) > 0 {
				weight += packages[bit]
				numBits++
			}
			if weight > targetWeight {
				break
			}
		}
		if weight == targetWeight {
			validGroupBits[i] = true
			if val, found := validByBitCount[numBits]; found {
				validByBitCount[numBits] = append(val, i)
			} else {
				validByBitCount[numBits] = []int{i}
			}
		}
	}

	sizes := []int{}
	for size := range validByBitCount {
		sizes = append(sizes, size)
	}
	sort.Ints(sizes)

	fmt.Println("Checking groups by size...")
	var opts []int
	allBits := (1 << len(packages)) - 1
	for _, size := range sizes {
		validOfSize := []int{}

		for _, try := range validByBitCount[size] {
		nextTry:
			for g2 := range validGroupBits {
				if g2&try > 0 {
					continue
				}
				for g3 := range validGroupBits {
					if g3&g2 > 0 || g3&try > 0 {
						continue
					}
					for g4 := range validGroupBits {
						if try|g2|g3|g4 == allBits {
							validOfSize = append(validOfSize, try)
							break nextTry
						}
					}
				}
			}
		}
		if len(validOfSize) > 0 {
			opts = validOfSize
			break
		}
	}

	fmt.Println("Finding best QE...")

	lowestQE := -1
	for _, group := range opts {
		qe := 1
		for bit := 0; bit < len(packages); bit++ {
			if (group & (1 << bit)) > 0 {
				qe *= packages[bit]
			}
		}
		if lowestQE == -1 || lowestQE > qe {
			lowestQE = qe
		}
	}
	return lowestQE
}

func totalWeight(packages []int) int {
	sum := 0
	for _, p := range packages {
		sum += p
	}
	return sum
}

func loadInput(filename string) []int {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return utils.GetInts(string(raw))
}
