package main

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	return santaCoin("input.txt", 5)
}

func part2() int {
	return santaCoin("input.txt", 6)
}

func santaCoin(filename string, numZeros int) int {
	input := utils.ReadInputLines(filename)[0]
	want := strings.Repeat("0", numZeros)
	for i := 0; ; i++ {
		h := md5.New()
		_, err := h.Write([]byte(fmt.Sprintf("%s%d", input, i)))
		if err != nil {
			panic(err)
		}
		hashString := fmt.Sprintf("%x", h.Sum(nil))
		if hashString[0:numZeros] == want {
			return i
		}
	}
}

var benchmark = false
