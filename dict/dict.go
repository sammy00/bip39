package dict

import "errors"

var (
	wordlists map[Language][]string

	lang     Language
	wordlist []string
)

func init() {
	wordlists = map[Language][]string{
		English: english(),
	}

	lang, wordlist = English, wordlists[English]
}

func LanguageInUse() Language {
	return lang
}

func Register(lang Language, wordlist []string) error {
	if _, ok := wordlists[lang]; ok {
		return errors.New("language already registered")
	}

	wordlists[lang] = wordlist
	return nil
}

func UseLanguage(lang Language) error {
	if _, ok := wordlists[lang]; !ok {
		return errors.New("non-registered language")
	}

	return nil
}

func WordListInUse() []string {
	return wordlist
}
