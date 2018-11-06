package bip39

func EntropyLenCompatible(n EntropyLen) bool {
	return 0 == n%32 && n >= Size128 && n <= Size256
}

// RecoverEntropy recovers the entropy without the checksum from
// the given mnemonic
func RecoverEntropy(mnemonic Mnemonic) ([]byte, error) {
	return nil, nil
}

// RecoverFullEntropy recovers the entropy plus the checksum from
// the given mnemonic
func RecoverFullEntropy(mnemonic Mnemonic) ([]byte, error) {
	return nil, nil
}

func ValidateFullEntropy(entropy []byte) bool {
	return false
}
