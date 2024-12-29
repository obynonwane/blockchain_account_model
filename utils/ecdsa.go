package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"math/big"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

func (s *Signature) String() string {
	// return 64 bytes hexadecimal string concatenation of R & S
	return fmt.Sprintf("%064x%064x", s.R, s.S)
}

func StringToBigIntTuple(s string) (big.Int, big.Int) {

	// returns the byte representation of the hexadecimal string
	bx, _ := hex.DecodeString(s[:64]) // Decode the first 64 characters
	by, _ := hex.DecodeString(s[64:]) //  Decode the next 64 characters

	var bix big.Int
	var biy big.Int

	_ = bix.SetBytes(bx) // Convert bx to big.Int
	_ = biy.SetBytes(by) // Convert by to big.Int

	return bix, biy
}

func SignatureFromString(s string) *Signature {
	x, y := StringToBigIntTuple(s)
	return &Signature{&x, &y}
}

/* converts public key from hexadecimal string to a pointer to ecdsa.PublicKey */
func PublicKeyFromString(s string) *ecdsa.PublicKey {
	x, y := StringToBigIntTuple(s)

	return &ecdsa.PublicKey{elliptic.P256(), &x, &y}
}

// converts hexadecimal string representation of the private to to a pointer to ecdsa.PrivateKey
func PrivateKeyFromString(s string, publicKey *ecdsa.PublicKey) *ecdsa.PrivateKey {

	// returns bytes when converting the hexadecimal string
	b, _ := hex.DecodeString(s[:])

	var bi big.Int
	_ = bi.SetBytes(b)

	return &ecdsa.PrivateKey{PublicKey: *publicKey, D: &bi}
}
