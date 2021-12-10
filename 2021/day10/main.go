package main

import (
	_ "embed"
	"fmt"
	"sort"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := solve(input)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve(in string) (int, int) {
	lines := utils.GetLines(in)
	totalErrorScore := 0
	compScores := make([]int, 0, len(lines))
	for _, line := range lines {
		errScore, compScore := parseLine(line)
		totalErrorScore += errScore
		if compScore > 0 {
			compScores = append(compScores, compScore)
		}
	}
	sort.Ints(compScores)
	middleCompleteScore := compScores[len(compScores)/2]
	return totalErrorScore, middleCompleteScore
}

var syntaxErrorValues = map[byte]int{')': 3, ']': 57, '}': 1197, '>': 25137}
var closePair = map[byte]byte{'(': ')', '[': ']', '{': '}', '<': '>'}
var autocompleteValues = map[byte]int{'(': 1, '[': 2, '{': 3, '<': 4}

func parseLine(line string) (int, int) {
	unclosed := []byte{}
	for _, ch := range []byte(line) {
		switch ch {
		case '(', '[', '{', '<':
			unclosed = append(unclosed, ch)
		case ')', ']', '}', '>':
			last := unclosed[len(unclosed)-1]
			if ch != closePair[last] {
				return syntaxErrorValues[ch], 0
			}
			unclosed = unclosed[:len(unclosed)-1]
		}
	}
	completeScore := closeScore(unclosed)
	return 0, completeScore
}

func closeScore(unclosed []byte) int {
	score := 0
	for i := len(unclosed) - 1; i >= 0; i-- {
		ch := unclosed[i]
		score *= 5
		score += autocompleteValues[ch]
	}
	return score
}

var benchmark = false
