package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Answer(t *testing.T) {
	prog := loadInputString()
	outputs := gatherOutputs(prog, 1)
	assert.Equal(t, []int{3063082071}, outputs)
}
func BenchmarkPart1(b *testing.B) {
	inputString := loadInputString()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gatherOutputs(inputString, 1)
	}
}

func TestPart2Answer(t *testing.T) {
	prog := loadInputString()
	outputs := gatherOutputs(prog, 2)
	assert.Equal(t, []int{81348}, outputs, "Should retun coordinate")
}

func BenchmarkPart2(b *testing.B) {
	inputString := loadInputString()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gatherOutputs(inputString, 2)
	}
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 3063082071
	//Part 2: 81348
}
