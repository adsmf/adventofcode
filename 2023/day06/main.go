package main

import (
	_ "embed"
	"fmt"
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
	races, bigRace := getRaces()
	multWays := 1
	for _, race := range races {
		multWays *= calcRace(race)
	}
	bigWays := calcRace(bigRace)
	return multWays, bigWays
}

func calcRace(race raceInfo) int {
	for i := 1; i <= race.time; i++ {
		remTime := race.time - i
		dist := remTime * i
		if dist > race.distance {
			lastWin := race.time - i
			return lastWin - i + 1
		}
	}
	return -1
}

func getRaces() ([]raceInfo, raceInfo) {
	smallRaces := []raceInfo{{}}
	bigRace := raceInfo{}
	lineNum, curIndex := 0, 0
	hasDigit := false
	for _, ch := range input {
		switch {
		case ch == '\n':
			lineNum++
			curIndex = 0
			hasDigit = false
		case ch == ' ' && hasDigit:
			hasDigit = false
			curIndex++
			if lineNum == 0 {
				smallRaces = append(smallRaces, raceInfo{})
			}
		case ch >= '0' && ch <= '9':
			hasDigit = true
			smallRace := &smallRaces[curIndex]
			if lineNum == 0 {
				bigRace.time = 10*bigRace.time + int(ch-'0')
				smallRace.time = 10*smallRace.time + int(ch-'0')
			} else {
				bigRace.distance = 10*bigRace.distance + int(ch-'0')
				smallRace.distance = 10*smallRace.distance + int(ch-'0')
			}
		}
	}
	return smallRaces, bigRace
}

type raceInfo struct {
	time     int
	distance int
}

var benchmark = false
