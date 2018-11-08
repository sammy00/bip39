package dict

import "github.com/derekparker/trie"

func LookUp(trie *trie.Trie, word string) (int, bool) {
	w, ok := trie.Find(word)

	if !ok {
		return -1, false
	}

	return w.Meta().(int), true
}
