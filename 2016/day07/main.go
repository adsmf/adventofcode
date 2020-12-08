package main

import (
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

var benchmark = false

func main() {
	addresses := load()
	p1 := part1(addresses)
	p2 := part2(addresses)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(addresses []address) int {
	count := 0
	for _, addr := range addresses {
		if addr.valid() {
			count++
		}
	}
	return count
}

func part2(addresses []address) int {
	count := 0
	for _, addr := range addresses {
		if addr.validSSL() {
			count++
		}
	}
	return count
}

func load() []address {
	addresses := []address{}
	for _, line := range utils.ReadInputLines("input.txt") {
		buffer := []rune{}
		parts := []string{}
		for _, char := range line + "§" {
			switch char {
			case '[', ']', '§':
				parts = append(parts, string(buffer))
				buffer = []rune{}
			default:
				buffer = append(buffer, char)
			}
		}
		addresses = append(addresses, address{parts: parts})
	}
	return addresses
}

type address struct {
	parts []string
}

func (a address) validSSL() bool {
	abas := map[string]bool{}
	babs := map[string]bool{}
	for i := 0; i < len(a.parts); i++ {
		part := a.parts[i]
		for j := 0; j <= len(part)-3; j++ {
			if part[j+0] != part[j+2] || part[j+0] == part[j+1] {
				continue
			}
			if i%2 == 0 {
				abas[part[j:j+3]] = true
			} else {
				babs[part[j:j+3]] = true
			}
		}
	}

	for aba := range abas {
		bab := fmt.Sprintf("%c%c%c", aba[1], aba[0], aba[1])
		if _, found := babs[bab]; found {
			return true
		}
	}

	return false
}

func (a address) valid() bool {
	for i := 1; i < len(a.parts); i += 2 {
		part := a.parts[i]
		for j := 0; j < len(part)-3; j++ {
			if part[j+0] == part[j+3] &&
				part[j+1] == part[j+2] &&
				part[j+0] != part[j+1] {
				return false
			}
		}
	}
	for i := 0; i < len(a.parts); i += 2 {
		part := a.parts[i]
		for j := 0; j < len(part)-3; j++ {
			if part[j+0] == part[j+3] &&
				part[j+1] == part[j+2] &&
				part[j+0] != part[j+1] {
				return true
			}
		}
	}
	return false
}
