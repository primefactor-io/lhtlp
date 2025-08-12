package homomorphic

import (
	"math/big"

	"github.com/primefactor-io/lhtlp/pkg/params"
	"github.com/primefactor-io/lhtlp/pkg/puzzle"
)

// MultiplyPlaintextValue multiplies the plaintext value with the value that is
// hidden in the puzzle.
func MultiplyPlaintextValue(params *params.Params, z *puzzle.Puzzle, p *big.Int) *puzzle.Puzzle {
	u := new(big.Int).Exp(z.U, p, params.N)     // u^p mod n
	v := new(big.Int).Exp(z.V, p, params.NExpY) // v^p mod n^y

	puzzle := puzzle.NewPuzzle(u, v)

	return puzzle
}
