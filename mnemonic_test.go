package bip39_test

import (
	"testing"

	"github.com/sammy00/bip39"
	"github.com/sammy00/bip39/dict"
)

func TestNewMnemonic_OK_en(t *testing.T) {
	testCases := bip39.ReadTrezorGoldenJSON(t)

	for i, c := range testCases {
		got, err := bip39.NewMnemonic(c.Entropy)

		if nil != err {
			t.Fatalf("#%d unexpected error: %v", i, err)
		}

		if got != c.Mnemonic {
			t.Fatalf("#%d invalid mnemonic: got %s, expect %s", i, got, c.Mnemonic)
		}
	}
}

func TestNewMnemonic_OK_jp(t *testing.T) {
	const lang = dict.Japanese

	var testCases []*bip39.GoldieJP
	bip39.ReadGoldenJSON(t, bip39.GoldenJP, &testCases)

	for i, c := range testCases {
		got, err := bip39.NewMnemonic(c.Entropy, lang)

		if nil != err {
			t.Fatalf("#%d unexpected error: %v", i, err)
		}

		if got != c.Mnemonic {
			t.Fatalf("#%d invalid mnemonic: got %s, expect %s", i, got, c.Mnemonic)
		}
	}
}

func TestValidateMnemonic_en_OK(t *testing.T) {
	testCases := bip39.ReadTrezorGoldenJSON(t)

	for i, c := range testCases {

		//t.Run("", func(st *testing.T) {
		if !bip39.ValidateMnemonic(c.Mnemonic) {
			t.Fatalf("#%d mnemonic should be valid", i)
		}
		//})
	}
}

func TestValidateMnemonic_jp_OK(t *testing.T) {
	const lang = dict.Japanese

	var testCases []*bip39.GoldieJP
	bip39.ReadGoldenJSON(t, bip39.GoldenJP, &testCases)

	for i, c := range testCases {
		if !bip39.ValidateMnemonic(c.Mnemonic, lang) {
			t.Fatalf("#%d mnemonic should be valid", i)
		}
	}
}
