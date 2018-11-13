package bip39

// EntropyLen alias int to ease semantic understanding
type EntropyLen = int

// enumerations of supported entropy length in bytes
const (
	Size128 = EntropyLen(16 + 4*iota)
	Size160
	Size192
	Size224
	Size256
)

// WordIndexBitSize specifies the bit width for splitting the full entropy
// (raw+checksum) into index groups
const WordIndexBitSize = 11
