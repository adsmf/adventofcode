package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 113220
	//Part 2: 30599555965
}

func TestAnswers(t *testing.T) {
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
