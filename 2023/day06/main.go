package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	races := getRaces()
	multWays := 1
	for _, race := range races {
		multWays *= calcRace(race)
	}
	return multWays
}

func part2() int {
	return calcRace(getBigRace())
}

func calcRace(race raceInfo) int {
	winWays := 0
	for i := 1; i < race.time; i++ {
		remTime := race.time - i
		dist := remTime * i
		if dist > race.distance {
			winWays++
		}
	}
	return winWays
}

func getRaces() []raceInfo {
	races := []raceInfo{}

	utils.EachLine(input, func(index int, line string) (done bool) {
		switch index {
		case 0:
			utils.EachInteger(line, func(index, time int) (done bool) {
				races = append(races, raceInfo{time: time})
				return false
			})
		case 1:
			utils.EachInteger(line, func(index, distance int) (done bool) {
				races[index].distance = distance
				return false
			})
		}
		return false
	})
	return races
}

func getBigRace() raceInfo {
	lineNum := 0
	time := 0
	distance := 0
	for _, ch := range input {
		switch {
		case ch == '\n':
			lineNum++
		case ch >= '0' && ch <= '9':
			if lineNum == 0 {
				time = 10*time + int(ch-'0')
			} else {
				distance = 10*distance + int(ch-'0')
			}
		}
	}
	return raceInfo{
		time:     time,
		distance: distance,
	}
}

type raceInfo struct {
	time     int
	distance int
}

var benchmark = false
