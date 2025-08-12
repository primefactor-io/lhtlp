package proofs

import (
	"crypto/rand"
	"math/big"

	"github.com/primefactor-io/lhtlp/pkg/params"
	"github.com/primefactor-io/lhtlp/pkg/puzzle"
	"github.com/primefactor-io/lhtlp/pkg/utils"
)

// Range proof is an instance of a Range proof.
type RangeProof struct {
	// D is the array that contains all puzzles.
	D []*puzzle.Puzzle
	// Values is the array that contains the individual puzzle values.
	Values []*PuzzleValues
}

// NewRangeProof creates a new instance of a Range proof.
func NewRangeProof(d []*puzzle.Puzzle, values []*PuzzleValues) *RangeProof {
	return &RangeProof{
		D:      d,
		Values: values,
	}
}

// PuzzleValues is an instance of a puzzle's values.
type PuzzleValues struct {
	// X is the puzzle's underlying plaintext value.
	X *big.Int
	// R is the puzzle's nonce.
	R *big.Int
}

// NewPuzzleValues creates a new instance of a puzzle's values.
func NewPuzzleValues(x, r *big.Int) *PuzzleValues {
	return &PuzzleValues{
		X: x,
		R: r,
	}
}

// GenerateRangeProof generates a Range proof which proves that all the puzzle's
// plaintext values (their x values) are an element of {0, ..., q} and in the
// range [-(q / 2), (q / 2)].
// Returns an error if the proof generation fails.
func GenerateRangeProof(bits int, params *params.Params, z []*puzzle.Puzzle, q *big.Int, wit []*PuzzleValues) (*RangeProof, error) {
	k := bits
	numPuzzles := len(z)

	if len(wit) != len(z) {
		return nil, ErrNumPuzzlesAndWitnesses
	}

	l := new(big.Int).SetInt64(int64(numPuzzles)) // l
	l4 := new(big.Int).Mul(big.NewInt(4), l)      // 4 * l
	b := new(big.Int).Div(q, big.NewInt(2))       // q / 2
	m := new(big.Int).Mul(b, l4)                  // (q / 2) * 4 * l = L
	n := new(big.Int).Div(m, big.NewInt(4))       // L / 4
	n2 := new(big.Int).Mul(big.NewInt(2), n)      // 2 * (L / 4)

	// Compute puzzles with drowning terms.
	y := make([]*big.Int, k)
	rPrime := make([]*big.Int, k)
	d := make([]*puzzle.Puzzle, k)

	for i := range k {
		// Sample random drowning term y_i in [0, 2 * (L / 4)).
		yi, err := rand.Int(rand.Reader, n2)
		if err != nil {
			return nil, ErrSampleY
		}

		// Compute D_i and r_i'.
		di, riPrime, err := puzzle.GeneratePuzzleAndReturnNonce(params, yi)
		if err != nil {
			return nil, ErrComputeD
		}

		d[i] = di
		y[i] = yi
		rPrime[i] = riPrime
	}

	// Compute puzzle values v and w.
	values := make([]*PuzzleValues, k)

	// Generate randomness via Fiat-Shamir transform.
	t, err := proofDataToHashBytes(k, numPuzzles, z, d)
	if err != nil {
		return nil, ErrGenerateRandomness
	}

	for i := range k {
		xjSum := big.NewInt(0)
		rjSum := big.NewInt(0)

		for j := range numPuzzles {
			index := (i * numPuzzles) + j
			bit := utils.BytesToBit(t, index)
			switch bit {
			case 0:
				continue
			case 1:
				xj := wit[j].X
				xjSum = new(big.Int).Add(xjSum, xj) // x_{j-1} + x_j

				rj := wit[j].R
				rjSum = new(big.Int).Add(rjSum, rj) // r_{j-1} + r_j
			default:
				// Bit value is neither 0 nor 1.
				return nil, ErrInvalidBit
			}
		}

		yi := y[i]
		vi := new(big.Int).Add(yi, xjSum)

		riPrime := rPrime[i]
		wi := new(big.Int).Add(riPrime, rjSum)

		values[i] = NewPuzzleValues(vi, wi)
	}

	proof := NewRangeProof(d, values)

	return proof, nil
}

// VerifyRangePoof verifies a Range proof which proves that all the puzzle's
// plaintext values (their x values) are an element of {0, ..., q} and in the
// range [-(q / 2), (q / 2)].
// Returns an error if the proof verification fails.
func VerifyRangePoof(proof *RangeProof, bits int, params *params.Params, z []*puzzle.Puzzle, q *big.Int) (bool, error) {
	k := bits
	numPuzzles := len(z)

	if len(proof.D) != len(proof.Values) {
		return false, ErrNumPuzzlesAndValues
	}

	zero := big.NewInt(0)
	l := new(big.Int).SetInt64(int64(numPuzzles)) // l
	l4 := new(big.Int).Mul(big.NewInt(4), l)      // 4 * l
	b := new(big.Int).Div(q, big.NewInt(2))       // q / 2
	m := new(big.Int).Mul(b, l4)                  // (q / 2) * 4 * l = L
	n := new(big.Int).Div(m, big.NewInt(2))       // L / 2
	n2 := new(big.Int).Mul(big.NewInt(2), n)      // 2 * (L / 2)

	// (Re)Generate randomness via Fiat-Shamir transform.
	t, err := proofDataToHashBytes(k, numPuzzles, z, proof.D)
	if err != nil {
		return false, ErrGenerateRandomness
	}

	for i := range k {
		vi := proof.Values[i].X
		wi := proof.Values[i].R

		// Check if v_i is an element of {0, ..., 2 * (L / 2)}.
		isViInSet := vi.Cmp(zero) >= 0 && vi.Cmp(n2) <= 0
		if !isViInSet {
			return false, nil
		}

		// Compute F_i.
		zjuProduct := big.NewInt(1)
		zjvProduct := big.NewInt(1)

		for j := range numPuzzles {
			index := (i * numPuzzles) + j
			bit := utils.BytesToBit(t, index)
			switch bit {
			case 0:
				continue
			case 1:
				zju := z[j].U
				in1 := new(big.Int).Mul(zjuProduct, zju)     // Z_{j-1}.u * Z_j.u
				zjuProduct = new(big.Int).Mod(in1, params.N) // Z_{j-1}.u * Z_j.u mod n

				zjv := z[j].V
				in2 := new(big.Int).Mul(zjvProduct, zjv)         // Z_{j-1}.v * Z_j.v
				zjvProduct = new(big.Int).Mod(in2, params.NExpY) // Z_{j-1}.v * Z_j.v mod n^y
			default:
				// Bit value is neither 0 nor 1.
				return false, ErrInvalidBit
			}
		}

		diu := proof.D[i].U
		in1 := new(big.Int).Mul(diu, zjuProduct) // D_i.U * (... * Z_{j-1}.u) * Z_j.u)
		fiu := new(big.Int).Mod(in1, params.N)   // D_i.U * (... * Z_{j-1}.u) * Z_j.u) mod n

		div := proof.D[i].V
		in2 := new(big.Int).Mul(div, zjvProduct)   // D_i.V * (... * Z_{j-1}.v) * Z_j.v
		fiv := new(big.Int).Mod(in2, params.NExpY) // D_i.V * (... * Z_{j-1}.v) * Z_j.v mod n^y

		fi := puzzle.NewPuzzle(fiu, fiv)

		fiPrime, err := puzzle.GeneratePuzzleWithCustomNonce(params, wi, vi)
		if err != nil {
			return false, ErrComputeFiPrime
		}

		// Check if puzzles are equal.
		if !fi.Equal(fiPrime) {
			return false, nil
		}
	}

	return true, nil
}

// proofDataToHashBytes implements the Fiat-Shamir transform to derive an array
// of bytes.
// Returns an error if random bytes can't be derived from the proof data.
func proofDataToHashBytes(k, l int, z []*puzzle.Puzzle, d []*puzzle.Puzzle) ([]byte, error) {
	var seed []byte

	// Puzzles (Z).
	for _, puzzle := range z {
		seed = append(seed, puzzle.U.Bytes()...)
		seed = append(seed, puzzle.V.Bytes()...)
	}
	// Puzzles (D).
	for _, puzzle := range d {
		seed = append(seed, puzzle.U.Bytes()...)
		seed = append(seed, puzzle.V.Bytes()...)
	}

	numBits := k * l
	randBytes, err := utils.GenerateRandomBytesSeeded(seed, numBits)
	if err != nil {
		return nil, ErrGenerateRandomBytes
	}

	return randBytes, nil
}
