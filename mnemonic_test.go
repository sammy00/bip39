package bip39_test

import (
	"testing"

	"github.com/sammy00/bip39"
)

func TestNewMnemonic_OK_en(t *testing.T) {
	testCases := bip39.ReadTrezorGoldenJSON(t)

	for _, c := range testCases {
		c := c

		t.Run("", func(st *testing.T) {
			got, err := bip39.NewMnemonic(c.Entropy)

			if nil != err {
				st.Fatalf("unexpected error: %v", err)
			}

			if got != c.Mnemonic {
				st.Fatalf("invalid mnemonic: got %s, expect %s", got, c.Mnemonic)
			}
		})
	}
}

/*
func TestNewMnemonic_OK_jp(t *testing.T) {
	testCases := bip39.ReadTrezorGoldenJSON(t)

	var jp []*bip39.GoldieJP
	bip39.ReadGoldenJSON(t, bip39.GoldenJP, &testCases)
	for _, v := range jp {
		testCases = append(testCases, jp...)
	}

	for _, c := range testCases {
		c := c

		t.Run("", func(st *testing.T) {
			got, err := bip39.NewMnemonic(c.Entropy)

			if nil != err {
				st.Fatalf("unexpected error: %v", err)
			}

			if got != c.Mnemonic {
				st.Fatalf("invalid mnemonic: got %s, expect %s", got, c.Mnemonic)
			}
		})
	}
}
*/
