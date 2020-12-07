package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

var benchmark = false

func main() {
	rooms := load("input.txt")
	p1 := part1(rooms)
	p2 := part2(rooms)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(rooms []room) int {
	count := 0
	for _, room := range rooms {
		if room.validate() {
			count += room.sector
		}
	}
	return count
}

func part2(rooms []room) int {
	for _, room := range rooms {
		if room.validate() {
			if room.decrypt() == "northpole object storage" {
				return room.sector
			}
		}
	}
	return -1
}

type room struct {
	enc      string
	sector   int
	checksum string
}

func (r room) decrypt() string {
	return strings.Map(func(c rune) rune {
		switch {
		case c == '-':
			return ' '
		case c >= 'a' && c <= 'z':
			return rune((int(c)+r.sector-'a')%26 + 'a')
		}
		return c
	}, r.enc)
}

func (r room) validate() bool {
	counts := map[rune]int{}
	for _, char := range r.enc {
		counts[char]++
	}

	chars := []rune{}
	for char := range counts {
		if char >= 'a' && char <= 'z' {
			chars = append(chars, char)
		}
	}
	sort.Slice(chars, func(i int, j int) bool {
		ci := chars[i]
		cj := chars[j]
		if counts[ci] == counts[cj] {
			return ci < cj
		}
		return counts[ci] >= counts[cj]
	})
	return r.checksum == string(chars[0:5])
}

func load(filename string) []room {
	rooms := []room{}
	for _, line := range utils.ReadInputLines(filename) {
		re := regexp.MustCompile(`^([a-z-]+)-(\d+)\[(.{5})\]$`)
		parts := re.FindStringSubmatch(line)
		sector, _ := strconv.Atoi(parts[2])
		newRoom := room{parts[1], sector, parts[3]}
		rooms = append(rooms, newRoom)
	}
	return rooms
}
