package bip39

import (
	"crypto/sha256"
	"errors"
	"math/big"
	"strings"

	"golang.org/x/text/unicode/norm"

	"github.com/sammy00/bip39/dict"
)

func EntropyLenCompatible(n EntropyLen) bool {
	return 0 == n%4 && n >= Size128 && n <= Size256
}

/*
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
*/

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
	// recover the full entropy by RecoverFullEntropy()
	// trim out checksum
	return nil, nil
}

// RecoverFullEntropy recovers the entropy plus the checksum from
// the given mnemonic
func RecoverFullEntropy(mnemonic Mnemonic, lang ...dict.Language) (
	[]byte, error) {
	// vanity check against language
	language, err := dict.LanguageToUse(lang...)
	if nil != err {
		return nil, err
	}

	// split mnemonic into words
	words := strings.Split(norm.NFKD.String(mnemonic), " ")
	// calculate the expected entropy length and checksum length
	// 	ENT/8=MS*4/3, CS=(ENT/8)/4
	nWord := len(words)
	n := nWord * 4 / 3
	// check compatibility of entropy length
	if nWord%3 != 0 || !EntropyLenCompatible(n) {
		return nil, errors.New("invalid mnemonic length")
	}

	// recover the full entropy and delegates further
	x, y := new(big.Int), new(big.Int)
	for _, word := range words {
		idx, ok := dict.LookUp(word, language)
		if !ok {
			return nil, errors.New("missing word")
		}

		y.SetInt64(int64(idx))
		x.Lsh(x, WordIndexBitSize)
		x.Or(x, y)
	}

	entropy := paddedPrepend(nil, x.Bytes(), n+1)
	//  checking to ValidateFullEntropy()
	//if err := ValidateFullEntropy(entropy); nil != err {
	//	return nil, errors.New("invalid full entropy")
	//}
	//fmt.Printf("world=%x\n", entropy)
	//entropy, err := recoverFullEntropy(mnemonic, lang...)
	//if nil != err {
	//	return nil, err
	//}

	if _, _, err := DecodeFullEntropy(entropy); nil != err {
		return nil, err
	}

	return entropy, nil
}

// RecoverNativeFullEntropy recovers the entropy plus the checksum from
// the given mnemonic
//func RecoverNativeFullEntropy(mnemonic Mnemonic, lang dict.Language) (
//	[]byte, error) {
//
//	return nil, nil
//}

func DecodeFullEntropy(data []byte) ([]byte, byte, error) {
	n := len(data) - 1
	if !EntropyLenCompatible(n) {
		return nil, 0, ErrEntropyLen
	}

	// k1 is the length of checksum
	k1, k2 := uint(n/4), uint(8-n/4)

	// decode checksum
	//checksumMask := byte(1<<uint8(k1) - 1)
	//checksum := data[n] & checksumMask
	checksum := data[n] & byte(1<<uint8(k1)-1)

	// decode raw entropy: shift all bytes right by checksumLen bits
	entropy := make([]byte, n)
	for i := range entropy {
		entropy[i] = (data[i] << k2) | (data[i+1] >> k1)
	}

	expectedChecksum := sha256.Sum256(entropy)[0] >> k2
	if expectedChecksum != checksum {
		return nil, 0, ErrChecksum
	}

	return entropy, checksum, nil
}

func recoverFullEntropy(mnemonic Mnemonic, lang ...dict.Language) (
	[]byte, error) {
	// vanity check against language
	language, err := dict.LanguageToUse(lang...)
	if nil != err {
		return nil, err
	}

	// split mnemonic into words
	words := strings.Split(norm.NFKD.String(mnemonic), " ")
	// calculate the expected entropy length and checksum length
	// 	ENT/8=MS*4/3, CS=(ENT/8)/4
	nWord := len(words)
	n := nWord * 4 / 3
	// check compatibility of entropy length
	if nWord%3 != 0 || !EntropyLenCompatible(n) {
		return nil, errors.New("invalid mnemonic length")
	}

	// recover the full entropy and delegates further
	x, y := new(big.Int), new(big.Int)
	for _, word := range words {
		idx, ok := dict.LookUp(word, language)
		if !ok {
			return nil, errors.New("missing word")
		}

		y.SetInt64(int64(idx))
		x.Lsh(x, WordIndexBitSize)
		x.Or(x, y)
	}

	//entropy := paddedPrepend(nil, x.Bytes(), n+1)
	return paddedPrepend(nil, x.Bytes(), n+1), nil
}
