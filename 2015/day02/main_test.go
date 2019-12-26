package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	type testDef struct {
		input string
		paper int
	}
	tests := []testDef{
		testDef{"2x3x4", 58},
		testDef{"1x1x10", 43},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			c := newCuboid(test.input)
			t.Logf("Papering cube: %s", test.input)
			t.Logf("Cubout: %v", c)
			assert.Equal(t, test.paper, c.paper())
		})
	}
}

func TestPart2Examples(t *testing.T) {

}

func TestAnswers(t *testing.T) {
	assert.Equal(t, 1588178, part1())
	assert.Equal(t, 0, part2())
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 1588178
	//Part 2: 0
}

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
