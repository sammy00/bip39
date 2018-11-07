package bip39

import (
	"crypto/sha256"
	"math/big"
)

func EntropyLenCompatible(n EntropyLen) bool {
	return 0 == n%4 && n >= Size128 && n <= Size256
}

func EntropyToFullBigInt(entropy []byte) (*big.Int, int, error) {
	n := len(entropy)
	if !EntropyLenCompatible(n) {
		return nil, 0, ErrEntropyLen
	}

	checksumB, checksumLen := sha256.Sum256(entropy)[0], n/32
	entropy = append(entropy, checksumB)

	x := new(big.Int).SetBytes(entropy)
	x.Rsh(x, uint(8-checksumLen))

	return x, n, nil
}

/*
func EntropyToWordIndices(entropy []byte) ([]int64, error) {
	n := len(entropy)
	if !EntropyLenCompatible(n) {
		return nil, ErrEntropyLen
	}

	// make up the full entropy as a big int
	checksumB, checksumLen := sha256.Sum256(entropy)[0], n/32
	entropy = append(entropy, checksumB)

	x := new(big.Int).SetBytes(entropy)
	x.Rsh(x, uint(8-checksumLen))

	// MS=(ENT+CS)/11=(ENT+ENT/32)/11=3*ENT/32
	sentenceLen := 3 * n / 32
	indices := make([]int64, sentenceLen)

	wordIndex := new(big.Int)
	for i := len(indices) - 1; i >= 0; i-- {
		wordIndex.And(x, bitMask11)
	}
}
*/

// RecoverEntropy recovers the entropy without the checksum from
// the given mnemonic
func RecoverEntropy(mnemonic Mnemonic) ([]byte, error) {
	return nil, nil
}

// RecoverFullEntropy recovers the entropy plus the checksum from
// the given mnemonic
func RecoverFullEntropy(mnemonic Mnemonic) ([]byte, error) {
	return nil, nil
}

func ValidateFullEntropy(entropy []byte) bool {
	return false
}
