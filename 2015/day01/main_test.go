package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {

}

func TestPart2Examples(t *testing.T) {

}

func TestAnswers(t *testing.T) {
	assert.Equal(t, 138, part1())
	assert.Equal(t, 1771, part2())
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 138
	//Part 2: 1771
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2()
	}
}
