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
	sumID := 0
	totalPower := 0

	utils.EachLine(input, func(index int, line string) (done bool) {
		validGame := true
		maxPulled := roundInfo{}
		start := strings.Index(line, ": ") + 1
		utils.EachSection(line[start:], ';', func(index int, roundStr string) (done bool) {
			round := roundInfo{}
			utils.EachSection(roundStr, ',', func(index int, pull string) (done bool) {
				colour, count := parsePull(pull)
				round.set(colour, count)
				return false
			})
			if !round.valid() {
				validGame = false
			}
			maxPulled = maxPulled.extend(round)
			return false
		})
		if validGame {
			sumID += index + 1
		}
		totalPower += maxPulled.power()
		return false
	})
	return sumID, totalPower
}

func parsePull(pull string) (string, int) {
	count := 0
	var colour string
	for pos := 0; pos < len(pull); pos++ {
		ch := pull[pos]
		if ch == ' ' && count > 0 {
			colour = pull[pos+1:]
			break
		}
		if ch >= '0' && ch <= '9' {
			count *= 10
			count += int(ch - '0')
		}
	}
	return colour, count
}

type roundInfo struct {
	red   int
	green int
	blue  int
}

func (r *roundInfo) set(colour string, count int) {
	switch colour {
	case "red":
		r.red = count
	case "green":
		r.green = count
	case "blue":
		r.blue = count
	}
}

func (r roundInfo) power() int  { return r.red * r.green * r.blue }
func (r roundInfo) valid() bool { return !(r.red > 12 || r.green > 13 || r.blue > 14) }

func (r roundInfo) extend(other roundInfo) roundInfo {
	if other.red > r.red {
		r.red = other.red
	}
	if other.green > r.green {
		r.green = other.green
	}
	if other.blue > r.blue {
		r.blue = other.blue
	}
	return r
}

var benchmark = false
