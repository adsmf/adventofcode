package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed example1.txt
var example1 string

func ExampleMain() {
	main()
	//Output:
	//Part 1: 10411
	//Part 2: 46721
}

func TestExample(t *testing.T) {
	initial := load(example1)
	p1 := calcEnergy(initial)
	require.Equal(t, 12521, p1)
	p2 := part2(initial)
	assert.Equal(t, 44169, p2)
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
