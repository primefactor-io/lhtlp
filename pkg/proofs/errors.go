package proofs

import "fmt"

var (
	// ErrNumPuzzlesAndWitnesses is returned if the number of puzzles is not equal to the number of witnesses.
	ErrNumPuzzlesAndWitnesses = fmt.Errorf("number of puzzles is not equal to number of witnesses")
	// ErrSampleY is returned if the drowning term y can't be sampled.
	ErrSampleY = fmt.Errorf("unable to sample drowning term y")
	// ErrComputeD is returned if D can't be computed.
	ErrComputeD = fmt.Errorf("unable to compute D")
	// ErrGenerateRandomness is returned if the randomness can't be generated.
	ErrGenerateRandomness = fmt.Errorf("unable to generate randomness")
	// ErrInvalidBit is returned if the bit is neither 0 nor 1.
	ErrInvalidBit = fmt.Errorf("bit is neither 0 nor 1")
	// ErrNumPuzzlesAndValues is returned if the number of puzzles is not equal to the number of values.
	ErrNumPuzzlesAndValues = fmt.Errorf("number of puzzles is not equal to number of values")
	// ErrComputeFiPrime is returned if Fi' can't be computed.
	ErrComputeFiPrime = fmt.Errorf("unable to compute Fi'")
	// ErrGenerateRandomBytes is returned if the random bytes can't be generated.
	ErrGenerateRandomBytes = fmt.Errorf("unable to generate random bytes")
)
