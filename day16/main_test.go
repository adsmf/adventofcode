package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples1(t *testing.T) {
	pattern := signal{0, 1, 0, -1}
	type testData struct {
		fftInput  signal
		times     int
		fftOutput signal
	}
	tests := []testData{
		testData{
			fftInput:  signal{1, 2, 3, 4, 5, 6, 7, 8},
			fftOutput: signal{4, 8, 2, 2, 6, 1, 5, 8},
			times:     1,
		},
		testData{
			fftInput:  signal{1, 2, 3, 4, 5, 6, 7, 8},
			fftOutput: signal{0, 1, 0, 2, 9, 4, 9, 8},
			times:     4,
		},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Day16Part1-ex%d", id), func(t *testing.T) {
			result := fftTimes(test.fftInput, pattern, test.times)
			assert.Equal(t, test.fftOutput, result)
		})
	}
}

func TestPart1Examples2(t *testing.T) {
	tests := map[string]string{
		"80871224585914546619083218645595": "24176176",
		"19617804207202209144916044189917": "73745418",
		"69317163492948606335995924319873": "52432133",
	}
	for input, output := range tests {
		t.Run(fmt.Sprintf("Day16Part1-%s", input), func(t *testing.T) {
			// result := fftTimes(output.fftInput, pattern, output.times)
			result := fftString(input, 100)
			assert.Equal(t, output, result)
		})
	}
}

func TestPart2Examples(t *testing.T) {

}

func TestAnswers(t *testing.T) {
	assert.Equal(t, "29795507", part1())
	assert.Equal(t, 0, part2())
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 29795507
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
