package main

import (
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1 := part1()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
	}
}

func part1() int {
	stars := []star{}
	for _, line := range utils.ReadInputLines("input.txt") {
		pos := utils.GetInts(line)
		stars = append(stars, star{pos[0], pos[1], pos[2], pos[3]})
	}
	constellations := map[star]map[star]bool{}

	for _, cur := range stars {
		matches := []star{}
		for root, constStars := range constellations {
			for constStar := range constStars {
				if cur.manhattan(constStar) <= 3 {
					matches = append(matches, root)
					constellations[root][cur] = true
					break
				}
			}
		}
		if len(matches) == 0 {
			constellations[cur] = map[star]bool{cur: true}
		}
		if len(matches) > 1 {
			first, others := matches[0], matches[1:]
			for _, oth := range others {
				for othStar := range constellations[oth] {
					constellations[first][othStar] = true
				}
				delete(constellations, oth)
			}
		}
	}
	return len(constellations)
}

type star struct{ x, y, z, t int }

func (s star) manhattan(to star) int {
	return dist(s.x, to.x) +
		dist(s.y, to.y) +
		dist(s.z, to.z) +
		dist(s.t, to.t)
}

func dist(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

var benchmark = false
