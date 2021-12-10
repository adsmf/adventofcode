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
	//Part 1: 399153
	//Part 2: 2995077699
}

func TestAnswers(t *testing.T) {
	p1, p2 := solve(input)
	assert.Equal(t, 399153, p1)
	assert.Equal(t, 2995077699, p2)
}

func TestExamples(t *testing.T) {
	p1, p2 := solve(example)
	assert.Equal(t, 26397, p1)
	assert.Equal(t, 288957, p2)
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
