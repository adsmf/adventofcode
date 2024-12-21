package main

import (
	_ "embed"
	"fmt"
	"math"

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

type routeSet ['v' + 1]['v' + 1][]string
type cacheSet map[cacheKey]int

func solve() (int, int) {
	routes := routeSet{}
	calcRoutes(&routes, numPadButtons)
	calcRoutes(&routes, dPadButtons)
	p1, p2 := 0, 0
	cache := cacheSet{}
	utils.EachLine(input, func(index int, line string) (done bool) {
		p1len := minTyped(line, routes, cache, 3)
		p2len := minTyped(line, routes, cache, 26)
		num := utils.GetInts(line)
		complexP1 := p1len * num[0]
		complexP2 := p2len * num[0]
		p1 += complexP1
		p2 += complexP2
		return false
	})
	return p1, p2
}

func minTyped(input string, routes routeSet, cache cacheSet, depth int) int {
	key := cacheKey{input, depth}
	if res, found := cache[key]; found {
		return res
	}
	route := 0
	if depth == 0 {
		route = len(input)
	} else {
		cur := 'A'
		for _, ch := range input {
			route += countMoves(byte(cur), byte(ch), depth, routes, cache)
			cur = ch
		}
	}
	cache[key] = route
	return route
}

func countMoves(start, end byte, depth int, routes routeSet, cache cacheSet) int {
	if start == end {
		return 1
	}
	bestSeqLen := math.MaxInt
	for _, seq := range routes[start][end] {
		seqLen := minTyped(seq, routes, cache, depth-1)
		bestSeqLen = min(bestSeqLen, seqLen)
	}
	return bestSeqLen
}

type cacheKey struct {
	input string
	depth int
}

func calcRoutes(routes *routeSet, buttons map[byte]point) {
	for start := range buttons {
		for end := range buttons {
			all := allMoves(buttons[start], buttons[end], buttons)
			routes[start][end] = all
		}
	}
}

type search struct {
	pos  point
	keys string
}

func allMoves(start, target point, buttons map[byte]point) []string {
	validPoints := map[point]bool{}
	for _, pos := range buttons {
		validPoints[pos] = true
	}
	open, next := []search{{start, ""}}, []search{}

	moves := []string{}
	for len(open) > 0 {
		for _, cur := range open {
			if cur.pos == target {
				cur.keys += "A"
				moves = append(moves, cur.keys)
				continue
			}
			if !validPoints[cur.pos] {
				continue
			}
			if cur.pos.y < target.y {
				next = append(next, search{cur.pos.add(point{0, 1}), cur.keys + "^"})
			}
			if cur.pos.y > target.y {
				next = append(next, search{cur.pos.add(point{0, -1}), cur.keys + "v"})
			}
			if cur.pos.x < target.x {
				next = append(next, search{cur.pos.add(point{1, 0}), cur.keys + ">"})
			}
			if cur.pos.x > target.x {
				next = append(next, search{cur.pos.add(point{-1, 0}), cur.keys + "<"})
			}
		}
		open, next = next, open[0:0]
	}
	return moves
}

/*
+---+---+---+
| 7 | 8 | 9 |
+---+---+---+
| 4 | 5 | 6 |
+---+---+---+
| 1 | 2 | 3 |
+---+---+---+
|   | 0 | A |
+---+---+---+
*/
var numPadButtons = map[byte]point{
	'9': {2, 3}, '8': {1, 3}, '7': {0, 3},
	'6': {2, 2}, '5': {1, 2}, '4': {0, 2},
	'3': {2, 1}, '2': {1, 1}, '1': {0, 1},
	'A': {2, 0}, '0': {1, 0},
}

/*
+---+---+---+
|   | ^ | A |
+---+---+---+
| < | v | > |
+---+---+---+
*/
var dPadButtons = map[byte]point{
	'A': {2, 1}, '^': {1, 1},
	'>': {2, 0}, 'v': {1, 0}, '<': {0, 0},
}

type point struct{ x, y int }

func (p point) add(o point) point { return point{p.x + o.x, p.y + o.y} }

var benchmark = false
