package utils

import (
	"fmt"
	"math/big"
)

// signature struct
type Signature struct {
	S *big.Int // public key S coordiante
	R *big.Int // computed by refering to info like the tx hash
}

func (s *Signature) String() string {
	return fmt.Sprintf("%x%x", s.R, s.S)
}
