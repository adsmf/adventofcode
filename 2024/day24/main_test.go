package main

import (
	"fmt"
	"testing"

	"github.com/adsmf/adventofcode/utils"
	"github.com/stretchr/testify/require"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 55920211035878
	//Part 2: btb,cmv,mwp,rdg,rmj,z17,z23,z30
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func TestHashingChar(t *testing.T) {
	for num := range 10 {
		val := byte(num + '0')
		t.Run(fmt.Sprintf("%c", val), func(t *testing.T) {
			require.Equal(t, val, unhashCh(hashChar(val)))
		})
	}
	for let := range 26 {
		val := byte(let + 'a')
		t.Run(fmt.Sprintf("%c", val), func(t *testing.T) {
			require.Equal(t, val, unhashCh(hashChar(val)))
		})
	}
}

func TestHashingFull(t *testing.T) {
	utils.EachLine(input, func(_ int, line string) (done bool) {
		if len(line) < 3 {
			return
		}
		val := line[:3]
		t.Run(val, func(t *testing.T) {
			hashed := hash(val)
			require.Equal(t, val, hashed.String())
		})
		return
	})
}
