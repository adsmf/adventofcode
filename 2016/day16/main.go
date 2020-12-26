package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	base := load("input.txt")
	p1 := calc(base, 272)
	p2 := calc(base, 35651584)
	if !benchmark {
		fmt.Printf("Part 1: %s\n", p1)
		fmt.Printf("Part 2: %s\n", p2)
	}
}

func calc(data string, diskSize int) string {
	for len(data) < diskSize {
		data = expand(data)
	}
	return checksum(data[:diskSize])
}

func expand(input string) string {
	expanded := make([]byte, len(input)*2+1)
	copy(expanded, []byte(input))
	expanded[len(input)] = '0'
	for offset := 0; offset < len(input); offset++ {
		if input[len(input)-offset-1] == '0' {
			expanded[len(input)+offset+1] = '1'
		} else {
			expanded[len(input)+offset+1] = '0'
		}
	}
	return string(expanded)
}

func checksum(input string) string {
	if len(input)%2 == 1 {
		return input
	}
	reduced := make([]byte, len(input)/2)
	for pairIndex := 0; pairIndex < len(input)/2; pairIndex++ {
		pair := input[pairIndex*2 : pairIndex*2+2]
		if pair[0] == pair[1] {
			reduced[pairIndex] = '1'
		} else {
			reduced[pairIndex] = '0'
		}
	}
	return checksum(string(reduced))
}

func load(filename string) string {
	raw, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(raw))
}

var benchmark = false
