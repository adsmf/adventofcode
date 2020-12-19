package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 12897
	//Part 2: 7031
}

func TestPart1Examples(t *testing.T) {
	type testDef struct {
		input string
		score int
	}
	tests := []testDef{
		{"{}", 1},
		{"{{{}}}", 6},
		{"{{},{}}", 5},
		{"{{{},{},{{}}}}", 16},
		{"{<a>,<a>,<a>,<a>}", 1},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
		{"{{<a!>},{<a!>},{<<><><<><a!>},{<ab><>}}", 3},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			result, _ := process([]byte(test.input))
			assert.Equal(t, test.score, result)
		})
	}
}

func TestPart2Examples(t *testing.T) {
	type testDef struct {
		input  string
		length int
	}
	tests := []testDef{
		{"<>", 0},
		{"<{!>}>", 2},
		{`<!!!>>`, 0},
		{`<{o"i!a,<{i<a>`, 10},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			offset, count := findCompleteGarbage([]byte(test.input)[1:])
			assert.Equal(t, offset+2, len(test.input))
			assert.Equal(t, count, test.length)
		})
	}
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
