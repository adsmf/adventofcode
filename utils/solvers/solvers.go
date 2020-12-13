package solvers

import (
	"fmt"
	"math/big"
)

// ChineseRemainderTheorem solution taken almost verbatim from
//   https://www.rosettacode.org/wiki/Chinese_remainder_theorem#Go
func ChineseRemainderTheorem(remainders, modulos []*big.Int) (*big.Int, error) {
	var one = big.NewInt(1)
	p := new(big.Int).Set(modulos[0])
	for _, n1 := range modulos[1:] {
		p.Mul(p, n1)
	}
	var result, quotient, coeff, gcd big.Int
	for i, n1 := range modulos {
		quotient.Div(p, n1)
		gcd.GCD(nil, &coeff, n1, &quotient)
		if gcd.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		result.Add(&result, coeff.Mul(remainders[i], coeff.Mul(&coeff, &quotient)))
	}
	return result.Mod(&result, p), nil
}
