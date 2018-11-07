package bip39

import (
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

//go:generate curl -s  https://raw.githubusercontent.com/trezor/python-mnemonic/master/vectors.json -o testdata/trezor.json

//go:generate curl -s https://raw.githubusercontent.com/bip32JP/bip32JP.github.io/master/test_JP_BIP39.json -o testdata/test_JP_BIP39.json

const GoldenBase = "testdata"
const GoldenTrezor = "trezor.json"
const GoldenJP = "test_JP_BIP39.json"

type Goldie struct {
	Entropy  []byte
	Mnemonic string
	Seed     []byte
}

type GoldieJP struct {
	Entropy            []byte
	Mnemonic           string
	Passphrase         string
	Seed               []byte
	ExtendedPrivateKey string
}

type GoldieTrezor GoldieJP

func (goldie *GoldieJP) UnmarshalJSON(data []byte) error {
	var jp map[string]string

	//fmt.Println("hello")
	if err := json.Unmarshal(data, &jp); nil != err {
		return err
	}

	var err error
	if goldie.Entropy, err = hex.DecodeString(jp["entropy"]); nil != err {
		return err
	}

	goldie.Mnemonic = jp["mnemonic"]
	goldie.Passphrase = jp["passphrase"]

	if goldie.Seed, err = hex.DecodeString(jp["seed"]); nil != err {
		return err
	}

	goldie.ExtendedPrivateKey = jp["bip32_xprv"]

	return nil
}

func (goldie *GoldieTrezor) UnmarshalJSON(data []byte) error {
	var trezor []string
	if err := json.Unmarshal(data, &trezor); nil != err {
		return err
	}

	var err error
	if goldie.Entropy, err = hex.DecodeString(trezor[0]); nil != err {
		return err
	}

	goldie.Mnemonic = trezor[1]
	goldie.Passphrase = "TREZOR"

	if goldie.Seed, err = hex.DecodeString(trezor[2]); nil != err {
		return err
	}

	goldie.ExtendedPrivateKey = trezor[3]

	return nil
}

func ReadGoldenJSON(t *testing.T, name string, v interface{}) {
	fd, err := os.Open(filepath.Join(GoldenBase, name))
	if nil != err {
		t.Fatal(err)
	}
	defer fd.Close()

	if err := json.NewDecoder(fd).Decode(v); nil != err {
		t.Fatal(err)
	}
}

func ReadTrezorGoldenJSON(t *testing.T) []*GoldieTrezor {
	trezor := make(map[string][]*GoldieTrezor)
	ReadGoldenJSON(t, GoldenTrezor, &trezor)

	return trezor["english"]
}
