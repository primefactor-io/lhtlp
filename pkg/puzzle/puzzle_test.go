package puzzle_test

import (
	"math/big"
	"testing"

	"github.com/primefactor-io/lhtlp/pkg/params"
	"github.com/primefactor-io/lhtlp/pkg/puzzle"
)

func TestPuzzle(t *testing.T) {
	t.Parallel()

	t.Run("Generate Puzzle / Solve Puzzle", func(t *testing.T) {
		t.Parallel()

		message := big.NewInt(42)

		params, _ := params.GenerateParams(128, 2, big.NewInt(1))
		puzzle1, _ := puzzle.GeneratePuzzle(params, message)

		mPrime := puzzle.SolvePuzzle(params, puzzle1)

		if mPrime.Cmp(message) != 0 {
			t.Errorf("want %v, got %v", message, mPrime)
		}
	})

	t.Run("Generate Puzzle / Solve Puzzle - Large Message", func(t *testing.T) {
		t.Parallel()

		params, _ := params.GenerateParams(128, 4, big.NewInt(1))

		message := new(big.Int).Mul(params.N, params.N) // n^2
		puzzle1, _ := puzzle.GeneratePuzzle(params, message)

		mPrime := puzzle.SolvePuzzle(params, puzzle1)

		if mPrime.Cmp(message) != 0 {
			t.Errorf("want %v, got %v", message, mPrime)
		}
	})

	t.Run("Generate Puzzle / Solve Puzzle - Custom Nonce", func(t *testing.T) {
		t.Parallel()

		nonce := big.NewInt(11)
		message := big.NewInt(42)

		params, _ := params.GenerateParams(128, 2, big.NewInt(1))
		puzzle1, _ := puzzle.GeneratePuzzleWithCustomNonce(params, nonce, message)

		mPrime := puzzle.SolvePuzzle(params, puzzle1)

		if mPrime.Cmp(message) != 0 {
			t.Errorf("want %v, got %v", message, mPrime)
		}
	})

	t.Run("Generate Puzzle / Solve Puzzle - Return Nonce", func(t *testing.T) {
		t.Parallel()

		message := big.NewInt(42)

		params, _ := params.GenerateParams(128, 2, big.NewInt(1))
		puzzle1, _, _ := puzzle.GeneratePuzzleAndReturnNonce(params, message)

		mPrime := puzzle.SolvePuzzle(params, puzzle1)

		if mPrime.Cmp(message) != 0 {
			t.Errorf("want %v, got %v", message, mPrime)
		}
	})

	t.Run("Puzzle Equality", func(t *testing.T) {
		t.Parallel()

		message := big.NewInt(42)

		params, _ := params.GenerateParams(128, 2, big.NewInt(1))

		p1, nonce, _ := puzzle.GeneratePuzzleAndReturnNonce(params, message)
		p2, _ := puzzle.GeneratePuzzleWithCustomNonce(params, nonce, message)

		if p1.Equal(p2) != true {
			t.Errorf("puzzles are not equal %v %v", p1, p2)
		}
	})
}
