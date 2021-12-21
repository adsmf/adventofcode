package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1 := part1(input)
	p2 := part2(input)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(input string) int {
	vals := utils.GetInts(input)
	pos1, pos2 := vals[1], vals[3]
	score1, score2 := 0, 0
	dieVal := 1
	rolled := 0
	for score1 < 1000 && score2 < 1000 {
		rolled += 3
		move := 0
		for i := 0; i < 3; i++ {
			move += dieVal
			dieVal++
			if dieVal > 100 {
				dieVal -= 100
			}
		}
		pos1 += move
		for pos1 > 10 {
			pos1 -= 10
		}
		score1 += pos1
		pos1, pos2 = pos2, pos1
		score1, score2 = score2, score1
	}
	return score1 * rolled
}

func part2(input string) int {
	cache := make(map[stateInfo]scorePair, 20000)
	vals := utils.GetInts(input)
	pos1, pos2 := vals[1], vals[3]
	combs := map[int]int{}
	for d1 := 1; d1 <= 3; d1++ {
		for d2 := 1; d2 <= 3; d2++ {
			for d3 := 1; d3 <= 3; d3++ {
				total := d1 + d2 + d3
				combs[total]++
			}
		}
	}
	wins, _ := quantum(cache, combs, pos1, pos2, 0, 0, true)
	return wins
}

func quantum(cache map[stateInfo]scorePair, combs map[int]int, pos1, pos2 int, score1, score2 int, p1plays bool) (int, int) {
	if score1 >= 21 {
		return 1, 0
	}
	if score2 >= 21 {
		return 0, 1
	}
	if scores, found := cache[stateInfo{pos1, pos2, score1, score2, p1plays}]; found {
		return scores.player1, scores.player2
	}
	if scores, found := cache[stateInfo{pos2, pos1, score2, score1, !p1plays}]; found {
		return scores.player2, scores.player1
	}
	scores := scorePair{0, 0}
	for move, times := range combs {
		if p1plays {
			nextPos := pos1 + move
			if nextPos > 10 {
				nextPos -= 10
			}
			s1, s2 := quantum(cache, combs, nextPos, pos2, score1+nextPos, score2, false)
			scores.player1 += s1 * times
			scores.player2 += s2 * times
		} else {
			nextPos := pos2 + move
			if nextPos > 10 {
				nextPos -= 10
			}
			s1, s2 := quantum(cache, combs, pos1, nextPos, score1, score2+nextPos, true)
			scores.player1 += s1 * times
			scores.player2 += s2 * times
		}
	}
	cache[stateInfo{pos1, pos2, score1, score2, p1plays}] = scores
	return scores.player1, scores.player2
}

type stateInfo struct {
	pos1, pos2     int
	score1, score2 int
	p1next         bool
}
type scorePair struct {
	player1, player2 int
}

var benchmark = false
