package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	filename := "examples/p1ex1.txt"
	expectedSum := 76

	input, err := ioutil.ReadFile(filename)
	assert.NoError(t, err)

	s := scaffold{}
	s.cameraViewRaw = string(input)

	s.processCamera()
	assert.Equal(t, expectedSum, s.sumAlignment())
}

func TestPart2Examples(t *testing.T) {

}

func TestAnswers(t *testing.T) {
	assert.Equal(t, 4372, part1())
	assert.Equal(t, 945911, part2())
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 4372
	//Part 2: 945911
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
