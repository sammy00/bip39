package dict

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

var English []string

func init() {
	data, err := ioutil.ReadFile(filepath.Join("dict", GoldenDictBase,
		"english.txt"))
	if nil != err {
		panic(err)
	}

	const expectCRC32 = "c1dbd296"
	if got := fmt.Sprintf("%x", crc32.ChecksumIEEE(data)); got != expectCRC32 {
		panic(errors.Errorf("invalid crc32: got %s, expect %s", got, expectCRC32))
	}

	English = strings.Split(strings.TrimSpace(string(data)), "\n")
}
