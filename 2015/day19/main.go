package main

import (
	"fmt"
	"github.com/adsmf/adventofcode/utils"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	l := loadInput("input.txt")
	return l.countOpts()
}

func part2() int {
	return -1
}

type lab struct {
	replacements map[string][]string
	initial      string
}

func (l lab) countOpts() int {
	opts := map[string]struct{}{}
	for i := 0; i < len(l.initial); i++ {
		pre := l.initial[:i]
		sub := l.initial[i:]
		for rep, withOpts := range l.replacements {
			if strings.HasPrefix(sub, rep) {
				suf := strings.TrimPrefix(sub, rep)
				for _, with := range withOpts {
					opts[pre+with+suf] = struct{}{}
				}
			}
		}
	}
	return len(opts)
}

func loadInput(filename string) lab {
	l := lab{
		replacements: map[string][]string{},
	}
	for _, line := range utils.ReadInputLines(filename) {
		if line == "" {
			continue
		}
		if strings.Contains(line, " => ") {
			parts := strings.Split(line, " => ")
			rep := parts[0]
			with := parts[1]
			if l.replacements[rep] == nil {
				l.replacements[rep] = []string{}
			}
			l.replacements[rep] = append(l.replacements[rep], with)
			continue
		}
		l.initial = line
	}
	return l
}
