package bip39

import (
	"crypto/sha256"
	"io"
	"math/big"
	"strings"

	"github.com/sammy00/bip39/dict"
)

var wordIndexBitMask = new(big.Int).SetInt64(1<<WordIndexBitSize - 1)

type Mnemonic = string

func GenerateMnemonic(rand io.Reader, n EntropyLen) (Mnemonic, error) {
	if !EntropyLenCompatible(n) {
		return "", ErrEntropyLen
	}

	entropy := make([]byte, n)
	if _, err := io.ReadFull(rand, entropy); nil != err {
		return "", err
	}

	return NewMnemonic(entropy)
}

func NewMnemonic(entropy []byte) (Mnemonic, error) {
	n := len(entropy)
	if !EntropyLenCompatible(n) {
		return "", ErrEntropyLen
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
	//indices := make([]int64, sentenceLen)
	words := make([]string, nWord)

	wordList, wordIndex := dict.WordListInUse(), new(big.Int)
	//fmt.Println(wordList)
	for i := nWord - 1; i >= 0; i-- {
		wordIndex.And(x, wordIndexBitMask)
		x.Rsh(x, WordIndexBitSize)

		words[i] = wordList[wordIndex.Int64()]
	}

	return strings.Join(words, " "), nil
}

func ValidateMnemonic(mnemonic Mnemonic) bool {
	return false
}
