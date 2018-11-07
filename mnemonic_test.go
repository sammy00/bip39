package bip39_test

import (
	"testing"

	"github.com/sammy00/bip39"
)

func TestNewMnemonic_OK(t *testing.T) {
	var testCases []*bip39.Goldie
	bip39.ReadGoldenJSON(t, bip39.GoldenTrezor, &testCases)

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
