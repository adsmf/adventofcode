package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	listing := parseListing()
	p1 := part1(listing)
	p2 := part2(listing)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2:\n%s\n", p2)
	}
}

func part1(listing programListing) int {
	regX := 1
	signal := 0
	cycle := 0
	countCycle := 20
	for i := 0; i < len(listing); i++ {
		cycles := 1
		if listing[i].op == opAddx {
			cycles = 2
		}
		for j := 0; j < cycles; j++ {
			cycle++
			if cycle == countCycle {
				signal += (cycle) * regX
				countCycle += 40
			}
		}
		regX += listing[i].value
	}
	return signal
}

func part2(listing programListing) string {
	regX := 1
	cycle := 0
	sb := strings.Builder{}
	for i := 0; i < len(listing); i++ {
		cycles := 1
		if listing[i].op == opAddx {
			cycles = 2
		}
		for j := 0; j < cycles; j++ {
			cycle++
			hPos := cycle % 40
			if regX >= hPos-2 && regX <= hPos {
				sb.WriteString("⬜️")
			} else {
				sb.WriteString("⬛️")
			}
			if hPos == 0 {
				sb.WriteByte('\n')
			}
		}
		regX += listing[i].value
	}
	return sb.String()
}

type programListing []instruction

type instruction struct {
	op    opCode
	value int
}

type opCode int

const (
	opNoop = iota
	opAddx
)

func parseListing() programListing {
	listing := programListing{}
	for _, line := range utils.GetLines(input) {
		words := strings.Split(line, " ")
		var inst instruction
		switch words[0] {
		case "noop":
			inst = instruction{op: opNoop, value: 0}
		case "addx":
			val, _ := strconv.Atoi(words[1])
			inst = instruction{
				op:    opAddx,
				value: val,
			}
		}
		listing = append(listing, inst)
	}
	return listing
}

var benchmark = false
