package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 4001724
	//Part 2: 587895
}
func TestReimplementations(t *testing.T) {
	ints, bitlen := parseInputInts()
	assert.Equal(t, part1initial(), part1(ints, bitlen))
	assert.Equal(t, part2initial(), part2(ints, bitlen))
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkMethods(b *testing.B) {
	b.Run("InitialP1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			part1initial()
		}
	})
	b.Run("InitialP2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			part2initial()
		}
	})
	b.Run("ParseInput", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			parseInputInts()
		}
	})
	ints, bitlen := parseInputInts()
	b.Run("FastP1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			part1(ints, bitlen)
		}
	})
	b.Run("FastP2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			part2(ints, bitlen)
		}
	})
}
