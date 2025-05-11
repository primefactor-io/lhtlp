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
