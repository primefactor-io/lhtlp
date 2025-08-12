package utils_test

import (
	"math/big"
	"slices"
	"testing"

	"github.com/primefactor-io/lhtlp/pkg/utils"
)

func TestUtils(t *testing.T) {
	t.Parallel()

	t.Run("GenerateRandomBytesSeeded", func(t *testing.T) {
		t.Parallel()

		bits := 128
		seed := []byte{0, 1, 2, 3, 4, 5}
		bytes, _ := utils.GenerateRandomBytesSeeded(seed, bits)

		if len(bytes)*8 != bits {
			t.Error("Random byte generation failed")
		}

		got := bytes
		want := []byte{
			213, 12, 38, 228, 46, 102, 162, 154,
			222, 213, 38, 53, 39, 181, 57, 114,
		}

		if !slices.Equal(got, want) {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("BytesToBit", func(t *testing.T) {
		t.Parallel()

		bytes := []byte{
			0, // 00000000
			1, // 10000000
			2, // 01000000
			3, // 11000000
			4, // 00100000
			5, // 10100000
			6, // 01100000
			7, // 11100000
			8, // 00010000
		}

		bits := []byte{
			0, 0, 0, 0, 0, 0, 0, 0, // 0
			1, 0, 0, 0, 0, 0, 0, 0, // 1
			0, 1, 0, 0, 0, 0, 0, 0, // 2
			1, 1, 0, 0, 0, 0, 0, 0, // 3
			0, 0, 1, 0, 0, 0, 0, 0, // 4
			1, 0, 1, 0, 0, 0, 0, 0, // 5
			0, 1, 1, 0, 0, 0, 0, 0, // 6
			1, 1, 1, 0, 0, 0, 0, 0, // 7
			0, 0, 0, 1, 0, 0, 0, 0, // 8
		}

		for i := range len(bytes) * 8 {
			want := bits[i]
			got := utils.BytesToBit(bytes, i)

			if got != want {
				t.Errorf("%v want %v, got %v", i+1, want, got)
			}
		}
	})

	t.Run("Factorial", func(t *testing.T) {
		t.Parallel()

		// 0! = 1
		zeroFac := utils.Factorial(big.NewInt(0))
		if zeroFac.Cmp(big.NewInt(1)) != 0 {
			t.Errorf("want 0! = 1, got %v", zeroFac)
		}

		// 1! = 1
		oneFac := utils.Factorial(big.NewInt(1))
		if oneFac.Cmp(big.NewInt(1)) != 0 {
			t.Errorf("want 1! = 1, got %v", oneFac)
		}

		// 2! = 2
		twoFac := utils.Factorial(big.NewInt(2))
		if twoFac.Cmp(big.NewInt(2)) != 0 {
			t.Errorf("want 2! = 2, got %v", oneFac)
		}

		// 3! = 6
		threeFac := utils.Factorial(big.NewInt(3))
		if threeFac.Cmp(big.NewInt(6)) != 0 {
			t.Errorf("want 3! = 6, got %v", oneFac)
		}

		// 20! = 2_432_902_008_176_640_000
		twentyFac := utils.Factorial(big.NewInt(20))
		if twentyFac.Cmp(big.NewInt(2_432_902_008_176_640_000)) != 0 {
			t.Errorf("want 20! = 2_432_902_008_176_640_000, got %v", twentyFac)
		}
	})

	t.Run("Exponentiate", func(t *testing.T) {
		t.Parallel()

		n := big.NewInt(2)
		x := 32

		res1, res2, res3 := utils.Exponentiate(n, x)

		// 2^31 = 2147483648
		if res1.Cmp(big.NewInt(2147483648)) != 0 {
			t.Errorf("want 2^31 = 2147483648, got %v", res1)
		}

		// 2^32 = 4294967296
		if res2.Cmp(big.NewInt(4294967296)) != 0 {
			t.Errorf("want 2^32 = 4294967296, got %v", res2)
		}

		// 2^33 = 8589934592
		if res3.Cmp(big.NewInt(8589934592)) != 0 {
			t.Errorf("want 2^33 = 8589934592, got %v", res3)
		}
	})
}
