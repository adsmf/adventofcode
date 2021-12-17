package main

import (
	_ "embed"
	"testing"

	"github.com/adsmf/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

//go:embed example.txt
var example string

func ExampleMain() {
	main()
	//Output:
	//Part 1: 30628
	//Part 2: 4433
}

func TestExample(t *testing.T) {
	vals := utils.GetInts(example)
	minX, maxX, minY, maxY := vals[0], vals[1], vals[2], vals[3]
	p1 := part1(minY)
	p2 := part2(minX, maxX, minY, maxY)
	assert.Equal(t, 45, p1)
	assert.Equal(t, 112, p2)
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
