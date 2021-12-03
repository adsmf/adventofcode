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

func BenchmarkFast(b *testing.B) {
	b.Run("ParseInput", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			parseInputInts()
		}
	})
	ints, bitlen := parseInputInts()
	b.Run("P1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			part1(ints, bitlen)
		}
	})
	b.Run("P2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			part2(ints, bitlen)
		}
	})
}

func BenchmarkInitial(b *testing.B) {
	b.Run("P1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			part1initial()
		}
	})
	b.Run("P2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			part2initial()
		}
	})
}
