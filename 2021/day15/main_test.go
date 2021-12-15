package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed example.txt
var example string

func ExampleMain() {
	main()
	//Output:
	//Part 1: 619
	//Part 2: 2922
}

func TestExamples(t *testing.T) {
	g, size := load(example)
	p1 := part1(g, size)
	p2 := part2(g, size)
	assert.Equal(t, 40, p1)
	assert.Equal(t, 315, p2)
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
