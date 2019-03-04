package bip39

import (
	"crypto/sha256"
	"io"
	"math/big"
	"strings"

	"github.com/sammyne/bip39/dict"
)

var wordIndexBitMask = new(big.Int).SetInt64(1<<WordIndexBitSize - 1)

type Mnemonic = string

func GenerateMnemonic(rand io.Reader, n EntropyLen,
	lang ...dict.Language) (Mnemonic, error) {
	if !EntropyLenCompatible(n) {
		return "", ErrEntropyLen
	}

	entropy := make([]byte, n)
	if _, err := io.ReadFull(rand, entropy); nil != err {
		return "", err
	}

	return NewMnemonic(entropy, lang...)
}

func NewMnemonic(entropy []byte, lang ...dict.Language) (Mnemonic, error) {
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
	//fmt.Printf("checksum=%x\n", checksumB&(1<<uint(checksumLen)-1))
	//fmt.Printf("checksum=%x\n", checksumB>>uint(8-checksumLen))
	//fmt.Printf("entropy'=%x\n", x.Bytes())

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

	//return norm.NFKD.String(strings.Join(words, " ")), nil
	//return strings.Join(words, " "), nil
	return strings.Join(words, dict.Whitespace(language)), nil
}

func ValidateMnemonic(mnemonic Mnemonic, lang ...dict.Language) bool {
	_, err := RecoverFullEntropy(mnemonic, lang...)
	return nil == err
}
