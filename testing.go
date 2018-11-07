package bip39

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

const GoldenBase = "testdata"
const GoldenTrezor = "trezor.golden"

type Goldie struct {
	Entropy  []byte
	Mnemonic string
	Seed     []byte
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
