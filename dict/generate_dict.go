// +build ignore

package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/sammy00/bip39/dict"
)

const tmpl = `// Code generated by sammy00. DO NOT EDIT.
package dict

import (
	"hash/crc32"
	"strings"

	"github.com/pkg/errors"
)

// %s return the wordlist for the given language specified as the 
// function name
func %s() []string {
	const golden = ` + "`%s`" + `

	const expectCRC32 = %d
	if got := crc32.ChecksumIEEE([]byte(golden)); got != expectCRC32 {
		panic(errors.Errorf("invalid crc32: got %%x, expect %%x", got, expectCRC32))
	}

	return strings.Split(strings.TrimSpace(golden), "\n")

	// wordlist := strings.Split(strings.TrimSpace(golden), "\n")
	// sort in ascending order to speed up lookup 
	// sort.Strings(wordlist)

	// return wordlist
}

//func init() {
//	if !sort.StringsAreSorted(%s()) {
//		panic("messy wordlist")
//	}
//}
`

func main() {
	const GoldenDictBase = "golden"

	dictNames := []string{
		//"chinese_simpilified",
		//"chinese_traditional",
		"english",
		//"french",
		//"italian",
		"japanese",
		//"korean",
		//"spanish",
	}

	for _, dictName := range dictNames {
		writeLocale(dictName)
	}
}

func writeLocale(locale string) {
	data, err := ioutil.ReadFile(filepath.Join(dict.GoldenDictBase,
		locale+".txt"))
	if nil != err {
		panic(err)
	}

	checksum, wordlistName := crc32.ChecksumIEEE(data), ToCamelCase(locale)

	os.Remove(locale + ".go")
	fd, err := os.OpenFile(locale+".go", os.O_CREATE|os.O_WRONLY, 0644)
	if nil != err {
		panic(err)
	}
	defer fd.Close()

	//fmt.Fprintf(fd, tmpl, wordlistName, wordlistName, data, checksum)
	// add on order checking
	fmt.Fprintf(fd, tmpl, wordlistName, wordlistName, data, checksum,
		wordlistName)
}

func ToCamelCase(s string) string {
	str := strings.Title(strings.Replace(s, "_", " ", -1))
	str = strings.Replace(str, " ", "", -1)

	return strings.ToLower(str[:1]) + str[1:]
	//return str
}
