package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	passphrases := utils.ReadInputLines("input.txt")
	p1 := part1(passphrases)
	p2 := part2(passphrases)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(passphrases []string) int {
	count := 0
	for _, passphrase := range passphrases {
		words := map[string]int{}
		valid := true
		for _, word := range strings.Split(passphrase, " ") {
			if _, found := words[word]; found {
				valid = false
				break
			}
			words[word]++
		}
		if valid {
			count++
		}
	}
	return count
}

func part2(passphrases []string) int {
	count := 0
	for _, passphrase := range passphrases {
		words := map[string]int{}
		valid := true
		for _, word := range strings.Split(passphrase, " ") {
			letters := []byte(word)
			sort.Slice(letters, func(i, j int) bool {
				return letters[i] < letters[j]
			})
			sortedWord := string(letters)
			if _, found := words[sortedWord]; found {
				valid = false
				break
			}
			words[sortedWord]++
		}
		if valid {
			count++
		}
	}
	return count
}

var benchmark = false
