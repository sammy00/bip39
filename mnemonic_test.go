package bip39_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/sammyne/bip39"
	"github.com/sammyne/bip39/dict"
)

func TestGenerateMnemonic_en_Error(t *testing.T) {
	testCases := []struct {
		entropy []byte
		n       int
		expect  error
	}{
		{ // no error for comparison
			[]byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			16,
			nil,
		},
		{ // invalid entropy length
			[]byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			17,
			bip39.ErrEntropyLen,
		},
		{ // not enough entropy
			[]byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			16,
			io.ErrUnexpectedEOF,
		},
	}

	for i, c := range testCases {
		rand := bytes.NewBuffer(c.entropy)

		if _, got := bip39.GenerateMnemonic(rand, c.n); got != c.expect {
			t.Fatalf("#%d unexpected error: got %v, expect %v", i, got, c.expect)
		}
	}
}

func TestGenerateMnemonic_en_OK(t *testing.T) {
	testCases := ReadTrezorGoldenJSON(t)

	for i, c := range testCases {
		rand := bytes.NewBuffer(c.Entropy)

		//got, err := bip39.NewMnemonic(c.Entropy)
		got, err := bip39.GenerateMnemonic(rand, rand.Len())

		if nil != err {
			t.Fatalf("#%d unexpected error: %v", i, err)
		}

		if got != c.Mnemonic {
			t.Fatalf("#%d invalid mnemonic: got %s, expect %s", i, got, c.Mnemonic)
		}
	}
}

func TestNewMnemonic_en_Error(t *testing.T) {
	testCases := []struct {
		entropy []byte
		lang    dict.Language
		expect  error
	}{
		{ // no error for comparison
			[]byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			dict.English,
			nil,
		},
		{ // invalid entropy length
			[]byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x01,
			},
			dict.English,
			bip39.ErrEntropyLen,
		},
		{ // unsupported language
			[]byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			dict.Reserved,
			dict.ErrUnknownLanguage,
		},
	}

	for i, c := range testCases {
		if _, got := bip39.NewMnemonic(c.entropy, c.lang); got != c.expect {
			t.Fatalf("#%d unexpected error: got %v, expect %v", i, got, c.expect)
		}
	}
}

func TestNewMnemonic_en_OK(t *testing.T) {
	testCases := ReadTrezorGoldenJSON(t)

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

	var testCases []*GoldieJP
	ReadGoldenJSON(t, GoldenJP, &testCases)

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
	testCases := ReadTrezorGoldenJSON(t)

	for i, c := range testCases {
		if !bip39.ValidateMnemonic(c.Mnemonic) {
			t.Fatalf("#%d mnemonic should be valid", i)
		}
	}
}

func TestValidateMnemonic_jp_OK(t *testing.T) {
	const lang = dict.Japanese

	var testCases []*GoldieJP
	ReadGoldenJSON(t, GoldenJP, &testCases)

	for i, c := range testCases {
		if !bip39.ValidateMnemonic(c.Mnemonic, lang) {
			t.Fatalf("#%d mnemonic should be valid", i)
		}
	}
}
