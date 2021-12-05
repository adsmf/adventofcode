package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	type testDef struct {
		// Test structure here
	}
	tests := []testDef{
		// Test data here
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			// Assertions here
		})
	}
}

func TestPart2Examples(t *testing.T) {
	type testDef struct {
		input string
		nice  bool
	}
	tests := []testDef{
		testDef{"qjhvhtzxzqqjkmpb", true},
		testDef{"xxyxx", true},
		testDef{"uurcxstgmygtbstg", false},
		testDef{"ieodomkazucvgmuy", false},
		testDef{"xxyzxx", false},
		testDef{"xxyzxxodo", true},
		testDef{"odoxxyzxx", true},
		testDef{"aaaodo", false},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %#v", test)
			assert.Equal(t, test.nice, niceString2(test.input))
		})
	}
}

func TestAnswers(t *testing.T) {
	assert.Equal(t, 255, part1())
	assert.Equal(t, 55, part2())
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 255
	//Part 2: 55
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
