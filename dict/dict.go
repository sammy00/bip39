package dict

import (
	"github.com/derekparker/trie"
)

type WordlistGenerator func() []string

var (
	//wordlists map[Language][]string
	wordlistGenerators map[Language]WordlistGenerator
	tries              map[Language]*trie.Trie

	language Language
	wordlist []string
)

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

func Disable(lang Language) {
	delete(tries, lang)
}

func Register(lang Language, generator WordlistGenerator,
	description string) error {
	if _, ok := wordlistGenerators[lang]; ok {
		return ErrOccupiedLanguage
	}

	wordlistGenerators[lang], languageDescriptions[lang] = generator, description

	return nil
}

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

func UseLanguage(lang Language) error {
	if _, ok := wordlistGenerators[lang]; !ok {
		//return errors.New("non-registered language")
		return ErrUnknownLanguage
	}

	language = lang
	Enable(language)

	return nil
}

func WordlistInUse() ([]string, Language) {
	return wordlist, language
}

func WordlistToUse(lang ...Language) ([]string, Language, error) {
	if 0 == len(lang) {
		wordlist, language := WordlistInUse()
		return wordlist, language, nil
	}

	//return WordListInUse(), nil

	generator, ok := wordlistGenerators[lang[0]]
	if !ok {
		return nil, lang[0], ErrUnknownLanguage
	}

	return generator(), lang[0], nil
}

func init() {
	wordlistGenerators = make(map[Language]WordlistGenerator)
	tries = make(map[Language]*trie.Trie)

	Register(English, english, languageDescriptions[English])
	Register(Japanese, japanese, languageDescriptions[Japanese])

	Enable(English)
	Enable(Japanese)

	language, wordlist = English, wordlistGenerators[English]()
}
