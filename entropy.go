package bip39

import "github.com/sammy00/bip39/dict"

func EntropyLenCompatible(n EntropyLen) bool {
	return 0 == n%4 && n >= Size128 && n <= Size256
}

/*
func EntropyToFullBigInt(entropy []byte) (*big.Int, int, error) {
	n := len(entropy)
	if !EntropyLenCompatible(n) {
		return nil, 0, ErrEntropyLen
	}

	checksumB, checksumLen := sha256.Sum256(entropy)[0], n/32
	entropy = append(entropy, checksumB)

	x := new(big.Int).SetBytes(entropy)
	x.Rsh(x, uint(8-checksumLen))

	return x, n, nil
}
*/

/*
func EntropyToWordIndices(entropy []byte) ([]int64, error) {
	n := len(entropy)
	if !EntropyLenCompatible(n) {
		return nil, ErrEntropyLen
	}

	// make up the full entropy as a big int
	checksumB, checksumLen := sha256.Sum256(entropy)[0], n/32
	entropy = append(entropy, checksumB)

	x := new(big.Int).SetBytes(entropy)
	x.Rsh(x, uint(8-checksumLen))

	// MS=(ENT+CS)/11=(ENT+ENT/32)/11=3*ENT/32
	sentenceLen := 3 * n / 32
	indices := make([]int64, sentenceLen)

	wordIndex := new(big.Int)
	for i := len(indices) - 1; i >= 0; i-- {
		wordIndex.And(x, bitMask11)
	}
}
*/

// RecoverEntropy recovers the entropy without the checksum from
// the given mnemonic
func RecoverEntropy(mnemonic Mnemonic) ([]byte, error) {
	// recover the full entropy by RecoverFullEntropy()
	// trim out checksum
	return nil, nil
}

// RecoverFullEntropy recovers the entropy plus the checksum from
// the given mnemonic
func RecoverFullEntropy(mnemonic Mnemonic, lang ...dict.Language) (
	[]byte, error) {
	// vanity check against language
	//language, err := dict.LanguageToUse(lang...)
	//if nil != err {
	//	return nil, err
	//}

	// split mnemonic into words
	// calculate the expected entropy length and checksum length
	// 	ENT/8=MS*4/3, CS=(ENT/8)/4
	// check compatibility of entropy length
	// recover the full entropy and delegates further
	//  checking to ValidateFullEntropy()

	return nil, nil
}

// RecoverNativeFullEntropy recovers the entropy plus the checksum from
// the given mnemonic
//func RecoverNativeFullEntropy(mnemonic Mnemonic, lang dict.Language) (
//	[]byte, error) {
//
//	return nil, nil
//}

func ValidateFullEntropy(entropy []byte) error {
	// check compatibility of entropy length
	// split raw entropy and checksum C1
	// checksum the raw entropy to get the checksum C2
	// ensure C1==C2

	return nil
}
