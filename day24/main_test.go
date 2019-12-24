package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	p := loadMap("examples/ex1.0.txt")
	t.Logf("%v\n", p)

	p.iter()
	t.Logf("After 1:\n%v\n", p)
	p1 := loadMap("examples/ex1.1.txt")
	assert.Equal(t, p.String(), p1.String())

	p.iter()
	t.Logf("After 2:\n%v\n", p)
	p2 := loadMap("examples/ex1.2.txt")
	assert.Equal(t, p.String(), p2.String())

}

func TestBiodiversity(t *testing.T) {
	p := loadMap("examples/biodiversity.txt")
	assert.Equal(t, 2129920, p.biodiversity())
}

func TestAnswers(t *testing.T) {
	assert.Equal(t, 11042850, part1())
	assert.Equal(t, 0, part2())
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 11042850
	//Part 2: 0
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
