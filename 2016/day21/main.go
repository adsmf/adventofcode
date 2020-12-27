package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %s\n", p1)
		fmt.Printf("Part 2: %s\n", p2)
	}
}

func part1() string {
	ops := load("input.txt")
	return process("abcdefgh", ops)
}

func part2() string {
	ops := load("input.txt")
	target := "fbgdceah"
	letters := make([]string, len(target))
	for i, let := range target {
		letters[i] = string(let)
	}
	perms := utils.PermuteStrings(letters)
	for _, perm := range perms {
		try := ""
		for _, let := range perm {
			try += let
		}
		result := process(try, ops)
		if result == target {
			return try
		}
	}
	return "UNSOLVED"
}

func process(seed string, operations []string) string {
	result := []byte(seed)
	for _, operation := range operations {
		parts := strings.Split(operation, " ")
		switch parts[0] {
		case "swap":
			if parts[1] == "position" {
				p1, p2 := utils.MustInt(parts[2]), utils.MustInt(parts[5])
				result[p1], result[p2] = result[p2], result[p1]
			} else {
				c1, c2 := parts[2][0], parts[5][0]
				for idx, c := range result {
					if c == c1 {
						result[idx] = c2
					} else if c == c2 {
						result[idx] = c1
					}
				}
			}
		case "reverse":
			start, end := utils.MustInt(parts[2]), utils.MustInt(parts[4])
			for i := 0; start+i < end-i; i++ {
				result[start+i], result[end-i] = result[end-i], result[start+i]
			}
		case "rotate":
			switch parts[1] {
			case "left":
				by := utils.MustInt(parts[2])
				result = append(result[by:], result[:by]...)
			case "right":
				by := len(result) - utils.MustInt(parts[2])
				result = append(result[by:], result[:by]...)
			case "based":
				by := bytes.IndexByte(result, parts[6][0])
				if by >= 4 {
					by++
				}
				by++
				by %= len(result)
				by = len(result) - by
				result = append(result[by:], result[:by]...)
			}
		case "move":
			from, to := utils.MustInt(parts[2]), utils.MustInt(parts[5])
			char := result[from]
			interim := make([]byte, len(result))
			copy(interim, result)
			interim = append(interim[:from], interim[from+1:]...)
			result = append(result[:from], result[from+1:]...)
			result = append(result[:to], char)
			result = append(result, interim[to:]...)
		}
	}
	return string(result)
}

func load(filename string) []string {
	lines := utils.ReadInputLines(filename)
	return lines
}

var benchmark = false
