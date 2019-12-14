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
	}
	tests := []testDef{
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
		testDef{
			program:  "1002,4,3,4,33",
			endState: "1002,4,3,4,99",
		},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Test %d", id), func(t *testing.T) {
			t.Logf("Test definition:\n%#v", test)
			inputStream := make(chan int)
			outputStream := make(chan int)

			var m Machine
			assert.NotPanics(t, func() {
				m = NewMachine(M19(inputStream, outputStream))
				m.LoadProgram(test.program)
			})
			t.Logf("Initial machine state:\n%v", m)

			if test.endState != "" {
				for i := 1; ; i++ {
					var rc ExecReturnCode
					assert.NotPanics(t, func() {
						rc = m.Step()
					})
					t.Logf("Step %d:\n%v\nRC: %d", i, m, rc)
					if rc != ExecRCNone {
						break
					}
				}
				assert.Equal(t, test.endState, m.ram.String())
			}
		})
	}
}
