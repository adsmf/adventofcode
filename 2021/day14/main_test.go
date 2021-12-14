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
	//Part 1: 2975
	//Part 2: 3015383850689
}

func TestExamples(t *testing.T) {
	polymer, rules := load(example)
	p1, p2 := expandPolymers(polymer, rules)
	assert.Equal(t, 1588, p1)
	assert.Equal(t, 2188189693529, p2)
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
