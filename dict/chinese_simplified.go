// +build ignore

package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/sammy00/bip39/dict"
)

//func init() {
func main() {
	data, err := ioutil.ReadFile(filepath.Join(dict.GoldenDictBase, "chinese_simplified.txt"))
	if nil != err {
		panic(err)
	}

	//fmt.Printf("%s\n", data)

	const expectCRC32 = "e3721bbf"
	if got := fmt.Sprintf("%x", crc32.ChecksumIEEE(data)); got != expectCRC32 {
		panic(errors.Errorf("invalid crc32: got %s, expect %s", got, expectCRC32))
	}
}
