package intcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestM19(t *testing.T) {
	type testDef struct {
		program  string
		endState string
		input    int
		output   int
	}
	tests := []testDef{
		// Add/multiply tests
		testDef{
			program:  "1,0,0,0,99",
			endState: "2,0,0,0,99",
		},
		testDef{
			program:  "1,0,0,0,99,0,0,0,0,0",
			endState: "2,0,0,0,99,0,0,0,0,0",
		},
		testDef{
			program:  "2,3,0,3,99",
			endState: "2,3,0,6,99",
		},
		testDef{
			program:  "2,4,4,5,99,0",
			endState: "2,4,4,5,99,9801",
		},
		testDef{
			program:  "1,1,1,4,99,5,6,0,99",
			endState: "30,1,1,4,2,5,6,0,99",
		},
		testDef{
			program:  "1,9,10,3,2,3,11,0,99,30,40,50",
			endState: "3500,9,10,70,2,3,11,0,99,30,40,50",
		},
		// Immediate mode
		testDef{
			program:  "1002,4,3,4,33",
			endState: "1002,4,3,4,99",
		},
		// I/O tests
		testDef{
			program: "3,9,8,9,10,9,4,9,99,-1,8",
			input:   7,
			output:  0,
		},
		testDef{
			program: "3,9,8,9,10,9,4,9,99,-1,8",
			input:   8,
			output:  1,
		},
		testDef{
			program: "3,9,8,9,10,9,4,9,99,-1,8",
			input:   9,
			output:  0,
		},

		testDef{
			program: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			input:   7,
			output:  999,
		},
		testDef{
			program: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			input:   8,
			output:  1000,
		},
		testDef{
			program: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			input:   9,
			output:  1001,
		},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Test %d", id), func(t *testing.T) {
			t.Logf("Test definition:\n%#v", test)
			inputStream := make(chan int, 1)
			outputStream := make(chan int)

			inputStream <- test.input

			var m Machine
			assert.NotPanics(t, func() {
				m = NewMachine(M19(inputStream, outputStream))
				m.LoadProgram(test.program)
			})
			t.Logf("Initial machine state:\n%v", m)

			for i := 1; ; i++ {
				var rc ExecReturnCode
				assert.NotPanics(t, func() {
					rc = m.Step()
				})
				t.Logf("Step %d:\n%v\nRC: %d", i, m, rc)
				if rc != ExecRCNone && rc != ExecRCInterrupt {
					break
				}
			}
			if test.endState != "" {
				assert.Equal(t, test.endState, m.ram.String())
			}

			assert.Equal(t, test.output, m.registers[M19RegisterOutput])
		})
	}
}
