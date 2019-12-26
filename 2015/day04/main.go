package main

import (
	"crypto/md5"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	return santaCoin("input.txt")
}

func part2() int {
	return 0
}

func santaCoin(filename string) int {
	input := utils.ReadInputLines(filename)[0]
	for i := 0; i < 1000000; i++ {
		h := md5.New()
		h.Write([]byte(fmt.Sprintf("%s%d", input, i)))
		hashString := fmt.Sprintf("%x", h.Sum(nil))
		if hashString[0:5] == "00000" {
			return i
		}
	}
	return -1
}
