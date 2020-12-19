package main

import (
	"fmt"
	"io/ioutil"

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
	inputBytes, _ := ioutil.ReadFile("input.txt")
	jumps := utils.GetInts(string(inputBytes))
	step := 0
	ip := 0
	for {
		step++
		jumps[ip], ip = jumps[ip]+1, ip+jumps[ip]
		if ip < 0 || ip >= len(jumps) {
			return step
		}
	}
}

func part2() int {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	jumps := utils.GetInts(string(inputBytes))
	step := 0
	ip, oldIP := 0, 0
	for {
		step++
		oldIP, ip, jumps[ip] = ip, ip+jumps[ip], jumps[ip]+1
		if ip < 0 || ip >= len(jumps) {
			return step
		}
		if jumps[oldIP] == 4 {
			jumps[oldIP] = 2
		}
	}
}

var benchmark = false
