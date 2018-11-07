package bip39

import (
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

const GoldenBase = "testdata"
const GoldenTrezor = "trezor.golden"
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

/*
func ReadTrezorGoldenJSON(t *testing.T) []*Goldie {
	fd, err := os.Open(filepath.Join(GoldenBase, name))
	if nil != err {
		t.Fatal(err)
	}
	defer fd.Close()
}
*/

func WriteGoldenJSON(name string, v interface{}) error {
	fd, err := os.OpenFile(filepath.Join(GoldenBase, name),
		os.O_CREATE|os.O_WRONLY, 0600)
	if nil != err {
		return err
	}
	defer fd.Close()

	encoder := json.NewEncoder(fd)
	encoder.SetIndent("", "  ")
	return encoder.Encode(v)
}
