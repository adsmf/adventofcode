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
			endState: "1,0,0,0,99",
		},
		testDef{
			program:  "1,0,0,0,99,0,0,0,0,0",
			endState: "1,0,0,0,99,0,0,0,0,0",
		},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Test %d", id), func(t *testing.T) {
			t.Logf("Test definition:\n%#v", test)
			inputStream := make(chan int)
			outputStream := make(chan int)
			m := NewMachine(M19(inputStream, outputStream))
			m.LoadProgram(test.program)
			t.Logf("Initial machine state:\n%v", m)

			if test.endState != "" {
				assert.Equal(t, test.endState, m.ram.String())
			}
		})
	}
}
