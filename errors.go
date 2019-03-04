package bip39

import "errors"

// enumerations of possible errors
var (
	// ErrChecksum signals the invalid checksum of a full entropy
	ErrChecksum = errors.New("invalid checksum")
	// ErrEntropyLen denotes a unsupported entropy length
	ErrEntropyLen = errors.New("incompatible entropy length")
	// ErrInvalidWord signals the mnemonic contains non-official words
	ErrInvalidWord = errors.New("mnemonic contains out-of-dict words")
	// ErrMnemonicLen signals a invalid mnemonic length
	ErrMnemonicLen = errors.New("invalid mnemonic length")
)
