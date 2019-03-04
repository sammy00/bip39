package dict_test

import (
	"testing"

	"github.com/derekparker/trie"
	"github.com/sammyne/bip39/dict"
)

func TestLookUp(t *testing.T) {
	type expect struct {
		index int
		ok    bool
	}

	mTrie := trie.New()
	mTrie.Add("hello", 1)
	mTrie.Add("world", 2)

	testCases := []struct {
		trie   *trie.Trie
		word   string
		expect expect
	}{
		{mTrie, "hello", expect{1, true}},
		{mTrie, "world", expect{2, true}},
		{mTrie, "hi", expect{0, false}},
	}

	for i, c := range testCases {
		index, ok := dict.LookUp(c.trie, c.word)

		if !ok && c.expect.ok {
			t.Fatalf("#%d fail to find word: %s", i, c.word)
		} else if ok && !c.expect.ok {
			t.Fatalf("#%d should not find the non-existent word: %s", i, c.word)
		}

		if c.expect.ok && c.expect.index != index {
			t.Fatalf("#%d invalid index: got %d, expect %d", i, c.expect.index, index)
		}
	}
}

func buildTrie(words []string, keys []int) *trie.Trie {
	t := trie.New()
	for i, word := range words {
		t.Add(word, keys[i])
	}

	return t
}
