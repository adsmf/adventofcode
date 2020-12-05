package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 832
	//Part 2: 517
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkBitwise(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loadBitwise("input.txt")
	}
}
