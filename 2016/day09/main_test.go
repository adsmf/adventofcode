package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	type testDef struct {
		input  string
		output string
	}
	tests := []testDef{
		{"ADVENT", "ADVENT"},
		{"A(1x5)BC", "ABBBBBC"},
		{"(3x3)XYZ", "XYZXYZXYZ"},
		{"A(2x2)BCD(2x2)EFG", "ABCBCDEFEFG"},
		{"(6x1)(1x3)A", "(1x3)A"},
		{"X(8x2)(3x3)ABCY", "X(3x3)ABC(3x3)ABCY"},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			unpacked := calculateLength(test.input, false)
			assert.Equal(t, len(test.output), unpacked)
		})
	}
}

func TestPart2Examples(t *testing.T) {
	type testDef struct {
		input  string
		length int
	}
	tests := []testDef{
		{"(3x3)XYZ", 9},
		{"X(8x2)(3x3)ABCY", 20},
		{"(27x12)(20x12)(13x14)(7x10)(1x12)A", 241920},
		{"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", 445},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			unpacked := calculateLength(test.input, true)
			assert.Equal(t, test.length, unpacked)
		})
	}
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 102239
	//Part 2: 10780403063
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkPart1(b *testing.B) {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	input := strings.TrimSpace(string(inputBytes))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculateLength(input, false)
	}
}

func BenchmarkPart2(b *testing.B) {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	input := strings.TrimSpace(string(inputBytes))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculateLength(input, true)
	}
}
