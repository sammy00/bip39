package bip39

type EntropyLen int

const (
	Size128 = EntropyLen(16 + 4*iota)
	Size160
	Size192
	Size224
	Size256
)
