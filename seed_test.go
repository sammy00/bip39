package bip39_test

import (
	"bytes"
	"testing"

	"github.com/sammy00/bip39"
)

func TestGenerateSeed(t *testing.T) {
	var testCases []*bip39.Goldie
	bip39.ReadGoldenJSON(t, bip39.GoldenTrezor, &testCases)

	const passphrase = "TREZOR"

	for _, c := range testCases {
		c := c

		t.Run("", func(st *testing.T) {
			got, err := bip39.GenerateSeed(c.Mnemonic, passphrase)

			if nil != err {
				st.Fatal(err)
			}

			if !bytes.Equal(got, c.Seed) {
				st.Fatalf("invalid seed: got %x, expect %x", got, c.Seed)
			}
		})
	}
}

func TestGenerateSeed_jp(t *testing.T) {
	var testCases []*bip39.GoldieJP
	bip39.ReadGoldenJSON(t, bip39.GoldenJP, &testCases)

	//const passphrase = "TREZOR"

	testCases = testCases[:1]
	for _, c := range testCases {
		c := c

		t.Run("", func(st *testing.T) {
			got, err := bip39.GenerateSeed(c.Mnemonic, c.Passphrase)

			if nil != err {
				st.Fatal(err)
			}

			if !bytes.Equal(got, c.Seed) {
				st.Fatalf("invalid seed: got %x, expect %x", got, c.Seed)
			}
		})
	}
}
