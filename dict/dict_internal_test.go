package dict

import (
	"reflect"
	"testing"

	"github.com/derekparker/trie"
)

func TestTrieToUse(t *testing.T) {
	type expect struct {
		trie *trie.Trie
		lang Language
		err  error
	}

	testCases := []struct {
		lang   []Language
		expect expect
	}{
		{
			nil,
			expect{tries[English], English, nil},
		},
		{
			[]Language{Japanese},
			expect{tries[Japanese], Japanese, nil},
		},
		{
			[]Language{Reserved},
			expect{nil, Reserved, ErrDisabledTrie},
		},
	}

	for i, c := range testCases {
		trie, lang, err := TrieToUse(c.lang...)

		if err != c.expect.err {
			t.Fatalf("#%d unexpected error: got %v, expect %v", i, err, c.expect.err)
		}

		if nil != err {
			continue
		}

		if lang != c.expect.lang {
			t.Fatalf("#%d invalid language: got %s, expect %s", i, lang,
				c.expect.lang)
		}

		wordlist, _, err := WordlistToUse(lang)
		if nil != err {
			t.Fatalf("#%d unexpect error during fetching wordlist: %v", i, err)
		}

		if !IsTrie4Wordlist(trie, wordlist) {
			t.Fatalf("#%d invalid trie for wordlist", i)
		}
	}
}

// this test assume English and Japanese are registered and enabled,
// where English goes as default
func TestWordlistToUse(t *testing.T) {
	type expect struct {
		wordlist []string
		lang     Language
		err      error
	}

	testCases := []struct {
		lang   []Language
		expect expect
	}{
		{[]Language{Japanese}, expect{japanese(), Japanese, nil}},
		{[]Language{Reserved}, expect{nil, Reserved, ErrUnknownLanguage}},
		{nil, expect{english(), English, nil}},
	}

	UseLanguage(English)
	Register(Japanese, japanese, languageDescriptions[Japanese])

	for i, c := range testCases {
		wordlist, lang, err := WordlistToUse(c.lang...)

		if err != c.expect.err {
			t.Fatalf("#%d unexpected error: got %v, expect %v", i, err, c.expect.err)
		}

		if nil != err {
			continue
		}

		if !reflect.DeepEqual(wordlist, c.expect.wordlist) {
			t.Fatalf("#%d invalid wordlist: got %v, expect %v", i, wordlist,
				c.expect.wordlist)
		}

		if lang != c.expect.lang {
			t.Fatalf("#%d invalid language: got %s, expect %s", i, lang,
				c.expect.lang)
		}
	}
}

func IsTrie4Wordlist(trie *trie.Trie, wordlist []string) bool {
	keys := trie.Keys()

	if len(keys) != len(wordlist) {
		return false
	}

	keySet := make(map[string]struct{}, len(keys))
	for _, k := range keys {
		keySet[k] = struct{}{}
	}

	if len(keySet) != len(wordlist) {
		return false
	}

	for _, word := range wordlist {
		if _, ok := keySet[word]; !ok {
			return false
		}
	}

	return true
}
