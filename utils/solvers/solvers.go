package solvers

import (
	"fmt"
	"math/big"

	"github.com/adsmf/adventofcode/utils"
	"golang.org/x/exp/constraints"
)

// ChineseRemainderTheorem solution taken almost verbatim from
//
//	https://www.rosettacode.org/wiki/Chinese_remainder_theorem#Go
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

// ChineseRemainderTheorem reimplemented from above without `big`
func ChineseRemainderTheoremStd[T constraints.Integer](remainders, modulos []T) (T, error) {
	p := modulos[0]
	for _, n1 := range modulos[1:] {
		p *= n1
	}
	res := T(0)
	for i, n1 := range modulos {
		q := p / n1
		gcd, coeff, _ := utils.ExtendedGreatestCommonDivisor(n1, q)
		if gcd != 1 {
			return 0, fmt.Errorf("%d not coprime", n1)
		}
		res += coeff * q * remainders[i]
	}
	return res % p, nil
}
