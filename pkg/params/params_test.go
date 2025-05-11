package params_test

import (
	"errors"
	"math/big"
	"testing"

	"github.com/primefactor-io/lhtlp/pkg/params"
)

func TestParamsGeneration(t *testing.T) {
	t.Parallel()

	t.Run("Error when prime numbers are equal", func(t *testing.T) {
		t.Parallel()

		_, err := params.GenerateParams(4, 2, big.NewInt(1))

		if !errors.Is(err, params.ErrEqualPrimeNumbers) {
			t.Errorf("want error %v, got %v", params.ErrEqualPrimeNumbers, err)
		}
	})
}
