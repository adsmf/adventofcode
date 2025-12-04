package main

import (
	_ "embed"
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	// Part 1: 1411
	// Part 2: 8557
}

//go:embed input.txt
var inputCopy []byte

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for b.Loop() {
		copy(input, inputCopy)
		main()
	}
}

// func BenchmarkPart1(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		part1()
// 	}
// }

// func BenchmarkPart2(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		part2()
// 	}
// }
