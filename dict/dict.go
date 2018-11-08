package dict

import (
	"errors"
	"sort"
)

type WordlistGenerator func() []string

var (
	//wordlists map[Language][]string
	wordlistGenerators map[Language]WordlistGenerator

	lang     Language
	wordlist []string
)

func LanguageInUse() Language {
	return lang
}

func LanguageToUse(lang ...Language) (Language, error) {
	if 0 == len(lang) {
		return LanguageInUse(), nil
	}

	if _, ok := wordlistGenerators[lang[0]]; !ok {
		return Reserved, errors.New("non-registered language")
	}

	return lang[0], nil
}

func LookUp(word string, lang ...Language) (int, bool) {
	wordlist, err := WordlistToUse(lang...)
	if nil != err {
		return -1, false
	}

	j := sort.SearchStrings(wordlist, word)

	return j, (j != len(wordlist)) && (wordlist[j] == word)
}

func LookUpMissing(lang Language, words ...string) int {

	wordlist, err := WordlistToUse(lang)
	if nil != err {
		return 0
	}

	wordlistLen := len(wordlist)
	for i, word := range words {

		j := sort.SearchStrings(wordlist, word)
		if j == wordlistLen || wordlist[j] != word {
			return i
		}
	}

	return -1
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

func WordListInUse() []string {
	return wordlist
}

func WordlistToUse(lang ...Language) ([]string, error) {
	if len(lang) > 0 {
		return Wordlist(lang[0])
	}

	return WordListInUse(), nil
}

func init() {
	//wordlists = map[Language][]string{
	//	English: english(),
	//}
	wordlistGenerators = map[Language]WordlistGenerator{
		English:  english,
		Japanese: japanese,
	}

	//lang, wordlist = English, wordlists[English]
	lang, wordlist = English, wordlistGenerators[English]()
}
