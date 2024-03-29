package main

import (
	"fmt"
	"testing"

	"github.com/adsmf/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	type testDef struct {
		file  string
		steps int
	}
	tests := []testDef{
		testDef{
			"p1ex1.txt",
			23,
		},
		testDef{
			"p1ex2.txt",
			58,
		},
	}
	for id, test := range tests {
		file := "examples/" + test.file
		expected := test.steps
		for i := 0; i < 1; i++ {
			t.Run(fmt.Sprintf("Part1 Test%d Iter%d", i, id), func(t *testing.T) {
				m := loadMap(file)
				t.Logf("Start:\n%v", m)
				steps := m.solve()
				assert.Greater(t, utils.MaxInt, steps)
				assert.Equal(t, expected, steps)
			})
		}
	}
}

func TestPart2Examples(t *testing.T) {
	type testDef struct {
		file  string
		steps int
	}
	tests := []testDef{
		testDef{
			"p2ex1.txt",
			396,
		},
	}
	for id, test := range tests {
		file := "examples/" + test.file
		expected := test.steps
		for i := 0; i < 1; i++ {
			t.Run(fmt.Sprintf("Part1 Test%d Iter%d", i, id), func(t *testing.T) {
				m := loadMap(file)
				m.reduce()
				t.Logf("Start:\n%v", m)
				m.recursive = true
				steps := m.solve()
				assert.Greater(t, utils.MaxInt, steps)
				assert.Equal(t, expected, steps)
			})
		}
	}
}

func TestAnswers(t *testing.T) {
	assert.Equal(t, 664, part1())
	assert.Equal(t, 7334, part2())
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 664
	//Part 2: 7334
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := loadMap("input.txt")
		m.solve()
	}
}

func BenchmarkPart1Reduced(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := loadMap("input.txt")
		m.reduce()
		m.solve()
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := loadMap("input.txt")
		m.recursive = true
		m.solve()
	}
}

func BenchmarkPart2Reduced(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := loadMap("input.txt")
		m.reduce()
		m.recursive = true
		m.solve()
	}
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
