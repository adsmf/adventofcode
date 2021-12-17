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
	//Part 1: 30628
	//Part 2: 4433
}

func TestExample(t *testing.T) {
	p1, p2 := solve(example)
	assert.Equal(t, 45, p1)
	assert.Equal(t, 112, p2)
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
