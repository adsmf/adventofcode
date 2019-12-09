package intcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModel2019(t *testing.T) {
	type testDef struct {
		program string
	}
	tests := []testDef{
		testDef{
			program: "1,0,0,0,99",
		},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Test %d", id), func(t *testing.T) {
			t.Logf("Test definition:\n%#v", test)
			inputStream := make(chan int)
			outputStream := make(chan int)
			m := NewMachine(Model2019(inputStream, outputStream))
			m.LoadProgram(test.program)
			t.Logf("Initial machine state:\n%v", m)
			assert.Equal(t, true, false)
		})
	}
}
