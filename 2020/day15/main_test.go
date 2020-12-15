package main

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/adsmf/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 866
	//Part 2: 1437692
}

func TestExamples(t *testing.T) {
	type testDef struct {
		input  string
		turns  int
		expect int
	}
	tests := []testDef{
		{"0,3,6", 4, 0},
		{"0,3,6", 5, 3},
		{"0,3,6", 6, 3},
		{"0,3,6", 7, 1},
		{"0,3,6", 8, 0},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%s-Turn%d", test.input, test.turns), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			numbers := utils.GetInts(test.input)
			result := play(numbers, test.turns)
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
	input, _ := ioutil.ReadFile("input.txt")
	numbers := utils.GetInts(string(input))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		play(numbers, 2020)
	}
}

func BenchmarkPart2(b *testing.B) {
	input, _ := ioutil.ReadFile("input.txt")
	numbers := utils.GetInts(string(input))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		play(numbers, 30000000)
	}
}
