package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed example1.txt
var example1 string

func ExampleMain() {
	main()
	//Output:
	//Part 1: 1599
	//Part 2: 418
}

func TestPart1Examples(t *testing.T) {
	g := load(example1)
	p1, p2 := runSim(g)
	assert.Equal(t, p1, 1656)
	assert.Equal(t, p2, 195)
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
