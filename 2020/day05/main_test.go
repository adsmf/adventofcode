package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 832
	//Part 2: 517
}

func TestLoaders(t *testing.T) {
	l1 := loadStringparse("input.txt")
	l2 := loadBitwise("input.txt")
	assert.EqualValues(t, l1, l2)
}

func BenchmarkStringparse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loadStringparse("input.txt")
	}
}

func BenchmarkBitwise(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loadBitwise("input.txt")
	}
}
