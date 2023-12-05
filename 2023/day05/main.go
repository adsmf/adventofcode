package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := analyse()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func analyse() (int, int) {
	var haveSeeds []int
	currentMap := -1
	mappings := seedMapList{}
	utils.EachLine(input, func(index int, line string) (done bool) {
		if strings.TrimSpace(line) == "" {
			return false
		}
		switch {
		case strings.HasPrefix(line, "seeds:"):
			haveSeeds = utils.GetInts(line)
		case strings.HasSuffix(line, "map:"):
			currentMap++
		default:
			parts := utils.GetInts(line)
			if len(parts) != 3 {
				fmt.Println("Not seed range:", line)
				return true
			}
			sr := seedRange{
				dest:  parts[0],
				src:   parts[1],
				width: parts[2],
			}
			mappings[currentMap] = append(mappings[currentMap], sr)
		}
		return false
	})

	lowestP1 := -1
	for _, seed := range haveSeeds {
		for i := 0; i < len(mappings); i++ {
			seed, _ = mappings[i].convertSkip(seed)
		}
		if lowestP1 == -1 || seed < lowestP1 {
			lowestP1 = seed
		}
	}

	lowestP2 := -1
	flattened := mappings.flatten()
	sort.Slice(flattened, func(i, j int) bool {
		return flattened[i].dest < flattened[j].dest
	})
	for _, seedRange := range flattened {
		for i := 0; i < len(haveSeeds); i += 2 {
			if haveSeeds[i]+haveSeeds[i+1]-1 >= seedRange.src &&
				haveSeeds[i] <= seedRange.src+seedRange.width-1 {
				bestSeed := haveSeeds[i]
				if bestSeed < seedRange.src {
					bestSeed = seedRange.src
				}
				lowestP2, _ = flattened.convertSkip(bestSeed)
				return lowestP1, lowestP2
			}
		}
	}
	return lowestP1, lowestP2
}

type seedMapList [7]seedMaping

func (s seedMapList) flatten() seedMaping {
	flat := seedMaping{}
	for idx := range s {
		sort.Slice(s[idx], func(i, j int) bool {
			return s[idx][i].src < s[idx][j].src
		})
	}
	for _, soilRange := range s[0] {
		for i := 0; i < soilRange.width; {
			converted, skip := s.convertSkip(soilRange.src + i)
			newRange := seedRange{
				src:   soilRange.src + i,
				dest:  converted,
				width: skip,
			}
			flat = append(flat, newRange)
			i += skip
			if skip == 0 {
				i++
			}
		}
	}
	return flat
}

func (s seedMapList) convertSkip(seed int) (int, int) {
	maxSkip := utils.MaxInt
	var skip int
	for i := 0; i < len(s); i++ {
		seed, skip = s[i].convertSkip(seed)
		if skip < maxSkip {
			maxSkip = skip
		}
	}
	return seed, maxSkip
}

type seedMaping []seedRange

func (s seedMaping) convertSkip(seed int) (int, int) {
	for _, sr := range s {
		if seed >= sr.src && seed < (sr.src+sr.width) {
			skip := sr.width - (seed - sr.src)
			return seed - sr.src + sr.dest, skip
		}
	}
	for i := 0; i < len(s); i++ {
		if s[i].src > seed {
			return seed, s[i].src - seed - 1
		}
	}
	return seed, utils.MaxInt
}

type seedRange struct {
	src   int
	dest  int
	width int
}

var benchmark = false
