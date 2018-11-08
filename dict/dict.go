package dict

import (
	"errors"

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
		return errors.New("language not registered")
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

func LanguageToUse(lang ...Language) (Language, error) {
	if 0 == len(lang) {
		return language, nil
	}

	if _, ok := wordlistGenerators[lang[0]]; !ok {
		return Reserved, errors.New("non-registered language")
	}

	return lang[0], nil
}

func Register(lang Language, generator WordlistGenerator,
	description string) error {
	if _, ok := wordlistGenerators[lang]; ok {
		return errors.New("language already registered")
	}

	wordlistGenerators[lang], languageDescriptions[lang] = generator, description

	return nil
}

func TrieToUse(lang ...Language) (*trie.Trie, Language, error) {
	if len(lang) > 0 {
		usable, ok := tries[lang[0]]
		if !ok {
			return nil, Reserved, errors.New("trie disabled")
		}

		return usable, lang[0], nil
	}

	if _, ok := tries[language]; !ok {
		return nil, Reserved, errors.New("trie disabled for language in use")
	}

	return tries[language], language, nil
}

func UseLanguage(lang Language) error {
	if _, ok := wordlistGenerators[lang]; !ok {
		return errors.New("non-registered language")
	}

	//Disable(language)

	language = lang
	Enable(language)

	return nil
}

func Wordlist(lang Language) ([]string, error) {
	generator, ok := wordlistGenerators[lang]
	if !ok {
		return nil, errors.New("non-registered language")
	}

	return generator(), nil
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
		return nil, Reserved, errors.New("non-registered language")
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
