package main

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed example1.txt
var example1 string

func ExampleMain() {
	main()
	//Output:
	//Part 1: 560
}

func TestAnswers(t *testing.T) {
}

func TestExamples(t *testing.T) {
	type testDef struct {
		input string
		moves int
	}
	tests := []testDef{
		{example1, 58},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			moves := settle(test.input)
			assert.Equal(t, test.moves, moves)
		})
	}
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
