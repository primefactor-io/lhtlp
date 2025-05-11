package puzzle

import (
	"crypto/rand"
	"math/big"

	"github.com/primefactor-io/lhtlp/pkg/params"
	"github.com/primefactor-io/lhtlp/pkg/utils"
)

// Puzzle is an instance of a puzzle.
type Puzzle struct {
	// U is the puzzle's u value.
	U *big.Int
	// V is the puzzle's v value.
	V *big.Int
}

// NewPuzzle creates a new instance of a puzzle.
func NewPuzzle(u, v *big.Int) *Puzzle {
	return &Puzzle{
		U: u,
		V: v,
	}
}

// Equal checks if two puzzles are equal.
func (p *Puzzle) Equal(other *Puzzle) bool {
	return p.U.Cmp(other.U) == 0 && p.V.Cmp(other.V) == 0
}

// GeneratePuzzle generates a puzzle that hides the plaintext.
// Returns an error if the generation of the puzzle fails.
func GeneratePuzzle(params *params.Params, plaintext *big.Int) (*Puzzle, error) {
	puzzle, _, err := GeneratePuzzleAndReturnNonce(params, plaintext)
	if err != nil {
		return nil, err
	}

	return puzzle, nil
}

// GeneratePuzzleAndReturnNonce generates a puzzle that hides the plaintext while
// also returning the nonce that was used for randomness.
// Returns an error if the generation of the puzzle fails.
func GeneratePuzzleAndReturnNonce(params *params.Params, plaintext *big.Int) (*Puzzle, *big.Int, error) {
	nExpMinusOne := new(big.Int).Sub(params.NExpY, big.NewInt(1)) // n^y - 1

	// Sample a random nonce r.
	nonce, err := rand.Int(rand.Reader, nExpMinusOne)
	if err != nil {
		return nil, nil, ErrSampleNonceR
	}

	puzzle, err := GeneratePuzzleWithCustomNonce(params, nonce, plaintext)
	if err != nil {
		return nil, nil, err
	}

	return puzzle, nonce, err
}

// GeneratePuzzleWithCustomNonce generates a puzzle that hides the plaintext
// using the passed-in nonce for randomness.
// Returns an error if the generation of the puzzle fails.
func GeneratePuzzleWithCustomNonce(params *params.Params, nonce, plaintext *big.Int) (*Puzzle, error) {
	r := nonce
	s := plaintext

	// Compute u.
	u := new(big.Int).Exp(params.G, r, params.N) // g^r mod n

	// Compute v.
	in1 := new(big.Int).Mul(r, params.NExpYMinusOne)     // r * n^(y - 1)
	in2 := new(big.Int).Exp(params.H, in1, params.NExpY) // h^(r * n^(y - 1)) mod n^y
	in3 := new(big.Int).Add(big.NewInt(1), params.N)     // 1 + n
	in4 := new(big.Int).Exp(in3, s, params.NExpY)        // (1 + n)^s mod n^y
	in5 := new(big.Int).Mul(in2, in4)                    // h^(r * n^(y - 1)) * (1 + n)^s
	v := new(big.Int).Mod(in5, params.NExpY)             // h^(r * n^(y - 1)) * (1 + n)^s mod n^y

	puzzle := NewPuzzle(u, v)

	return puzzle, nil
}

// SolvePuzzle solves the puzzle and returns the plaintext that was hidden inside
// of it.
func SolvePuzzle(params *params.Params, puzzle *Puzzle) *big.Int {
	// Compute w = u^(2^t) mod n by repeated squaring.
	i := big.NewInt(0)
	w := new(big.Int).Set(puzzle.U) // w = u
	for {
		w = new(big.Int).Exp(w, big.NewInt(2), params.N) // w^2 mod n
		i = new(big.Int).Add(i, big.NewInt(1))           // i + 1
		if i.Cmp(params.T) == 0 {
			break
		}
	}

	// Compute a = (1 + n)^s mod n^y.
	in1 := new(big.Int).Exp(w, params.NExpYMinusOne, params.NExpY) // w^(n^(y - 1)) mod n^y
	in2 := new(big.Int).ModInverse(in1, params.NExpY)              // w^-(n^(y - 1)) mod n^y
	in3 := new(big.Int).Mul(puzzle.V, in2)                         // v * w^-(n^(y - 1))
	a := new(big.Int).Mod(in3, params.NExpY)                       // v * w^-(n^(y - 1)) mod n^y

	// Compute s via the polynomial-time discrete-logarithm algorithm described in
	// section "3 A Generalisation of Paillier’s Probabilistic Encryption Scheme"
	// of the paper https://www.brics.dk/RS/00/45/BRICS-RS-00-45.pdf.
	var s = big.NewInt(0)
	for j := 1; j <= params.Y-1; j++ {
		// Compute n^j and n^(j + 1).
		_, n1, n2 := utils.Exponentiate(params.N, j)

		// Compute t1 = L(a mod n^(j + 1)) = ((1 + n)^s mod n^(j + 1) - 1) / n.
		in1 := new(big.Int).Mod(a, n2)              // a mod n^(j + 1)
		in2 := new(big.Int).Sub(in1, big.NewInt(1)) // a - 1
		t1 := new(big.Int).Div(in2, params.N)       // (a - 1) / n

		t2 := s

		for k := 2; k <= j; k++ {
			s = new(big.Int).Sub(s, big.NewInt(1)) // s - 1

			in1 := new(big.Int).Mul(t2, s) // t2 * s
			t2 = new(big.Int).Mod(in1, n1) // t2 * s mod n^j

			// Compute n^(k - 1).
			k0, _, _ := utils.Exponentiate(params.N, k)

			in2 := new(big.Int).Mul(t2, k0)              // t2 * n^(k - 1)
			in3 := utils.Factorial(big.NewInt(int64(k))) // k!
			in4 := new(big.Int).Div(in2, in3)            // t2 * n^(k - 1) / k!
			in5 := new(big.Int).Sub(t1, in4)             // t1 - t2 * n^(k - 1) / k!
			t1 = new(big.Int).Mod(in5, n1)               // t1 - t2 * n^(k - 1) / k! mod n^j
		}

		s = t1
	}

	return s
}
