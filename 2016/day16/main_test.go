package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 10100011010101011
	//Part 2: 01010001101011001
}

func TestAnswers(t *testing.T) {
}

func TestPart1Expansion(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"1", "100"},
		{"0", "001"},
		{"11111", "11111000000"},
		{"111100001010", "1111000010100101011110000"},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			result := expand(test.input)
			assert.Equal(t, test.output, result)
		})
	}
}

func TestPart1Checksum(t *testing.T) {
	tests := []struct {
		input    string
		checksum string
	}{
		{"110010110100", "100"},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			result := checksum(test.input)
			assert.Equal(t, test.checksum, result)
		})
	}
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
