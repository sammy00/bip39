package bip39_test

import (
	"bytes"
	"testing"

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
