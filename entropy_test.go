package bip39_test

import (
	"bytes"
	"testing"

	"github.com/sammy00/bip39/dict"

	"github.com/sammy00/bip39"
)

func TestDecodeFullEntropy(t *testing.T) {
	testCases := ReadTrezorGoldenJSON(t)

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

func TestRecoverFullEntropy_Error(t *testing.T) {
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
		{ // switch the last "about" and the "abandon" before it to
			// produce invalid checksum
			"abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about abandon",
			dict.English,
			bip39.ErrChecksum,
		},
	}

	for i, c := range testCases {
		_, err := bip39.RecoverFullEntropy(c.mnemonic, c.lang)

		if nil == err || err.Error() != c.expect.Error() {
			t.Fatalf("#%d got unexpected error %v, expect %v", i, err, c.expect)
		}
	}
}

func TestDecodeFullEntropy_Error(t *testing.T) {
	testCases := []struct {
		data   []byte
		expect error // the expected error
	}{
		{ // no error for comparison
			[]byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x03,
			},
			nil,
		},
		{ // invalid entropy length
			[]byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x03, 0x04,
			},
			bip39.ErrEntropyLen,
		},
		{ // invalid checksum
			[]byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x04,
			},
			bip39.ErrChecksum,
		},
	}

	/*
		const mnemonic = "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about"
		entropy, _ := bip39.RecoverFullEntropy(mnemonic)
		for _, v := range entropy {
			fmt.Printf("0x%02x,", v)
		}
		fmt.Println()
	*/

	for i, c := range testCases {
		_, _, err := bip39.DecodeFullEntropy(c.data)

		if err != c.expect {
			t.Fatalf("#%d got unexpected error %v, expect %v", i, err, c.expect)
		}
	}
}
