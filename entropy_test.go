package bip39_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/sammy00/bip39/dict"

	"github.com/sammy00/bip39"
)

func TestDecodeFullEntropy(t *testing.T) {
	testCases := bip39.ReadTrezorGoldenJSON(t)

	for i, c := range testCases {

		data, err := bip39.RecoverFullEntropy(c.Mnemonic)
		if nil != err {
			t.Fatalf("#%d unexpected error: %v", i, err)
		}

		entropy, _, err := bip39.DecodeFullEntropy(data)
		if nil != err {
			t.Fatalf("#%d unexpected error: %v", i, err)
		}

		if !bytes.Equal(entropy, c.Entropy) {
			t.Fatalf("#%d invalid entropy: got %x, expect %x", i, entropy, c.Entropy)
		}
	}
}

func TestDecodeFullEntropy_Error(t *testing.T) {
	testCases := []struct {
		mnemonic bip39.Mnemonic
		lang     dict.Language
		expect   error // expected error
	}{
		{ // language is disabled
			"", dict.Reserved, dict.ErrDisabledTrie,
		},
		{ // #(words) should be a multiple of 3
			"hello world", dict.English, bip39.ErrMnemonicLen,
		},
		{ // invalid mnemonic length
			"hello world hi", dict.English, bip39.ErrMnemonicLen,
		},
		{ // about00 is out of dict
			"abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about00",
			dict.English,
			bip39.ErrInvalidWord,
		},
		{ // switch the last about and the abandon before it to
			// produce invalid checksum
			"abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about abandon",
			dict.English,
			bip39.ErrChecksum,
		},
	}

	wordlist, _ := dict.WordlistInUse()
	for _, word := range wordlist[:20] {
		fmt.Printf(`"%s",`, word)
	}
	fmt.Println()

	for i, c := range testCases {
		_, err := bip39.RecoverFullEntropy(c.mnemonic, c.lang)

		if nil == err || err.Error() != c.expect.Error() {
			t.Fatalf("#%d got unexpected error %v, expect %v", i, err, c.expect)
		}
	}
}
