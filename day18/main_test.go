package main

import (
	"fmt"
	"testing"

	"github.com/adsmf/adventofcode2019/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	type testDef struct {
		file  string
		steps int
	}
	tests := []testDef{
		testDef{
			"amf1.txt",
			2,
		},
		testDef{
			"amf2.txt",
			8,
		},
		testDef{
			file:  "p1ex0.txt",
			steps: 8,
		},
		testDef{
			file:  "p1ex1.txt",
			steps: 86,
		},
		testDef{
			file:  "p1ex2.txt",
			steps: 132,
		},
		testDef{
			file:  "p1ex3.txt",
			steps: 136,
		},
		testDef{
			file:  "p1ex4.txt",
			steps: 81,
		},
	}
	for id, test := range tests {
		file := "examples/" + test.file
		expected := test.steps
		for i := 0; i < 1; i++ {
			t.Run(fmt.Sprintf("Part1 Test%d Iter%d", i, id), func(t *testing.T) {
				crawl := loadMap(file)
				t.Logf("Start:\n%v", crawl)
				steps := crawl.collectKeys()
				assert.Greater(t, utils.MaxInt, steps)
				assert.Equal(t, expected, steps)
			})
		}
	}
}

func TestKeyring(t *testing.T) {
	type testDef struct {
		initial keyring
		add     int
		expect  keyring
	}

	tests := []testDef{
		testDef{
			initial: keyring(""),
			add:     'a',
			expect:  keyring("a"),
		},
		testDef{
			initial: keyring("a"),
			add:     'a',
			expect:  keyring("a"),
		},
		testDef{
			initial: keyring("a"),
			add:     'b',
			expect:  keyring("ab"),
		},
		testDef{
			initial: keyring("bc"),
			add:     'a',
			expect:  keyring("abc"),
		},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Keyring %d", id), func(t *testing.T) {
			result := test.initial.with(test.add)
			assert.Equal(t, test.expect, result)
		})
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
			8,
		},
		testDef{
			"p2ex2.txt",
			32,
		},
	}
	for id, test := range tests {
		file := "examples/" + test.file
		expected := test.steps
		for i := 0; i < 1; i++ {
			t.Run(fmt.Sprintf("Part2 Test%d Iter%d", i, id), func(t *testing.T) {
				crawl := loadMap(file)
				t.Logf("Start:\n%v", crawl)
				steps := crawl.collectKeys()
				assert.Greater(t, utils.MaxInt, steps)
				assert.Equal(t, expected, steps)
			})
		}
	}
}

func TestAnswers(t *testing.T) {
	assert.Equal(t, 6162, part1())
	assert.Equal(t, 1556, part2())
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 6162
	//Part 2: 1556
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := loadMap("input.txt")
		v.collectKeys()
	}
}
func BenchmarkPart1Reduced(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := loadMap("input.txt")
		v.reduce()
		v.collectKeys()
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := loadMap("input2.txt")
		v.collectKeys()
	}
}

func BenchmarkPart2Reduced(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := loadMap("input2.txt")
		v.reduce()
		v.collectKeys()
	}
}
