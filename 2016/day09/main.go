package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

var benchmark = false

func main() {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	input := strings.TrimSpace(string(inputBytes))
	p1 := calculateLength(input, false)
	p2 := calculateLength(input, true)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

var matcher = regexp.MustCompile(`^([^(]*)\(([^)]+)\)(.*)$`)

func calculateLength(input string, expandData bool) int {
	matches := matcher.FindStringSubmatch(input)
	if len(matches) == 4 {
		prefix, remainder := matches[1], matches[3]
		expand := utils.GetInts(matches[2])
		expandSection, remainder := remainder[:expand[0]], string(remainder[expand[0]:])
		expansionLen := len(expandSection)
		if expandData {
			expansionLen = calculateLength(expandSection, true)
		}
		return len(prefix) + expansionLen*expand[1] + calculateLength(remainder, expandData)
	}
	return len(input)
}
