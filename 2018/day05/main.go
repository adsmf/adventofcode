package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	polymerBytes, _ := ioutil.ReadFile("input.txt")
	polymer := string(polymerBytes)
	p1, reacted := part1(polymer)
	p2 := part2(reacted)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}
func part1(polymer string) (int, string) {
	polymer = fullyReact(polymer)
	return len(polymer), polymer
}

func part2(polymer string) int {
	bestLength := len(polymer)
	for char := 'a'; char <= 'z'; char++ {
		re, _ := regexp.Compile(fmt.Sprintf("(?i)%c", char))
		testPoly := string(re.ReplaceAll([]byte(polymer), []byte{}))
		testPoly = fullyReact(testPoly)
		if len(testPoly) < bestLength {
			bestLength = len(testPoly)
		}
	}
	return bestLength
}

func fullyReact(polymer string) string {
	for {
		newPolymer, reacted := reactOnce(polymer)
		if !reacted {
			break
		}
		polymer = newPolymer
	}
	return polymer
}

func reactOnce(polymer string) (string, bool) {
	reacted := false
	caseDiff := byte('A' ^ 'a')
	newPolymer := strings.Builder{}

	for i := 0; i < len(polymer)-1; i++ {
		current := byte(polymer[i])
		next := byte(polymer[i+1])
		if current^next == caseDiff {
			i++
			reacted = true
			continue
		}
		newPolymer.WriteByte(current)
		if i == len(polymer)-2 {
			newPolymer.WriteByte(next)
		}
	}
	return newPolymer.String(), reacted
}

var benchmark = false
