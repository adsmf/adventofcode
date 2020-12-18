package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	input := string(inputBytes)[:len(inputBytes)-1]
	p1, p2 := sum(input)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func sum(input string) (int, int) {
	s1, s2 := 0, 0
	o1, o2, mod := 1, len(input)/2, len(input)
	for pos, char := range input {
		if byte(char) == input[(pos+o1)%mod] {
			s1 += int(char - '0')
		}
		if byte(char) == input[(pos+o2)%mod] {
			s2 += int(char - '0')
		}
	}

	return s1, s2
}

var benchmark = false
