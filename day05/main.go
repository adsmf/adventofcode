package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	polymerBytes, _ := ioutil.ReadFile("input.txt")
	polymer := string(polymerBytes)
	polymer = fullyReact(polymer)
	fmt.Printf("Polymer length: %d\n", len(polymer))
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
