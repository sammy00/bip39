// +build ignore

package main

import (
	"encoding/hex"

	"github.com/sammy00/bip39"
)

func main() {
	trezor := make(map[string][][]string)
	bip39.ReadGoldenJSON(nil, "trezor.json", &trezor)

	//fmt.Printf("%+v\n", trezor)

	// convert to mnemonic goldies to ease testing
	mnemonics := make([]*bip39.Goldie, len(trezor["english"]))
	for i, c := range trezor["english"] {
		entropy, _ := hex.DecodeString(c[0])
		seed, _ := hex.DecodeString(c[2])

		mnemonics[i] = &bip39.Goldie{
			Entropy:  entropy,
			Mnemonic: c[1],
			Seed:     seed,
		}
	}

	if err := bip39.WriteGoldenJSON("trezor.golden", mnemonics); nil != err {
		panic(err)
	}
}
