package homomorphic_test

import (
	"math/big"
	"testing"

	"github.com/primefactor-io/lhtlp/pkg/homomorphic"
	"github.com/primefactor-io/lhtlp/pkg/params"
	"github.com/primefactor-io/lhtlp/pkg/puzzle"
)

func TestMultiplyPlaintextValue(t *testing.T) {
	t.Parallel()

	t.Run("Generate Puzzle / Multiply Plaintext Value / Solve Puzzle", func(t *testing.T) {
		t.Parallel()

		message1 := big.NewInt(24)
		message2 := big.NewInt(42)
		expected := big.NewInt(1_008)

		params, _ := params.GenerateParams(128, 2, big.NewInt(1))
		puzzle1, _ := puzzle.GeneratePuzzle(params, message1)

		puzzle2 := homomorphic.MultiplyPlaintextValue(params, puzzle1, message2)

		result := puzzle.SolvePuzzle(params, puzzle2)

		if result.Cmp(expected) != 0 {
			t.Errorf("want %v, got %v", expected, result)
		}
	})
}
