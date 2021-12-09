package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 489
	//Part 2: 1056330
}

func TestPart1Examples(t *testing.T) {
	data := loadData(example)
	p1 := part1(data)
	assert.Equal(t, 15, p1)
}

func TestPart2Examples(t *testing.T) {
	data := loadData(example)
	p2 := part2(data)
	assert.Equal(t, 1134, p2)
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
