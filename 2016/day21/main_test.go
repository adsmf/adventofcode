package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: gfdhebac
	//Part 2: dhaegfbc
}

func TestPart1Examples(t *testing.T) {
	ops := load("example.txt")
	result := process("abcde", ops)
	assert.Equal(t, "decab", result)
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
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
