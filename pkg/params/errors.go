package params

import "fmt"

var (
	// ErrGeneratePrimeP is returned if the prime p can't be generated.
	ErrGeneratePrimeP = fmt.Errorf("unable to generate prime p")
	// ErrGeneratePrimeQ is returned if the prime q can't be generated.
	ErrGeneratePrimeQ = fmt.Errorf("unable to generate prime q")
	// ErrEqualPrimeNumbers is returned if the prime numbers are equal.
	ErrEqualPrimeNumbers = fmt.Errorf("equal prime numbers")
	// ErrSampleGPrime is returned if the random g' can't be sampled.
	ErrSampleGPrime = fmt.Errorf("unable to sample random g'")
)
