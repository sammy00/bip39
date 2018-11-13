package dict_test

import (
	"testing"

	"github.com/sammy00/bip39/dict"
)

func TestLanguage_String(t *testing.T) {
	testCases := []struct {
		lang   dict.Language
		expect string
	}{
		{dict.English, "English"},
		{dict.Japanese, "Japanese"},
		{dict.Reserved, ""},
	}

	for i, c := range testCases {
		if got := c.lang.String(); got != c.expect {
			t.Fatalf("#%d invalid description: got %s, expect %s", i, got, c.expect)
		}
	}
}
