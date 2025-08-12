package homomorphic_test

import (
	"math/big"
	"testing"

	"github.com/primefactor-io/lhtlp/pkg/homomorphic"
	"github.com/primefactor-io/lhtlp/pkg/params"
	"github.com/primefactor-io/lhtlp/pkg/puzzle"
)

func TestAddPlaintextValues(t *testing.T) {
	t.Parallel()

	t.Run("Generate 2 Puzzles / Add Message Values / Solve Puzzle", func(t *testing.T) {
		t.Parallel()

		message1 := big.NewInt(24)
		message2 := big.NewInt(42)
		expected := big.NewInt(66)

		params, _ := params.GenerateParams(128, 2, big.NewInt(1))
		puzzle1, _ := puzzle.GeneratePuzzle(params, message1)
		puzzle2, _ := puzzle.GeneratePuzzle(params, message2)

		puzzle3 := homomorphic.AddPlaintextValues(params, puzzle1, puzzle2)

		result := puzzle.SolvePuzzle(params, puzzle3)

		if result.Cmp(expected) != 0 {
			t.Errorf("want %v, got %v", expected, result)
		}
	})

	t.Run("Generate 3 Puzzles / Add Message Values / Solve Puzzle", func(t *testing.T) {
		t.Parallel()

		message1 := big.NewInt(24)
		message2 := big.NewInt(42)
		message3 := big.NewInt(11)
		expected := big.NewInt(77)

		params, _ := params.GenerateParams(128, 2, big.NewInt(1))
		puzzle1, _ := puzzle.GeneratePuzzle(params, message1)
		puzzle2, _ := puzzle.GeneratePuzzle(params, message2)
		puzzle3, _ := puzzle.GeneratePuzzle(params, message3)

		puzzle4 := homomorphic.AddPlaintextValues(params, puzzle1, puzzle2, puzzle3)

		result := puzzle.SolvePuzzle(params, puzzle4)

		if result.Cmp(expected) != 0 {
			t.Errorf("want %v, got %v", expected, result)
		}
	})

	t.Run("Generate 10 Puzzles / Add Message Values / Solve Puzzle", func(t *testing.T) {
		t.Parallel()

		message1 := big.NewInt(1)
		message2 := big.NewInt(2)
		message3 := big.NewInt(4)
		message4 := big.NewInt(8)
		message5 := big.NewInt(16)
		message6 := big.NewInt(32)
		message7 := big.NewInt(64)
		message8 := big.NewInt(128)
		message9 := big.NewInt(256)
		message10 := big.NewInt(512)
		expected := big.NewInt(1_023)

		params, _ := params.GenerateParams(128, 2, big.NewInt(1))
		puzzle1, _ := puzzle.GeneratePuzzle(params, message1)
		puzzle2, _ := puzzle.GeneratePuzzle(params, message2)
		puzzle3, _ := puzzle.GeneratePuzzle(params, message3)
		puzzle4, _ := puzzle.GeneratePuzzle(params, message4)
		puzzle5, _ := puzzle.GeneratePuzzle(params, message5)
		puzzle6, _ := puzzle.GeneratePuzzle(params, message6)
		puzzle7, _ := puzzle.GeneratePuzzle(params, message7)
		puzzle8, _ := puzzle.GeneratePuzzle(params, message8)
		puzzle9, _ := puzzle.GeneratePuzzle(params, message9)
		puzzle10, _ := puzzle.GeneratePuzzle(params, message10)

		puzzle11 := homomorphic.AddPlaintextValues(params, puzzle1, puzzle2, puzzle3,
			puzzle4, puzzle5, puzzle6, puzzle7, puzzle8, puzzle9, puzzle10)

		result := puzzle.SolvePuzzle(params, puzzle11)

		if result.Cmp(expected) != 0 {
			t.Errorf("want %v, got %v", expected, result)
		}
	})

	t.Run("Generate 2 Puzzles / Add Message Values / Solve Puzzle - Large Messages", func(t *testing.T) {
		t.Parallel()

		params, _ := params.GenerateParams(128, 4, big.NewInt(1))

		nn := new(big.Int).Mul(params.N, params.N) // n^2

		message1 := new(big.Int).Set(nn) // n^2
		message2 := new(big.Int).Set(nn) // n^2

		expected := new(big.Int).Add(message1, message2) // n^2 + n^2

		puzzle1, _ := puzzle.GeneratePuzzle(params, message1)
		puzzle2, _ := puzzle.GeneratePuzzle(params, message2)

		puzzle3 := homomorphic.AddPlaintextValues(params, puzzle1, puzzle2)

		result := puzzle.SolvePuzzle(params, puzzle3)

		if result.Cmp(expected) != 0 {
			t.Errorf("want %v, got %v", expected, result)
		}
	})
}

func TestAddPlaintextValue(t *testing.T) {
	t.Parallel()

	t.Run("Generate Puzzle / Add Plaintext Value / Solve Puzzle", func(t *testing.T) {
		t.Parallel()

		message1 := big.NewInt(24)
		message2 := big.NewInt(42)
		expected := big.NewInt(66)

		params, _ := params.GenerateParams(128, 2, big.NewInt(1))
		puzzle1, _ := puzzle.GeneratePuzzle(params, message1)

		puzzle2 := homomorphic.AddPlaintextValue(params, puzzle1, message2)

		result := puzzle.SolvePuzzle(params, puzzle2)

		if result.Cmp(expected) != 0 {
			t.Errorf("want %v, got %v", expected, result)
		}
	})
}
