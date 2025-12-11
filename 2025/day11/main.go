package main

import (
	_ "embed"
	"fmt"
	"slices"

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
	mappings := map[idType][]idType{}
	nextID := idType(0)
	ids := make(map[string]idType, 1000)
	getID := func(node string) idType {
		if id, found := ids[node]; found {
			return id
		}
		id := nextID
		nextID++
		ids[node] = id
		return id
	}
	outNodes := make([]idType, 0, 30)
	utils.EachLine(input, func(index int, line string) (done bool) {
		in := line[0:3]
		outNodes = outNodes[0:0]
		for i := 5; i < len(line); i += 4 {
			out := line[i : i+3]
			outNodes = append(outNodes, getID(out))
		}
		mappings[getID(in)] = slices.Clone(outNodes)
		return
	})
	connections := make([][]idType, 750)
	connections = connections[:nextID]
	for from, to := range mappings {
		connections[from] = to
	}
	keyIDs := keyIDinfo{
		out: ids["out"],
		dac: ids["dac"],
		fft: ids["fft"],
	}
	p1 := countWays(make(map[search]int, 150), connections, keyIDs, ids["you"], false, false)
	p2 := countWays(make(map[search]int, 1500), connections, keyIDs, ids["svr"], true, true)
	return p1, p2
}

func countWays(cache map[search]int, mappings [][]idType, ids keyIDinfo, node idType, needDAC, needFFT bool) int {
	if cached, found := cache[search{node, needDAC, needFFT}]; found {
		return cached
	}
	if node == ids.out {
		if needDAC || needFFT {
			return 0
		}
		return 1
	}
	needDAC = needDAC && node != ids.dac
	needFFT = needFFT && node != ids.fft
	sub := 0
	for _, next := range mappings[node] {
		sub += countWays(cache, mappings, ids, next, needDAC, needFFT)
	}
	cache[search{node, needDAC, needFFT}] = sub
	return sub
}

type keyIDinfo struct {
	out idType
	dac idType
	fft idType
}

type search struct {
	node    idType
	needDAC bool
	needFFT bool
}

type idType uint16

var benchmark = false
