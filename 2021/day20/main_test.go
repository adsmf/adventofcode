package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed example1.txt
var example string

func ExampleMain() {
	main()
	//Output:
	//Part 1: 5846
	//Part 2: 21149
}

func TestExample(t *testing.T) {
	p1, p2 := solve(example)
	assert.Equal(t, 35, p1)
	assert.Equal(t, 3351, p2)
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
