package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 805932
	//Part 2: 133029050096658
}

func TestPart1Examples(t *testing.T) {
	p1 := part1(" 1 4 2 8")
	assert.Equal(t, 739785, p1)
}

func TestPart2Examples(t *testing.T) {
	p2 := part2(" 1 4 2 8")
	assert.Equal(t, 444356092776315, p2)
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
