package bip39

import (
	"crypto/sha256"
	"fmt"
	"io"
	"math/big"
	"strings"

	"github.com/sammy00/bip39/dict"
	"golang.org/x/text/unicode/norm"
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

	var wordlist []string
	if 0 == len(lang) {
		wordlist = dict.WordListInUse()
	} else {
		var err error
		if wordlist, err = dict.Wordlist(lang[0]); nil != err {
			return "", err
		}
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

	//return strings.Join(words, " "), nil

	//native := strings.Join(words, " ")
	//return string(norm.NFKD.Bytes([]byte(native))), nil

	//return norm.NFKD.String(strings.Join(words, " ")), nil
	//whiteSpace := norm.NFKD.String("\u3000")
	whiteSpace := "\u3000"
	fmt.Printf("*%s*\n", whiteSpace)
	fmt.Println("*\u3000*")
	fmt.Println("* *")
	return norm.NFKD.String(strings.Join(words, whiteSpace)), nil
	//return strings.Join(words, whiteSpace), nil
}

func ValidateMnemonic(mnemonic Mnemonic) bool {
	return false
}
