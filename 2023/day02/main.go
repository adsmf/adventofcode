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
	games := loadGames()
	p1 := part1(games)
	p2 := part2(games)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(games gameList) int {
	sumID := 0
	for _, game := range games {
		if game.valid() {
			sumID += game.id
		}
	}
	return sumID
}

func part2(games gameList) int {
	totalPower := 0
	for _, game := range games {
		totalPower += game.power()
	}
	return totalPower
}

func loadGames() gameList {
	games := gameList{}
	utils.EachLine(input, func(line string) (done bool) {
		game := gameInfo{}

		split := strings.Split(line, ": ")
		game.id = utils.GetInts(split[0])[0]
		for _, roundStr := range strings.Split(split[1], "; ") {
			round := roundInfo{}
			var color string
			var count int
			for _, pull := range strings.Split(roundStr, ", ") {
				_, err := fmt.Sscanf(pull, "%d %s", &count, &color)
				if err != nil {
					panic(err)
				}
				round[color] = count
			}
			game.rounds = append(game.rounds, round)
		}
		games = append(games, game)
		return false
	})
	return games
}

type gameList []gameInfo

type gameInfo struct {
	id     int
	rounds []roundInfo
}

func (g gameInfo) valid() bool {
	for _, round := range g.rounds {
		if round["red"] > 12 || round["green"] > 13 || round["blue"] > 14 {
			return false
		}
	}
	return true
}

func (g gameInfo) power() int {
	minCubes := map[string]int{}
	for _, round := range g.rounds {
		for color, count := range round {
			if minCubes[color] < count {
				minCubes[color] = count
			}
		}
	}
	power := 1
	for _, count := range minCubes {
		power *= count
	}
	return power
}

type roundInfo map[string]int

var benchmark = false
