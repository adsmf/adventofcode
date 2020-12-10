package main

import (
	"io/ioutil"
	"testing"

	"github.com/adsmf/adventofcode/utils"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 2590
	//Part 2: 226775649501184
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkPart1(b *testing.B) {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	adapters := utils.GetInts(string(inputBytes))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(adapters)
	}
}

func BenchmarkPart2(b *testing.B) {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	adapters := utils.GetInts(string(inputBytes))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(adapters)
	}
}
