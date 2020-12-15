package main

import (
	"io/ioutil"
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 866
	//Part 2: 1437692
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkPart1(b *testing.B) {
	input, _ := ioutil.ReadFile("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		play(string(input), 2020)
	}
}

func BenchmarkPart2(b *testing.B) {
	input, _ := ioutil.ReadFile("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		play(string(input), 30000000)
	}
}
