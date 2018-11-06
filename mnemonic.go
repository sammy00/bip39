package bip39

import (
	"io"
)

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

	return "", nil
}

func ValidateMnemonic(mnemonic Mnemonic) bool {
	return false
}
