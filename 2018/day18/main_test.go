package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 675100
	//Part 2: 191820
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lumberLife(10)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lumberLife(1000000000)
	}
}
