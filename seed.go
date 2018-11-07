package bip39

import (
	"crypto/sha512"

	"golang.org/x/crypto/pbkdf2"
)

func GenerateSeed(mnemonic Mnemonic, passphrase ...string) ([]byte, error) {
	saltS := "mnemonic"
	if 0 != len(passphrase) {
		saltS += passphrase[0]
	}

	password, salt := []byte(mnemonic), []byte(saltS)
	//pwd, salt := norm.NFKD.Bytes([]byte(mnemonic)), norm.NFKD.Bytes([]byte(saltS))
	return pbkdf2.Key(password, salt, 2048, 64, sha512.New), nil
}
