package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 5108096
	//Part 2: 10553942650264
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loadInput()
	}
}
func BenchmarkP1(b *testing.B) {
	g := loadInput()
	for i := 0; i < b.N; i++ {
		part1(g)
	}
}

func BenchmarkP2(b *testing.B) {
	g := loadInput()
	for i := 0; i < b.N; i++ {
		part2(g)
	}
}
