package main

import (
	"fmt"
	"testing"

	"github.com/adsmf/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 102
	//Part 2: 327300950120029
}

func TestPart2Examples(t *testing.T) {
	type testDef struct {
		input  string
		expect int
	}
	tests := []testDef{
		{"2,5", 4},
		{"7,5,2", 14},
		{"17,x,13,19", 3417},
		{"67,7,59,61", 754018},
		{"67,x,7,59,61", 779210},
		{"67,7,x,59,61", 1261476},
		{"1789,37,47,1889", 1202161486},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			result := part2(test.input)
			assert.Equal(t, test.expect, result)
		})
	}
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkPart1(b *testing.B) {
	lines := utils.ReadInputLines("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(lines)
	}
}

func BenchmarkPart2(b *testing.B) {
	lines := utils.ReadInputLines("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(lines[1])
	}
}
