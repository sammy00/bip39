package bip39_test

import (
	"bytes"
	"testing"

	"github.com/sammy00/bip39"
)

func TestGenerateSeed_en(t *testing.T) {
	testCases := bip39.ReadTrezorGoldenJSON(t)

	for i, c := range testCases {
		got, err := bip39.GenerateSeed(c.Mnemonic, c.Passphrase)

		if nil != err {
			t.Fatalf("#%d unexpected: %v", i, err)
		}

		if !bytes.Equal(got, c.Seed) {
			t.Fatalf("#%d invalid seed: got %x, expect %x", i, got, c.Seed)
		}
	}
}

func TestGenerateSeed_jp(t *testing.T) {
	var testCases []*bip39.GoldieJP
	bip39.ReadGoldenJSON(t, bip39.GoldenJP, &testCases)

	for i, c := range testCases {

		got, err := bip39.GenerateSeed(c.Mnemonic, c.Passphrase)
		if nil != err {
			t.Fatalf("#%d unexpected: %v", i, err)
		}

		if !bytes.Equal(got, c.Seed) {
			t.Fatalf("#%d invalid seed: got %x, expect %x", i, got, c.Seed)
		}
	}
}
