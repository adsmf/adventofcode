package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	type testDef struct {
		input  string
		output string
		rounds int
		// Test structure here
	}
	tests := []testDef{
		testDef{"1", "11", 1},
		testDef{"111221", "312211", 1},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)

			result := test.input
			for i := 0; i < test.rounds; i++ {
				result = next(result)
			}
			assert.Equal(t, test.output, result)
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

func ExampleMain() {
	main()
	//Output:
	//Part 1: 492982
	//Part 2: 6989950
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
