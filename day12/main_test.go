package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	type testDef struct {
		inputFile string
		states    map[int]string
	}

	tests := []testDef{
		testDef{
			inputFile: "ex1.txt",
			states: map[int]string{
				0: `pos=<x= -1, y=  0, z=  2>, vel=<x=  0, y=  0, z=  0>
pos=<x=  2, y=-10, z= -7>, vel=<x=  0, y=  0, z=  0>
pos=<x=  4, y= -8, z=  8>, vel=<x=  0, y=  0, z=  0>
pos=<x=  3, y=  5, z= -1>, vel=<x=  0, y=  0, z=  0>
`,
				1: `pos=<x=  2, y= -1, z=  1>, vel=<x=  3, y= -1, z= -1>
pos=<x=  3, y= -7, z= -4>, vel=<x=  1, y=  3, z=  3>
pos=<x=  1, y= -7, z=  5>, vel=<x= -3, y=  1, z= -3>
pos=<x=  2, y=  2, z=  0>, vel=<x= -1, y= -3, z=  1>
`,
				2: `pos=<x=  5, y= -3, z= -1>, vel=<x=  3, y= -2, z= -2>
pos=<x=  1, y= -2, z=  2>, vel=<x= -2, y=  5, z=  6>
pos=<x=  1, y= -4, z= -1>, vel=<x=  0, y=  3, z= -6>
pos=<x=  1, y= -4, z=  2>, vel=<x= -1, y= -6, z=  2>
`,
				3: `pos=<x=  5, y= -6, z= -1>, vel=<x=  0, y= -3, z=  0>
pos=<x=  0, y=  0, z=  6>, vel=<x= -1, y=  2, z=  4>
pos=<x=  2, y=  1, z= -5>, vel=<x=  1, y=  5, z= -4>
pos=<x=  1, y= -8, z=  2>, vel=<x=  0, y= -4, z=  0>
`,
				4: `pos=<x=  2, y= -8, z=  0>, vel=<x= -3, y= -2, z=  1>
pos=<x=  2, y=  1, z=  7>, vel=<x=  2, y=  1, z=  1>
pos=<x=  2, y=  3, z= -6>, vel=<x=  0, y=  2, z= -1>
pos=<x=  2, y= -9, z=  1>, vel=<x=  1, y= -1, z= -1>
`,
				10: `pos=<x=  2, y=  1, z= -3>, vel=<x= -3, y= -2, z=  1>
pos=<x=  1, y= -8, z=  0>, vel=<x= -1, y=  1, z=  3>
pos=<x=  3, y= -6, z=  1>, vel=<x=  3, y=  2, z= -3>
pos=<x=  2, y=  0, z=  4>, vel=<x=  1, y= -1, z= -1>
`,
			},
		},
		testDef{
			inputFile: "ex2.txt",
			states: map[int]string{
				0: `pos=<x= -8, y=-10, z=  0>, vel=<x=  0, y=  0, z=  0>
pos=<x=  5, y=  5, z= 10>, vel=<x=  0, y=  0, z=  0>
pos=<x=  2, y= -7, z=  3>, vel=<x=  0, y=  0, z=  0>
pos=<x=  9, y= -8, z= -3>, vel=<x=  0, y=  0, z=  0>
`,
				10: `pos=<x= -9, y=-10, z=  1>, vel=<x= -2, y= -2, z= -1>
pos=<x=  4, y= 10, z=  9>, vel=<x= -3, y=  7, z= -2>
pos=<x=  8, y=-10, z= -3>, vel=<x=  5, y= -1, z= -2>
pos=<x=  5, y=-10, z=  3>, vel=<x=  0, y= -4, z=  5>
`,
				100: `pos=<x=  8, y=-12, z= -9>, vel=<x= -7, y=  3, z=  0>
pos=<x= 13, y= 16, z= -3>, vel=<x=  3, y=-11, z= -5>
pos=<x=-29, y=-11, z= -1>, vel=<x= -3, y=  7, z=  4>
pos=<x= 16, y=-13, z= 23>, vel=<x=  7, y=  1, z=  1>
`,
			},
		},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part 1 - %d", id), func(t *testing.T) {
			for steps, result := range test.states {
				system := loadInput("examples/" + test.inputFile)
				if steps > 0 {
					t.Logf("Running %d steps", steps)
					system.run(steps)
				}
				assert.Equal(t, result, system.String())
			}
		})
	}
}
func TestPart1Energy(t *testing.T) {
	system := loadInput("examples/ex2.txt")
	system.run(100)
	assert.Equal(t, 1940, system.energy())
}

func TestPart2Examples(t *testing.T) {

	planets := loadInput("examples/ex1.txt")
	repeat := findRepeat(planets)
	assert.Equal(t, 2772, repeat)

	// planets = loadInput("examples/ex2.txt")
	// repeat = findRepeat(planets)
	// assert.Equal(t, 4686774924, repeat)
}

func TestAnswers(t *testing.T) {
	assert.Equal(t, 6227, part1())

	// assert.Equal(t, 0, part2())
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 6227

	//Part 2: ???
}

func BenchmarkStep(b *testing.B) {
	planets := loadInput("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		planets.step()
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
	}
}

// func BenchmarkPart2(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		part2()
// 	}
// }
