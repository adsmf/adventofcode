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
	p1, p2 := analyse()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func analyse() (int, int) {
	var haveSeeds []int
	currentMap := -1
	mappings := [7]seedMap{}
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
	convertSkip := func(seed int) (int, int) {
		maxSkip := utils.MaxInt
		var skip int
		for i := 0; i < len(mappings); i++ {
			seed, skip = mappings[i].convertSkip(seed)
			if skip < maxSkip {
				maxSkip = skip
			}
		}
		return seed, maxSkip
	}
	// part 1
	lowestP1 := -1
	for _, seed := range haveSeeds {
		for i := 0; i < len(mappings); i++ {
			seed = mappings[i].convert(seed)
		}
		if lowestP1 == -1 || seed < lowestP1 {
			lowestP1 = seed
		}
	}
	// part 2
	lowestP2 := -1
	for i := 0; i < len(haveSeeds); i += 2 {
		for j := 0; j < haveSeeds[i+1]; j++ {
			skip := 0
			seed := haveSeeds[i] + j
			seed, skip = convertSkip(seed)
			if lowestP2 == -1 || seed < lowestP2 {
				lowestP2 = seed
			}
			if skip > 0 {
				j += skip - 1
			}
		}
	}
	return lowestP1, lowestP2
}

type seedMap []seedRange

func (s seedMap) convert(seed int) int {
	for _, sr := range s {
		if seed >= sr.src && seed < (sr.src+sr.width) {
			return seed - sr.src + sr.dest
		}
	}
	return seed
}

func (s seedMap) convertSkip(seed int) (int, int) {
	for _, sr := range s {
		if seed >= sr.src && seed < (sr.src+sr.width) {
			skip := sr.width - (seed - sr.src)
			return seed - sr.src + sr.dest, skip
		}
	}
	return seed, 0 // TODO: not zero
}

type seedRange struct {
	src   int
	dest  int
	width int
}

var benchmark = false
