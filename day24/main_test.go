package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnswers(t *testing.T) {
	assert.Equal(t, 11042850, part1())
	assert.Equal(t, 1967, part2())
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 11042850
	//Part 2: 1967
}

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

func TestPart2Example(t *testing.T) {
	p := loadMap("examples/ex1.0.txt")
	p.iterRecursiveN(0)
	t.Logf("Initial:\n%v", p)
	result := p.iterRecursiveN(10)
	t.Logf("After 10:\n%v", p)
	assert.Equal(t, 99, result)
}

func TestRecursiveNeighbours19(t *testing.T) {
	pos := point{3, 3, 0}
	expected := []point{
		point{2, 3, 0},
		point{4, 3, 0},
		point{3, 2, 0},
		point{3, 4, 0},
	}
	neighbours := pos.recursiveNeighbours(5, 5)
	t.Logf("Expected:\n%v", expected)
	t.Logf("Got:\n%v", neighbours)
	assert.Subset(t, expected, neighbours)
	assert.Subset(t, neighbours, expected)
}

func TestRecursiveNeighboursE(t *testing.T) {
	pos := point{4, 0, 0}
	expected := []point{
		point{3, 0, 0},
		point{3, 2, 1},
		point{4, 1, 0},
		point{2, 1, 1},
	}
	neighbours := pos.recursiveNeighbours(5, 5)
	t.Logf("Expected:\n%v", expected)
	t.Logf("Got:\n%v", neighbours)
	assert.Subset(t, neighbours, expected)
	assert.Subset(t, expected, neighbours)
}

func TestRecursiveNeighboursU(t *testing.T) {
	pos := point{0, 4, -1}
	expected := []point{
		point{1, 2, 0},
		point{1, 4, -1},
		point{0, 3, -1},
		point{2, 3, 0},
	}
	neighbours := pos.recursiveNeighbours(5, 5)
	t.Logf("Expected:\n%v", expected)
	t.Logf("Got:\n%v", neighbours)
	assert.Subset(t, neighbours, expected)
	assert.Subset(t, expected, neighbours)
}

func TestRecursiveNeighbours14(t *testing.T) {
	pos := point{3, 2, 0}
	expected := []point{
		// Same level
		point{3, 1, 0},
		point{3, 3, 0},
		point{4, 2, 0},
		// Inner level
		point{4, 0, -1},
		point{4, 1, -1},
		point{4, 2, -1},
		point{4, 3, -1},
		point{4, 4, -1},
	}
	neighbours := pos.recursiveNeighbours(5, 5)
	t.Logf("Expected:\n%v", expected)
	t.Logf("Got:\n%v", neighbours)
	assert.Len(t, neighbours, 8)
	assert.Subset(t, neighbours, expected)
	assert.Subset(t, expected, neighbours)
}

func TestRecursiveNeighbours8(t *testing.T) {
	pos := point{2, 1, 0}
	expected := []point{
		// Same level
		point{2, 0, 0},
		point{1, 1, 0},
		point{3, 1, 0},
		// Inner level
		point{0, 0, -1},
		point{1, 0, -1},
		point{2, 0, -1},
		point{3, 0, -1},
		point{4, 0, -1},
	}
	neighbours := pos.recursiveNeighbours(5, 5)
	t.Logf("Expected:\n%v", expected)
	t.Logf("Got:\n%v", neighbours)
	assert.Len(t, neighbours, 8)
	assert.Subset(t, neighbours, expected)
	assert.Subset(t, expected, neighbours)
}

func TestRecursiveNeighbours18(t *testing.T) {
	pos := point{2, 3, 0}
	expected := []point{
		// Same level
		point{2, 4, 0},
		point{1, 3, 0},
		point{3, 3, 0},
		// Inner level
		point{0, 4, -1},
		point{1, 4, -1},
		point{2, 4, -1},
		point{3, 4, -1},
		point{4, 4, -1},
	}
	neighbours := pos.recursiveNeighbours(5, 5)
	t.Logf("Expected:\n%v", expected)
	t.Logf("Got:\n%v", neighbours)
	assert.Len(t, neighbours, 8)
	assert.Subset(t, neighbours, expected)
	assert.Subset(t, expected, neighbours)
}

func TestRecursiveNeighboursL(t *testing.T) {
	pos := point{1, 2, -1}
	expected := []point{
		// Same level
		point{1, 1, -1},
		point{1, 3, -1},
		point{0, 2, -1},
		// Inner level
		point{0, 0, -2},
		point{0, 1, -2},
		point{0, 2, -2},
		point{0, 3, -2},
		point{0, 4, -2},
	}
	neighbours := pos.recursiveNeighbours(5, 5)
	t.Logf("Expected:\n%v", expected)
	t.Logf("Got:\n%v", neighbours)
	assert.Len(t, neighbours, 8)
	assert.Subset(t, neighbours, expected)
	assert.Subset(t, expected, neighbours)
}

func TestBiodiversity(t *testing.T) {
	p := loadMap("examples/biodiversity.txt")
	assert.Equal(t, 2129920, p.biodiversity())
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
