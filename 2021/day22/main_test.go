package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed example1.txt
var example1 string

//go:embed example2.txt
var example2 string

func ExampleMain() {
	main()
	//Output:
	//Part 1: 611176
	//Part 2: 1201259791805392
}

func TestPart1Examples(t *testing.T) {
	r := loadRanges(example1)
	p1, _ := solve(r)
	assert.Equal(t, 590784, p1)
}

func TestPart2Example(t *testing.T) {
	r := loadRanges(example2)
	p1, p2 := solve(r)
	assert.Equal(t, 474140, p1)
	assert.Equal(t, 2758514936282235, p2)
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
