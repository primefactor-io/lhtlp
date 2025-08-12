package homomorphic

import (
	"math/big"

	"github.com/primefactor-io/lhtlp/pkg/params"
	"github.com/primefactor-io/lhtlp/pkg/puzzle"
)

// AddPlaintextValues adds the plaintext values that were hidden in the puzzles.
func AddPlaintextValues(params *params.Params, puzzles ...*puzzle.Puzzle) *puzzle.Puzzle {
	var u *big.Int
	var v *big.Int

	for i, puzzle := range puzzles {
		// Compute new u.
		var in1 *big.Int = puzzle.U
		if i != 0 {
			in1 = new(big.Int).Mul(u, puzzle.U) // u_{i-1} * u_{i}
		}
		u = new(big.Int).Mod(in1, params.N) // u_a * u_b mod n

		// Compute new v.
		var in2 *big.Int = puzzle.V
		if i != 0 {
			in2 = new(big.Int).Mul(v, puzzle.V) // v_{i-1} * v_{i}
		}
		v = new(big.Int).Mod(in2, params.NExpY) // v_a * v_b mod n^y
	}

	puzzle := puzzle.NewPuzzle(u, v)

	return puzzle
}

// AddPlaintextValue adds the plaintext value to the value that is hidden in the
// puzzle.
func AddPlaintextValue(params *params.Params, z *puzzle.Puzzle, p *big.Int) *puzzle.Puzzle {
	// Compute u'.
	uPrime := new(big.Int).Exp(params.G, p, params.N) // g^p mod n

	// Compute v'.
	in1 := new(big.Int).Mul(p, params.NExpYMinusOne)     // p * n^(y - 1)
	in2 := new(big.Int).Exp(params.H, in1, params.NExpY) // h^(p * n^(y - 1)) mod n^y
	in3 := new(big.Int).Add(big.NewInt(1), params.N)     // 1 + n
	in4 := new(big.Int).Exp(in3, p, params.NExpY)        // (1 + n)^p mod n^y
	in5 := new(big.Int).Mul(in2, in4)                    // h^(p * n^(y - 1)) * (1 + n)^p
	vPrime := new(big.Int).Mod(in5, params.NExpY)        // h^(p * n^(y - 1)) * (1 + n)^p mod n^y

	in6 := new(big.Int).Mul(z.U, uPrime) // u * u'
	u := new(big.Int).Mod(in6, params.N) // u * u' mod n

	in7 := new(big.Int).Mul(z.V, vPrime)     // v * v'
	v := new(big.Int).Mod(in7, params.NExpY) // v * v' mod n^y

	puzzle := puzzle.NewPuzzle(u, v)

	return puzzle
}
