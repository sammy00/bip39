package bip39_test

import (
	"testing"
	"unicode"

	"github.com/sammy00/bip39"
	"github.com/sammy00/bip39/dict"
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

func TestNewMnemonic_OK_jp(t *testing.T) {
	const lang = dict.Japanese

	var testCases []*bip39.GoldieJP
	bip39.ReadGoldenJSON(t, bip39.GoldenJP, &testCases)

	testCases = testCases[:1]

	for _, c := range testCases {
		c := c

		t.Run("", func(st *testing.T) {
			got, err := bip39.NewMnemonic(c.Entropy, lang)

			if nil != err {
				st.Fatalf("unexpected error: %v", err)
			}

			if got != c.Mnemonic {
				//if !strings.EqualFold(got, c.Mnemonic) {
				st.Log(got, len(got))
				st.Log(c.Mnemonic, len(c.Mnemonic))

				gotS := []rune(got)
				mnemonicS := []rune(c.Mnemonic)
				for i := range mnemonicS {
					if gotS[i] != mnemonicS[i] {
						st.Logf("%+q - %+q", gotS[i], mnemonicS[i])
						st.Log(i, string(gotS[i:]), string(mnemonicS[i:]))
						break
					}
				}
				//words := strings.Split(c.Mnemonic, " ")
				//words := strings.Fields(c.Mnemonic)
				//st.Log(words[0])
				//j := strings.IndexFunc(c.Mnemonic, func(x rune) bool {
				//	return unicode.IsSpace(x)
				//})
				for _, v := range c.Mnemonic {
					if unicode.IsSpace(v) {
						st.Logf("%+q", v)
						break
					}
				}

				st.Fatalf("invalid mnemonic: got %s, expect %s", got, c.Mnemonic)
			}
		})
	}
}
