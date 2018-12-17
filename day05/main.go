package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	polymerBytes, _ := ioutil.ReadFile("input.txt")
	polymer := string(polymerBytes)
	// part1(polymer)
	part2(polymer)
}
func part1(polymer string) {
	polymer = fullyReact(polymer)
	fmt.Printf("Polymer length: %d\n", len(polymer))
}

func part2(polymer string) {
	bestLength := len(polymer)
	for char := 'a'; char <= 'z'; char++ {
		re, _ := regexp.Compile(fmt.Sprintf("(?i)%c", char))
		testPoly := string(re.ReplaceAll([]byte(polymer), []byte{}))
		testPoly = fullyReact(testPoly)
		if len(testPoly) < bestLength {
			bestLength = len(testPoly)
			fmt.Printf("Better alternative: %c => %d\n", char, len(testPoly))
		}

	}
}

func fullyReact(polymer string) string {
	for {
		newPolymer, reacted := reactOnce(polymer)
		if !!!reacted {
			break
		}
		polymer = newPolymer
	}
	return polymer
}

func reactOnce(polymer string) (string, bool) {
	reacted := false
	caseDiff := byte('A' ^ 'a')
	newPolymer := ""

	for i := 0; i < len(polymer)-1; i++ {
		current := byte(polymer[i])
		next := byte(polymer[i+1])
		if current^next == caseDiff {
			i++
			reacted = true
			continue
		}
		newPolymer = fmt.Sprintf("%s%c", newPolymer, current)
		if i == len(polymer)-2 {
			newPolymer = fmt.Sprintf("%s%c", newPolymer, next)
		}
	}

	return newPolymer, reacted
}
