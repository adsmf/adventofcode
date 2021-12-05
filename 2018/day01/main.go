package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	fileBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	parts := strings.Split(fileString, "\n")
	freq := 0
	seenFreqs := make(map[int]bool)
	p1 := 0
	first := true
	for {
		for _, mod := range parts {
			modInt, _ := strconv.Atoi(mod)
			freq = freq + modInt
			if seenFreqs[freq] {
				return p1, freq
			}
			seenFreqs[freq] = true
		}
		if first {
			first = false
			p1 = freq
		}
	}
}

var benchmark = false
