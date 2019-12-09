package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay5Part1Examples(t *testing.T) {
	tests := map[string]string{
		// Previous
		"1,0,0,0,99":                    "2,0,0,0,99",
		"2,3,0,3,99":                    "2,3,0,6,99",
		"2,4,4,5,99,0":                  "2,4,4,5,99,9801",
		"1,1,1,4,99,5,6,0,99":           "30,1,1,4,2,5,6,0,99",
		"1,9,10,3,2,3,11,0,99,30,40,50": "3500,9,10,70,2,3,11,0,99,30,40,50",
		// New
		"1002,4,3,4,33": "1002,4,3,4,99",
		// Mine
		"3,1,99": "3,1002,99",
	}

	for program, expected := range tests {
		t.Run("Day5 "+program, func(t *testing.T) {
			inputs := make(chan int64, 1)
			outputs := make(chan int64)
			var output int64
			go func() {
				output = <-outputs
				t.Logf("Got output: %d", output)
			}()
			inputs <- 1002
			mach := newMachine(program, inputs, outputs)
			assert.Equal(t, int64(0), output)
			assert.NotPanics(t, func() {
				mach.run(-1)
			})
			assert.Equal(t, expected, mach.String())
		})
	}
}

func TestDay5Part2Examples(t *testing.T) {
	type part2inputtest map[int64]int64

	tests := map[string]part2inputtest{
		"3,9,8,9,10,9,4,9,99,-1,8": part2inputtest{
			7: 0,
			8: 1,
			9: 0,
		},
		"3,9,7,9,10,9,4,9,99,-1,8": part2inputtest{
			7: 1,
			8: 0,
			9: 0,
		},
		"3,3,1108,-1,8,3,4,3,99": part2inputtest{
			7: 0,
			8: 1,
			9: 0,
		},
		"3,3,1107,-1,8,3,4,3,99": part2inputtest{
			7: 1,
			8: 0,
			9: 0,
		},
		"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9": part2inputtest{
			-1: 1,
			0:  0,
			1:  1,
			2:  1,
		},
		"3,3,1105,-1,9,1101,0,0,12,4,12,99,1": part2inputtest{
			-1: 1,
			0:  0,
			1:  1,
			2:  1,
		},
		"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99": part2inputtest{
			7: 999,
			8: 1000,
			9: 1001,
		},
	}

	for prog, expected := range tests {
		for input, expectedOutput := range expected {
			t.Run(fmt.Sprintf("Day 5 part 2 - %s - %d", prog, input), func(t *testing.T) {
				assert.NotPanics(t, func() {
					outputs := gatherOutputs(prog, -1, input)
					assert.Equal(t, expectedOutput, outputs[0])
				})
			})
		}
	}
}

func TestDay9Part1Examples(t *testing.T) {
	type part2inputtest map[int64][]int64

	tests := map[string]part2inputtest{
		// Documented
		"109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99": part2inputtest{
			0: []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		"1102,34915192,34915192,7,4,7,99,0": part2inputtest{
			0: []int64{1219070632396864},
		},
		"104,1125899906842624,99": part2inputtest{
			0: []int64{1125899906842624},
		},
		"104,112589990684262400,99": part2inputtest{
			0: []int64{112589990684262400},
		},
		// Mine
		"109,12,203,1,4,12,4,13,204,2,99,11,12,13,14": part2inputtest{
			0: []int64{12, 0, 14},
			1: []int64{12, 1, 14},
		},

		////////////////
		// Positional //
		////////////////
		// - Add
		"1,3,3,10,4,10,99": part2inputtest{
			0: []int64{20},
		},
		// - Mult
		"2,3,3,10,4,10,99": part2inputtest{
			0: []int64{100},
		},
		// - Input
		"3,10,4,10,99": part2inputtest{
			-100: []int64{-100},
			0:    []int64{0},
			100:  []int64{100},
		},
		// - Output
		"1101,12,13,20,4,20,99": part2inputtest{
			0: []int64{25},
		},
		// - JNZ
		"1101,8,1,0,5,2,1,99,104,1,99": part2inputtest{
			0: []int64{1},
		},
		// - JEZ
		"1101,8,0,0,6,2,1,99,104,1,99": part2inputtest{
			0: []int64{1},
		},
		// - CLT
		// - CEQ
		// - URB
		"1101,5,3,20,9,20,204,4,99,1,2,3,4": part2inputtest{
			0: []int64{4},
		},

		///////////////
		// Immediate //
		///////////////
		// - Add
		"1101,3,3,10,4,10,99": part2inputtest{
			0: []int64{6},
		},
		// - Mult
		"1102,3,3,10,4,10,99": part2inputtest{
			0: []int64{9},
		},
		// - Input
		// ----- N/A -----
		// - Output
		"104,1,99": part2inputtest{
			0: []int64{1},
		},
		// - JNZ
		"1101,3,3,20,1105,1,8,99,104,1,99": part2inputtest{
			0: []int64{1},
		},
		// - JEZ
		"1101,3,3,20,1106,0,8,99,104,1,99": part2inputtest{
			0: []int64{1},
		},
		// - CLT
		// - CEQ
		// - URB
		"109,4,204,4,99,1,2,3,4": part2inputtest{
			0: []int64{4},
		},

		//////////////
		// Relative //
		//////////////
		// - Add
		"109,2,22201,3,3,10,4,12,99": part2inputtest{
			0: []int64{20},
		},
		// - Mult
		"109,2,22202,3,3,10,4,12,99": part2inputtest{
			0: []int64{100},
		},
		// - Input
		"109,2,203,10,4,12,99": part2inputtest{
			-100: []int64{-100},
			0:    []int64{0},
			100:  []int64{100},
		},
		// - Output
		"109,2,204,0,99": part2inputtest{
			0: []int64{204},
		},
		// - JNZ
		"109,2,1101,10,1,20,2205,2,1,99,104,1,99": part2inputtest{
			0: []int64{1},
		},
		// - JEZ
		"109,2,1101,10,0,20,2206,2,1,99,104,1,99": part2inputtest{
			0: []int64{1},
		},
		// - CLT
		// - CEQ
		// - URB

		// 	"1101,3,3,10,4,10,99": part2inputtest{
		// 		0: []int64{6},
		// 	},
	}

	for prog, expected := range tests {
		for input, expectedOutput := range expected {
			t.Run(fmt.Sprintf("Day 9 part 2 - %s - %d", prog, input), func(t *testing.T) {
				assert.NotPanics(t, func() {
					outputs := gatherOutputs(prog, -1, input)
					assert.Equal(t, expectedOutput, outputs)
				})
			})
		}
	}
}

func TestParamModeDecode(t *testing.T) {
	assert.Equal(t, int64(0), paramMode(10, 0))
	assert.Equal(t, int64(1), paramMode(10, 1))
	assert.Equal(t, int64(2), paramMode(212, 0))
	assert.Equal(t, int64(1), paramMode(212, 1))
	assert.Equal(t, int64(2), paramMode(212, 2))
}

func TestPart1Answer(t *testing.T) {
	prog := loadInputString()
	outputs := gatherOutputs(prog, -1, 1)
	assert.Equal(t, []int64{3063082071}, outputs)
}

func TestPart2Answer(t *testing.T) {
	prog := loadInputString()
	outputs := gatherOutputs(prog, -1, 2)
	assert.Equal(t, []int64{81348}, outputs, "Should retun coordinate")
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 3063082071
	//Part 2: 81348
}
