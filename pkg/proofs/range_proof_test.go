package proofs_test

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/primefactor-io/lhtlp/pkg/params"
	"github.com/primefactor-io/lhtlp/pkg/proofs"
	"github.com/primefactor-io/lhtlp/pkg/puzzle"
)

func TestRangeProof(t *testing.T) {
	t.Parallel()

	t.Run("Prove / Verify - Single Puzzle - Valid (m = 0)", func(t *testing.T) {
		t.Parallel()

		bits := 128
		q := big.NewInt(1000)

		m := big.NewInt(0)

		params, _ := params.GenerateParams(bits, 2, big.NewInt(1))
		p, r, _ := puzzle.GeneratePuzzleAndReturnNonce(params, m)
		v := proofs.NewPuzzleValues(m, r)

		puzzles := []*puzzle.Puzzle{p}
		values := []*proofs.PuzzleValues{v}

		proof, _ := proofs.GenerateRangeProof(bits, params, puzzles, q, values)
		isValid, _ := proofs.VerifyRangePoof(proof, bits, params, puzzles, q)

		if isValid != true {
			t.Error("Range proof verification failed")
		}
	})

	t.Run("Prove / Verify - Single Puzzle - Valid (m >= 0 and m < q)", func(t *testing.T) {
		t.Parallel()

		bits := 128
		q := big.NewInt(1000)

		m, _ := rand.Int(rand.Reader, q)

		params, _ := params.GenerateParams(bits, 2, big.NewInt(1))
		p, r, _ := puzzle.GeneratePuzzleAndReturnNonce(params, m)
		v := proofs.NewPuzzleValues(m, r)

		puzzles := []*puzzle.Puzzle{p}
		values := []*proofs.PuzzleValues{v}

		proof, _ := proofs.GenerateRangeProof(bits, params, puzzles, q, values)
		isValid, _ := proofs.VerifyRangePoof(proof, bits, params, puzzles, q)

		if isValid != true {
			t.Error("Range proof verification failed")
		}
	})

	t.Run("Prove / Verify - Single Puzzle - Valid (m = q)", func(t *testing.T) {
		t.Parallel()

		bits := 128
		q := big.NewInt(1000)

		m := q

		params, _ := params.GenerateParams(bits, 2, big.NewInt(1))
		p, r, _ := puzzle.GeneratePuzzleAndReturnNonce(params, m)
		v := proofs.NewPuzzleValues(m, r)

		puzzles := []*puzzle.Puzzle{p}
		values := []*proofs.PuzzleValues{v}

		proof, _ := proofs.GenerateRangeProof(bits, params, puzzles, q, values)
		isValid, _ := proofs.VerifyRangePoof(proof, bits, params, puzzles, q)

		if isValid != true {
			t.Error("Range proof verification failed")
		}
	})

	t.Run("Prove / Verify - Single Puzzle - Invalid (m < 0)", func(t *testing.T) {
		t.Parallel()

		bits := 128
		q := big.NewInt(1000)
		q2 := new(big.Int).Mul(big.NewInt(2), q) // 2 * q

		m := new(big.Int).Neg(q2) // -(2 * q)

		params, _ := params.GenerateParams(bits, 2, big.NewInt(1))
		p, r, _ := puzzle.GeneratePuzzleAndReturnNonce(params, m)
		v := proofs.NewPuzzleValues(m, r)

		puzzles := []*puzzle.Puzzle{p}
		values := []*proofs.PuzzleValues{v}

		proof, _ := proofs.GenerateRangeProof(bits, params, puzzles, q, values)
		isValid, _ := proofs.VerifyRangePoof(proof, bits, params, puzzles, q)

		if isValid != false {
			t.Error("Range proof verification failed")
		}
	})

	t.Run("Prove / Verify - Single Puzzle - Invalid (m > q)", func(t *testing.T) {
		t.Parallel()

		bits := 128
		q := big.NewInt(1000)
		q2 := new(big.Int).Mul(big.NewInt(2), q) // 2 * q

		m := q2

		params, _ := params.GenerateParams(bits, 2, big.NewInt(1))
		p, r, _ := puzzle.GeneratePuzzleAndReturnNonce(params, m)
		v := proofs.NewPuzzleValues(m, r)

		puzzles := []*puzzle.Puzzle{p}
		values := []*proofs.PuzzleValues{v}

		proof, _ := proofs.GenerateRangeProof(bits, params, puzzles, q, values)
		isValid, _ := proofs.VerifyRangePoof(proof, bits, params, puzzles, q)

		if isValid != false {
			t.Error("Range proof verification failed")
		}
	})

	t.Run("Prove / Verify - Multiple Puzzles - Valid (m = 0)", func(t *testing.T) {
		t.Parallel()

		bits := 128
		q := big.NewInt(1000)

		m1 := big.NewInt(0)
		m2 := big.NewInt(0)
		m3 := big.NewInt(0)

		params, _ := params.GenerateParams(bits, 2, big.NewInt(1))

		p1, r1, _ := puzzle.GeneratePuzzleAndReturnNonce(params, m1)
		p2, r2, _ := puzzle.GeneratePuzzleAndReturnNonce(params, m2)
		p3, r3, _ := puzzle.GeneratePuzzleAndReturnNonce(params, m3)
		v1 := proofs.NewPuzzleValues(m1, r1)
		v2 := proofs.NewPuzzleValues(m2, r2)
		v3 := proofs.NewPuzzleValues(m3, r3)

		puzzles := []*puzzle.Puzzle{p1, p2, p3}
		values := []*proofs.PuzzleValues{v1, v2, v3}

		proof, _ := proofs.GenerateRangeProof(bits, params, puzzles, q, values)
		isValid, _ := proofs.VerifyRangePoof(proof, bits, params, puzzles, q)

		if isValid != true {
			t.Error("Range proof verification failed")
		}
	})

	t.Run("Prove / Verify - Multiple Puzzles - Valid (m >= 0 and m < q)", func(t *testing.T) {
		t.Parallel()

		bits := 128
		q := big.NewInt(1000)

		m1, _ := rand.Int(rand.Reader, q)
		m2, _ := rand.Int(rand.Reader, q)
		m3, _ := rand.Int(rand.Reader, q)

		params, _ := params.GenerateParams(bits, 2, big.NewInt(1))

		p1, r1, _ := puzzle.GeneratePuzzleAndReturnNonce(params, m1)
		p2, r2, _ := puzzle.GeneratePuzzleAndReturnNonce(params, m2)
		p3, r3, _ := puzzle.GeneratePuzzleAndReturnNonce(params, m3)
		v1 := proofs.NewPuzzleValues(m1, r1)
		v2 := proofs.NewPuzzleValues(m2, r2)
		v3 := proofs.NewPuzzleValues(m3, r3)

		puzzles := []*puzzle.Puzzle{p1, p2, p3}
		values := []*proofs.PuzzleValues{v1, v2, v3}

		proof, _ := proofs.GenerateRangeProof(bits, params, puzzles, q, values)
		isValid, _ := proofs.VerifyRangePoof(proof, bits, params, puzzles, q)

		if isValid != true {
			t.Error("Range proof verification failed")
		}
	})

	t.Run("Prove / Verify - Multiple Puzzles - Valid (m = q)", func(t *testing.T) {
		t.Parallel()

		bits := 128
		q := big.NewInt(1000)

		m1 := q
		m2 := q
		m3 := q

		params, _ := params.GenerateParams(bits, 2, big.NewInt(1))

		p1, r1, _ := puzzle.GeneratePuzzleAndReturnNonce(params, m1)
		p2, r2, _ := puzzle.GeneratePuzzleAndReturnNonce(params, m2)
		p3, r3, _ := puzzle.GeneratePuzzleAndReturnNonce(params, m3)
		v1 := proofs.NewPuzzleValues(m1, r1)
		v2 := proofs.NewPuzzleValues(m2, r2)
		v3 := proofs.NewPuzzleValues(m3, r3)

		puzzles := []*puzzle.Puzzle{p1, p2, p3}
		values := []*proofs.PuzzleValues{v1, v2, v3}

		proof, _ := proofs.GenerateRangeProof(bits, params, puzzles, q, values)
		isValid, _ := proofs.VerifyRangePoof(proof, bits, params, puzzles, q)

		if isValid != true {
			t.Error("Range proof verification failed")
		}
	})

	t.Run("Prove / Verify - Multiple Puzzles - Invalid (m < 0)", func(t *testing.T) {
		t.Parallel()

		bits := 128
		q := big.NewInt(1000)
		q2 := new(big.Int).Mul(big.NewInt(2), q) // 2 * q

		m1 := new(big.Int).Neg(q2) // -(2 * q)
		m2 := new(big.Int).Neg(q2) // -(2 * q)
		m3 := new(big.Int).Neg(q2) // -(2 * q)

		params, _ := params.GenerateParams(bits, 2, big.NewInt(1))

		p1, r1, _ := puzzle.GeneratePuzzleAndReturnNonce(params, m1)
		p2, r2, _ := puzzle.GeneratePuzzleAndReturnNonce(params, m2)
		p3, r3, _ := puzzle.GeneratePuzzleAndReturnNonce(params, m3)
		v1 := proofs.NewPuzzleValues(m1, r1)
		v2 := proofs.NewPuzzleValues(m2, r2)
		v3 := proofs.NewPuzzleValues(m3, r3)

		puzzles := []*puzzle.Puzzle{p1, p2, p3}
		values := []*proofs.PuzzleValues{v1, v2, v3}

		proof, _ := proofs.GenerateRangeProof(bits, params, puzzles, q, values)
		isValid, _ := proofs.VerifyRangePoof(proof, bits, params, puzzles, q)

		if isValid != false {
			t.Error("Range proof verification failed")
		}
	})

	t.Run("Prove / Verify - Multiple Puzzles - Invalid (m > q)", func(t *testing.T) {
		t.Parallel()

		bits := 128
		q := big.NewInt(1000)
		q2 := new(big.Int).Mul(big.NewInt(2), q) // 2 * q

		m1 := q2
		m2 := q2
		m3 := q2

		params, _ := params.GenerateParams(bits, 2, big.NewInt(1))

		p1, r1, _ := puzzle.GeneratePuzzleAndReturnNonce(params, m1)
		p2, r2, _ := puzzle.GeneratePuzzleAndReturnNonce(params, m2)
		p3, r3, _ := puzzle.GeneratePuzzleAndReturnNonce(params, m3)
		v1 := proofs.NewPuzzleValues(m1, r1)
		v2 := proofs.NewPuzzleValues(m2, r2)
		v3 := proofs.NewPuzzleValues(m3, r3)

		puzzles := []*puzzle.Puzzle{p1, p2, p3}
		values := []*proofs.PuzzleValues{v1, v2, v3}

		proof, _ := proofs.GenerateRangeProof(bits, params, puzzles, q, values)
		isValid, _ := proofs.VerifyRangePoof(proof, bits, params, puzzles, q)

		if isValid != false {
			t.Error("Range proof verification failed")
		}
	})
}
