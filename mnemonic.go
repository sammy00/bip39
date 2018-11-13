package bip39

import (
	"crypto/sha256"
	"io"
	"math/big"
	"strings"

	"github.com/sammy00/bip39/dict"
)

// wordIndexBitMask is the bits masker assisting in decoding word indices
// out of the full entropy
var wordIndexBitMask = new(big.Int).SetInt64(1<<WordIndexBitSize - 1)

// GenerateMnemonic constructs a mnemonic randomly based on the n bytes read
// out of the provided random sources for the provided language (as the
// global default if none)
func GenerateMnemonic(rand io.Reader, n EntropyLen,
	lang ...dict.Language) (string, error) {
	if !EntropyLenCompatible(n) {
		return "", ErrEntropyLen
	}

	entropy := make([]byte, n)
	if _, err := io.ReadFull(rand, entropy); nil != err {
		return "", err
	}

	return NewMnemonic(entropy, lang...)
}

// NewMnemonic constructs the mnemonic w.r.t a language (the global default
// configured by dict pkg if none provided) for the given entropy
func NewMnemonic(entropy []byte, lang ...dict.Language) (string, error) {
	n := len(entropy)
	if !EntropyLenCompatible(n) {
		return "", ErrEntropyLen
	}

	wordlist, language, err := dict.WordlistToUse(lang...)
	if nil != err {
		return "", err
	}

	// make up the full entropy as a big int
	// checksumLen=n*8/32=n/4
	checksumB, checksumLen := sha256.Sum256(entropy)[0], n/4
	entropy = append(entropy, checksumB)

	x := new(big.Int).SetBytes(entropy)
	x.Rsh(x, uint(8-checksumLen))

	// MS=(ENT+CS)/11=(ENT+ENT/32)/11=3*ENT/32
	// if measured in bytes, we got
	//   MS=(3*ENT/8)/(32/8)=3*(ENT/8)/4
	nWord := 3 * n / 4
	words := make([]string, nWord)

	wordIndex := new(big.Int)
	for i := nWord - 1; i >= 0; i-- {
		wordIndex.And(x, wordIndexBitMask)
		x.Rsh(x, WordIndexBitSize)

		words[i] = wordlist[wordIndex.Int64()]
	}

	return strings.Join(words, dict.Whitespace(language)), nil
}

// ValidateMnemonic checks if the mnemonic is valid against a language
// (default as the one set by dict pkg)
func ValidateMnemonic(mnemonic string, lang ...dict.Language) bool {
	_, err := RecoverFullEntropy(mnemonic, lang...)
	return nil == err
}
