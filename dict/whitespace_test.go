package dict_test

import (
	"testing"

	"github.com/sammyne/bip39/dict"
)

func TestWhitespace(t *testing.T) {
	testCases := []struct {
		lang   dict.Language
		expect string
	}{
		{dict.English, " "},
		{dict.Japanese, dict.IdeographicSpaces},
	}

	for i, c := range testCases {
		if got := dict.Whitespace(c.lang); got != c.expect {
			t.Fatalf("#%d invalid whitespace: got %s, expect %s", i, got, c.expect)
		}
	}
}
