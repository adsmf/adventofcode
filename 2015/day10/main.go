package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := play()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func play() (int, int) {
	p1Str := repeat(input, 40)
	p2Str := repeat(p1Str, 10)
	return len(p1Str), len(p2Str)
}

func next(start string) string {
	new := strings.Builder{}
	var lastChar rune
	count := 0
	for _, char := range start + string('\x00') {
		if char != lastChar {
			if count > 0 {
				new.WriteString(strconv.Itoa(count))
				new.WriteByte(byte(lastChar))
				count = 0
			}
			lastChar = char
		}
		count++
	}
	return new.String()
}

func repeat(in string, times int) string {
	result := strings.TrimSpace(in)
	for i := 0; i < times; i++ {
		result = next(result)
	}
	return result
}

var benchmark = false
