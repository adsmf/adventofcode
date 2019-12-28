package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	return len(repeat("input.txt", 40))
}

func part2() int {
	return len(repeat("input.txt", 50))
}

func next(start string) string {
	new := ""
	var lastChar rune
	count := 0
	for _, char := range start + string('\x00') {
		if char != lastChar {
			if count > 0 {
				new += strconv.Itoa(count) + string(lastChar)
				count = 0
			}
			lastChar = char
		}
		count++
	}
	return new
}

func repeat(filename string, times int) string {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	result := string(input)
	result = strings.TrimSpace(result)
	for i := 0; i < times; i++ {
		fmt.Println(i, "\t", len(result))
		result = next(result)
	}
	return result
}
