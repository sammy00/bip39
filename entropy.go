package bip39

import (
	"crypto/sha256"
	"math/big"
	"strings"

	"golang.org/x/text/unicode/norm"

	"github.com/sammyne/bip39/dict"
)

func EntropyLenCompatible(n EntropyLen) bool {
	return 0 == n%4 && n >= Size128 && n <= Size256
}

// RecoverFullEntropy recovers the entropy plus the checksum from
// the given mnemonic
func RecoverFullEntropy(mnemonic Mnemonic, lang ...dict.Language) (
	[]byte, error) {
	// vanity check against language
	trie, _, err := dict.TrieToUse(lang...)
	if nil != err {
		return nil, err
	}

	// split mnemonic into words
	words := strings.Fields(norm.NFKD.String(mnemonic))

	// calculate the expected entropy length and checksum length
	// 	ENT/8=MS*4/3, CS=(ENT/8)/4
	nWord := len(words)
	n := nWord * 4 / 3
	// check compatibility of entropy length
	if nWord%3 != 0 || !EntropyLenCompatible(n) {
		return nil, ErrMnemonicLen
	}

	// recover the full entropy and delegates further
	x, y := new(big.Int), new(big.Int)
	for _, word := range words {
		//idx, ok := dict.LookUp(word, language)
		idx, ok := dict.LookUp(trie, word)
		if !ok {
			return nil, ErrInvalidWord
		}

		y.SetInt64(int64(idx))
		x.Lsh(x, WordIndexBitSize)
		x.Or(x, y)
	}

	entropy := paddedPrepend(nil, x.Bytes(), n+1)
	if _, _, err := DecodeFullEntropy(entropy); nil != err {
		return nil, err
	}

	return entropy, nil
}

func DecodeFullEntropy(data []byte) ([]byte, byte, error) {
	n := len(data) - 1
	if !EntropyLenCompatible(n) {
		return nil, 0, ErrEntropyLen
	}

	// k1 is the length of checksum
	k1, k2 := uint(n/4), uint(8-n/4)

	// decode checksum
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
