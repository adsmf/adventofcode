package main

import (
	"fmt"
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 454
	//Part 2: 8,51
}

func TestAnswers(t *testing.T) {
}

func TestPart1Examples(t *testing.T) {
	type testDef struct {
		// Test structure here
	}
	tests := []testDef{
		// Test data here
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			// Assertions here
		})
	}
}

func TestPart2Examples(t *testing.T) {
	type testDef struct {
		// Test structure here
	}
	tests := []testDef{
		// Test data here
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			// Assertions here
		})
	}
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
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
