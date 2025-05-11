package puzzle

import "fmt"

// ErrSampleNonceR is returned if the random nonce r can't be sampled.
var ErrSampleNonceR = fmt.Errorf("unable to sample random nonce r")
