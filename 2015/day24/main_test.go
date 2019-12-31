package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	p := []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}
	total := totalWeight(p)
	target := total / 3
	fmt.Printf("Packages: %v\nTotal: %d (3*%d)\n", p, total, target)
	result := findSmallest3(p, target)
	assert.Equal(t, 99, result)
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

func TestAnswers(t *testing.T) {
}

// func ExampleMain() {
// 	main()
// 	//Output:
// 	//Part 1: -1
// 	//Part 2: -1
// }

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2()
	}
}
