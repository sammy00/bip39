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

	lang     Language
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

/*
func LanguageInUse() Language {
	return lang
}
*/

func LanguageToUse(language ...Language) (Language, error) {
	if 0 == len(language) {
		return lang, nil
	}

	if _, ok := wordlistGenerators[language[0]]; !ok {
		return Reserved, errors.New("non-registered language")
	}

	return language[0], nil
}

func LookUp(word string, lang ...Language) (int, bool) {
	language, err := LanguageToUse(lang...)
	if nil != err {
		return -1, false
	}

	trie, ok := tries[language]
	if !ok {
		return -1, false
	}

	w, ok := trie.Find(word)
	return w.Meta().(int), ok
}

func Register(lang Language, generator WordlistGenerator,
	description string) error {
	if _, ok := wordlistGenerators[lang]; ok {
		return errors.New("language already registered")
	}

	wordlistGenerators[lang], languageDescriptions[lang] = generator, description

	return nil
}

func UseLanguage(lang Language) error {
	//if _, ok := wordlists[lang]; !ok {
	if _, ok := wordlistGenerators[lang]; !ok {
		return errors.New("non-registered language")
	}

	return nil
}

func Wordlist(lang Language) ([]string, error) {
	generator, ok := wordlistGenerators[lang]
	if !ok {
		return nil, errors.New("non-registered language")
	}

	return generator(), nil
}

//func WordListInUse() []string {
func WordlistInUse() ([]string, Language) {
	return wordlist, lang
}

func WordlistToUse(lang ...Language) ([]string, Language, error) {
	if 0 == len(lang) {
		wordlist, language := WordlistInUse()
		return wordlist, language, nil
		//return Wordlist(lang[0])
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

	lang, wordlist = English, wordlistGenerators[English]()
}
