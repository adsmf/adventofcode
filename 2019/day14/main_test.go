package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	type testDef struct {
		file string
		ore  int
	}
	tests := []testDef{
		testDef{"ex1.txt", 31},
		testDef{"ex2.txt", 165},
		testDef{"ex3.txt", 13312},
		testDef{"ex4.txt", 180697},
		testDef{"ex5.txt", 2210736},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part 1 test %d", id), func(t *testing.T) {
			reactions := loadInput("examples/" + test.file)
			req, _ := calculateRequired(chemicalQuantity{"FUEL", 1}, reactions, surplusMap{})
			assert.Equal(t, test.ore, req)
		})
	}
}

func TestPart2Examples(t *testing.T) {

}

func TestAnswers(t *testing.T) {
	assert.Equal(t, 301997, part1())
	assert.Equal(t, 6216589, part2())
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 301997
	//Part 2: 6216589
}

func BenchmarkPart1(b *testing.B) {
	reactions := loadInput("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculateRequired(chemicalQuantity{"FUEL", 1}, reactions, surplusMap{})
	}
}

func BenchmarkPart2(b *testing.B) {
	reactions := loadInput("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calcMaxHold(1000000000000, reactions)
	}
}
