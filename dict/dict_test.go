package dict_test

import (
	"testing"

	"github.com/sammy00/bip39/dict"
)

func TestEnable(t *testing.T) {
	testCases := []struct {
		lang   dict.Language
		expect error
	}{
		{dict.Japanese, nil},
		{dict.English, nil},
		{dict.Spanish, dict.ErrUnknownLanguage},
	}

	dict.Disable(dict.Japanese)
	defer dict.Enable(dict.Japanese)

	for i, c := range testCases {
		if err := dict.Enable(c.lang); err != c.expect {
			t.Fatalf("#%d unexpected error: got %v, expect %v", i, err, c.expect)
		}
	}
}

func TestRegister(t *testing.T) {
	testCases := []struct {
		lang        dict.Language
		generator   dict.WordlistGenerator
		description string
		expect      error
	}{
		{
			dict.Language(0xff),
			func() []string { return []string{"hello", "world"} },
			"dummy",
			nil,
		},
		{
			dict.English,
			func() []string { return []string{"hello", "world"} },
			"dummy",
			dict.ErrOccupiedLanguage,
		},
	}

	for i, c := range testCases {
		if got := dict.Register(c.lang, c.generator,
			c.description); got != c.expect {
			t.Fatalf("#%d unexpected error: got %v, expect %v", i, got, c.expect)
		}
	}
}

func TestUseLanguage(t *testing.T) {
	testCases := []struct {
		lang   dict.Language
		expect error
	}{
		{dict.English, nil},
		{dict.Reserved, dict.ErrUnknownLanguage},
	}

	for i, c := range testCases {
		if got := dict.UseLanguage(c.lang); got != c.expect {
			t.Fatalf("#%d unexpected error: got %v, expect %v", i, got, c.expect)
		}
	}
}
