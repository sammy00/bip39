package bip39

import (
	"crypto/sha512"

	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/text/unicode/norm"
)

// GenerateSeed converts a mnemonic into a 512-bits seed based on a passphrase
// which is empty if not provided.
func GenerateSeed(mnemonic string, passphrase ...string) ([]byte, error) {
	saltS := "mnemonic"
	if 0 != len(passphrase) {
		saltS += passphrase[0]
	}

	password := norm.NFKD.Bytes([]byte(mnemonic))
	salt := norm.NFKD.Bytes([]byte(saltS))

	return pbkdf2.Key(password, salt, 2048, 64, sha512.New), nil
}
