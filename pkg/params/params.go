package params

import (
	"crypto/rand"
	"math/big"
	"sync"

	"github.com/primefactor-io/lhtlp/pkg/utils"
)

// Params is an instance of protocol parameters.
type Params struct {
	// Y is the exponent.
	Y int
	// T is the difficulty.
	T *big.Int
	// N is the product of p and q.
	N *big.Int
	// G is the value g.
	G *big.Int
	// H is the value h.
	H *big.Int
	// NExpY is the value n^y.
	NExpY *big.Int
	// NExpYMinusOne is the value n^(y - 1).
	NExpYMinusOne *big.Int
}

// NewParams creates a new instance of protocol parameters.
func NewParams(y int, t, n, g, h, nExpY, nExpYMinusOne *big.Int) *Params {
	return &Params{
		Y:             y,
		T:             t,
		N:             n,
		G:             g,
		H:             h,
		NExpY:         nExpY,
		NExpYMinusOne: nExpYMinusOne,
	}
}

// GenerateParams generates protocol parameters based on the desired security
// (expressed in bits) and difficulty.
// Returns an error if the generation of the protocol parameters fails.
func GenerateParams(bits, y int, difficulty *big.Int) (*Params, error) {
	// Prime numbers p and q should have roughly the same size.
	primeBits := bits / 2

	// Generate prime numbers p and q.
	var p *big.Int
	var q *big.Int
	errCh := make(chan error, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	// Generate prime p.
	go func() {
		defer wg.Done()

		var err error
		p, err = rand.Prime(rand.Reader, primeBits)
		if err != nil {
			err = ErrGeneratePrimeP
		}

		errCh <- err
	}()

	// Generate prime q.
	go func() {
		defer wg.Done()

		var err error
		q, err = rand.Prime(rand.Reader, primeBits)
		if err != nil {
			err = ErrGeneratePrimeQ
		}

		errCh <- err
	}()

	wg.Wait()

	if err := <-errCh; err != nil {
		return nil, err
	}

	// Check if prime numbers are equal.
	if p.Cmp(q) == 0 {
		return nil, ErrEqualPrimeNumbers
	}

	t := difficulty
	n := new(big.Int).Mul(p, q)                     // p * q
	nMinusOne := new(big.Int).Sub(n, big.NewInt(1)) // n - 1

	// Compute n^(y - 1) and n^y.
	nExpYMinusOne, nExpY, _ := utils.Exponentiate(n, y)

	pMinusOne := new(big.Int).Sub(p, big.NewInt(1)) // p - 1
	qMinusOne := new(big.Int).Sub(q, big.NewInt(1)) // q - 1

	phiN := new(big.Int).Mul(pMinusOne, qMinusOne)    // (p - 1) * (q - 1)
	phiNHalf := new(big.Int).Div(phiN, big.NewInt(2)) // phiN / 2

	// Randomly sample g'.
	gPrime, err := rand.Int(rand.Reader, nMinusOne)
	if err != nil {
		return nil, ErrSampleGPrime
	}

	// Compute g.
	in1 := new(big.Int).Exp(gPrime, big.NewInt(2), n) // g'^2 mod n
	g := new(big.Int).ModInverse(in1, n)              // -g'^2 mod n

	// Compute h'.
	hPrime := new(big.Int).Exp(big.NewInt(2), t, phiNHalf) // 2^t mod (phiN / 2)

	// Compute h.
	h := new(big.Int).Exp(g, hPrime, n) // g^(2^t) mod n

	params := NewParams(y, t, n, g, h, nExpY, nExpYMinusOne)

	return params, nil
}
