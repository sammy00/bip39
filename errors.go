package bip39

import "errors"

var (
	ErrEntropyLen = errors.New("incompatible entropy length")
	ErrChecksum   = errors.New("invalid checksum")
)
