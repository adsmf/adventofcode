package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	w := loadCircuit("input.txt")
	s := w.run(signals{})
	return int(s["a"])
}

func part2() int {
	w := loadCircuit("input.txt")
	s := w.run(signals{})
	a := s["a"]
	w = loadCircuit("input.txt")
	s = w.run(signals{"b": a})
	return int(s["a"])
}

type signals map[string]uint16
type wiring map[string]string

func (w wiring) run(overrides signals) signals {
	s := signals{}
	for len(w) > 0 {
		foundSignals := []string{}
		for res, inp := range w {
			if val, found := overrides[res]; found {
				s[res] = val
				foundSignals = append(foundSignals, res)
				continue
			}
			if val, found := decodeParam(inp, s); found {
				s[res] = val
				foundSignals = append(foundSignals, res)
				continue
			}
			if strings.HasPrefix(inp, "NOT ") {
				notOf := inp[4:]
				if val, found := decodeParam(notOf, s); found {
					s[res] = -val - 1
					foundSignals = append(foundSignals, res)
					continue
				}
			}
			if strings.Contains(inp, " AND ") {
				sigs := strings.Split(inp, " AND ")
				valA, foundA := decodeParam(sigs[0], s)
				valB, foundB := decodeParam(sigs[1], s)
				if foundA && foundB {
					s[res] = valA & valB
					foundSignals = append(foundSignals, res)
					continue
				}
			}
			if strings.Contains(inp, " OR ") {
				sigs := strings.Split(inp, " OR ")
				valA, foundA := decodeParam(sigs[0], s)
				valB, foundB := decodeParam(sigs[1], s)
				if foundA && foundB {
					s[res] = valA | valB
					foundSignals = append(foundSignals, res)
					continue
				}
			}

			if strings.Contains(inp, " LSHIFT ") {
				sigs := strings.Split(inp, " LSHIFT ")
				valA, foundA := decodeParam(sigs[0], s)
				valB, foundB := decodeParam(sigs[1], s)
				if foundA && foundB {
					s[res] = valA << valB
					foundSignals = append(foundSignals, res)
					continue
				}
			}
			if strings.Contains(inp, " RSHIFT ") {
				sigs := strings.Split(inp, " RSHIFT ")
				valA, foundA := decodeParam(sigs[0], s)
				valB, foundB := decodeParam(sigs[1], s)
				if foundA && foundB {
					s[res] = valA >> valB
					foundSignals = append(foundSignals, res)
					continue
				}
			}
		}
		if len(foundSignals) == 0 {
			fmt.Printf("Couldn't find anything else\n%v\n", w)
			break
		}
		for _, f := range foundSignals {
			delete(w, f)
		}
	}
	return s
}

func decodeParam(inp string, s signals) (uint16, bool) {
	if sig, found := s[inp]; found {
		return sig, true
	}
	num, err := strconv.Atoi(inp)
	if err == nil {
		return uint16(num), true
	}
	return 0, false
}

func loadCircuit(filename string) wiring {
	w := wiring{}

	lines := utils.ReadInputLines(filename)
	for _, line := range lines {
		if !strings.Contains(line, " -> ") {
			continue
		}
		parts := strings.Split(line, " -> ")
		w[parts[1]] = parts[0]
	}

	return w
}
