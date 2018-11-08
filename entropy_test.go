package bip39_test

import (
	"bytes"
	"testing"

	"github.com/sammy00/bip39"
)

func TestDecodeFullEntropy(t *testing.T) {
	testCases := bip39.ReadTrezorGoldenJSON(t)

	for _, c := range testCases {

		t.Run("", func(st *testing.T) {
			data, err := bip39.RecoverFullEntropy(c.Mnemonic)
			if nil != err {
				st.Fatal(err)
			}

			entropy, _, err := bip39.DecodeFullEntropy(data)
			if nil != err {
				st.Fatal(err)
			}

			if !bytes.Equal(entropy, c.Entropy) {
				st.Fatalf("invalid entropy: got %x, expect %x", entropy, c.Entropy)
			}
		})
	}
}
