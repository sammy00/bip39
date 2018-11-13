package dict

import (
	"github.com/derekparker/trie"
)

// WordlistGenerator is the signature of function generating a wordlist
type WordlistGenerator func() []string

var (
	wordlistGenerators map[Language]WordlistGenerator
	tries              map[Language]*trie.Trie

	language Language
	wordlist []string
)

// Enable constructs the trie based on the wordlist bound to the given language
func Enable(lang Language) error {
	generator, ok := wordlistGenerators[lang]
	if !ok {
		return ErrUnknownLanguage
	}

	if _, ok := tries[lang]; ok {
		return nil
	}

	tries[lang] = trie.New()
	for i, word := range generator() {
		tries[lang].Add(word, i)
	}

	return nil
}

// Disable clean up the trie bound to the given language
func Disable(lang Language) {
	delete(tries, lang)
}

// Register takes into record a given language, its wordlist generator and a
// brief description
func Register(lang Language, generator WordlistGenerator,
	description string) error {
	if _, ok := wordlistGenerators[lang]; ok {
		return ErrOccupiedLanguage
	}

	wordlistGenerators[lang], languageDescriptions[lang] = generator, description

	return nil
}

// TrieToUse gets the trie bound to the given language if any. If no language
// provided, the global default will be employed. Both the trie and language
// bound to it (is the global default language if no language is provided)
// will be returned
func TrieToUse(lang ...Language) (*trie.Trie, Language, error) {
	l := language // default as the language in use
	if len(lang) > 0 {
		l = lang[0]
	}

	if _, ok := tries[l]; !ok {
		return nil, l, ErrDisabledTrie
	}

	return tries[l], l, nil
}

// UseLanguage overrides the global default language to use.
func UseLanguage(lang Language) error {
	if _, ok := wordlistGenerators[lang]; !ok {
		return ErrUnknownLanguage
	}

	language = lang
	Enable(language)

	return nil
}

// WordlistToUse returns the wordlist bound to a language, which is the
// provided one if any, otherwise the global default language. The 2nd
// output is the language bound to the wordlist (the 1st output) in case
// of no error.
func WordlistToUse(lang ...Language) ([]string, Language, error) {
	if 0 == len(lang) {
		return wordlist, language, nil
	}

	generator, ok := wordlistGenerators[lang[0]]
	if !ok {
		return nil, lang[0], ErrUnknownLanguage
	}

	return generator(), lang[0], nil
}

// init registers and enables the English and Japanese, and
// set the english as the global default
func init() {
	wordlistGenerators = make(map[Language]WordlistGenerator)
	tries = make(map[Language]*trie.Trie)

	Register(English, english, "English")
	Register(Japanese, japanese, "Japanese")

	Enable(English)
	Enable(Japanese)

	language, wordlist = English, wordlistGenerators[English]()
}
