package utils

import "math/big"

// Factorial computes the factorial of x which is x! = 1 * 2 * 3 * ... * x.
// See: https://stackoverflow.com/a/54952509
func Factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)

	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}

	return n.Mul(x, Factorial(n.Sub(x, n)))
}

// Exponentiate computes the exponentiations n^(x - 1), n^x and n^(x + 1).
// Note: The caller needs to ensure that x has a value in the correct range
// (e.g. that n^(x - 1) can be computed correctly).
func Exponentiate(n *big.Int, x int) (*big.Int, *big.Int, *big.Int) {
	var n1 = big.NewInt(1) // n^(x - 1)
	var n2 = big.NewInt(1) // n^x
	var n3 = big.NewInt(1) // n^(x + 1)

	for i := 1; i <= x+1; i++ {
		if i <= x-1 {
			n1 = new(big.Int).Mul(n1, n)
		}

		if i <= x {
			n2 = new(big.Int).Mul(n2, n)
		}

		if i <= x+1 {
			n3 = new(big.Int).Mul(n3, n)
		}
	}

	return n1, n2, n3
}
