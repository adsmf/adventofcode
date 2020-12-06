package main

import (
	"testing"
)

func TestAnswers(t *testing.T) {
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 117
	//Part 2: 909
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
