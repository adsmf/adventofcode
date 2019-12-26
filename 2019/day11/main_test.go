package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestPart2Examples(t *testing.T) {
// 	tests := map[string]string{
// 		"amf":      "PFKHECZU",
// 		"amk":      "BCKFPCRA",
// 		"asm":      "BLULZJLZ",
// 		"dnwe":     "JHARBGCU",
// 		"hindessm": "RPJCFZKF",
// 	}

// 	for input, result := range tests {
// 		t.Run(fmt.Sprintf("Decode %s", input), func(t *testing.T) {
// 			inputString, err := ioutil.ReadFile("inputs/" + input + ".txt")
// 			assert.NoError(t, err)
// 			hull := runPainter(strings.TrimSpace(string(inputString)), 1)
// 			t.Logf("Expected:\n%s", hull.print())
// 			decode := hull.decode()
// 			assert.Equal(t, result, decode)
// 		})
// 	}
// }

func TestAnswers(t *testing.T) {
	assert.Equal(t, 1930, part1())
}

func ExampleMain() {
	main()
	//Output:
	// Part 1: 1930
	// Part 2:
	// .###..####.#..#.#..#.####..##..####.#..#
	// .#..#.#....#.#..#..#.#....#..#....#.#..#
	// .#..#.###..##...####.###..#......#..#..#
	// .###..#....#.#..#..#.#....#.....#...#..#
	// .#....#....#.#..#..#.#....#..#.#....#..#
	// .#....#....#..#.#..#.####..##..####..##.
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
