package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"math/big"
)

// GenerateRandomBytesSeeded uses a seed to generate a byte slice that contains
// the number of desired bits.
// Returns an error if the random bytes can't be generated.
func GenerateRandomBytesSeeded(seed []byte, bits int) ([]byte, error) {
	// Derive AES key from seed.
	h := sha256.New()
	h.Write(seed)
	key := h.Sum(nil)

	// Initialize AES.
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, ErrInitializeAES
	}

	// Set IV to 16 zero bytes.
	iv := make([]byte, aes.BlockSize)
	// Initialize stream cipher via CTR mode.
	stream := cipher.NewCTR(block, iv)

	// Use keystream to derive random bits.
	numBytes := (bits + 7) / 8
	randBytes := make([]byte, numBytes)
	stream.XORKeyStream(randBytes, randBytes)

	// Remove excess bits if length of random bits is too large.
	mask := uint32(1<<uint(bits%8)) - 1
	if mask > 0 {
		randBytes[0] &= byte(mask)
	}

	return randBytes, nil
}

// BytesToBit returns the bit at position i within the passed-in byte slice.
// The passed-in byte slice is interpreted as a continuous stream of bits.
func BytesToBit(bytes []byte, i int) byte {
	selector := i / 8
	byt := bytes[selector]

	bit := i % 8

	mask := byte(0x01)

	return (byt >> bit) & mask
}

// Factorial computes the factorial of x which is x! = 1 * 2 * 3 * ... * x.
// See: https://stackoverflow.com/a/54952509
func Factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)

	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}

	return n.Mul(x, Factorial(n.Sub(x, n)))
}

// Exponentiate computes the exponentiations n^(x - 1), n^x and n^(x + 1).
// Note: The caller needs to ensure that x has a value in the correct range
// (e.g. that n^(x - 1) can be computed correctly).
func Exponentiate(n *big.Int, x int) (*big.Int, *big.Int, *big.Int) {
	var n1 = big.NewInt(1) // n^(x - 1)
	var n2 = big.NewInt(1) // n^x
	var n3 = big.NewInt(1) // n^(x + 1)

	for i := 1; i <= x+1; i++ {
		if i <= x-1 {
			n1 = new(big.Int).Mul(n1, n)
		}

		if i <= x {
			n2 = new(big.Int).Mul(n2, n)
		}

		if i <= x+1 {
			n3 = new(big.Int).Mul(n3, n)
		}
	}

	return n1, n2, n3
}
