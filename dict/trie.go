package dict

import "github.com/derekparker/trie"

// LookUp searchs a given word against the given trie, and return the index
// (stored in its meta) of the word.
func LookUp(trie *trie.Trie, word string) (int, bool) {
	w, ok := trie.Find(word)

	if !ok {
		return -1, false
	}

	return w.Meta().(int), true
}
