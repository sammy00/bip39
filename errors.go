package bip39

import "errors"

var (
	ErrChecksum    = errors.New("invalid checksum")
	ErrEntropyLen  = errors.New("incompatible entropy length")
	ErrInvalidWord = errors.New("mnemonic contains out-of-dict words")
	ErrMnemonicLen = errors.New("invalid mnemonic length")
)
