package dict

import "errors"

type WordlistGenerator func() []string

var (
	//wordlists map[Language][]string
	wordlistGenerators map[Language]WordlistGenerator

	lang     Language
	wordlist []string
)

func init() {
	//wordlists = map[Language][]string{
	//	English: english(),
	//}
	wordlistGenerators = map[Language]WordlistGenerator{
		English: english,
	}

	//lang, wordlist = English, wordlists[English]
	lang, wordlist = English, wordlistGenerators[English]()
}

func LanguageInUse() Language {
	return lang
}

func Register(lang Language, generator WordlistGenerator,
	description string) error {
	/*
		if _, ok := wordlists[lang]; ok {
			return errors.New("language already registered")
		}

		wordlists[lang], languageDescriptions[lang] = wordlist, description
		return nil
	*/
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

func WordListInUse() []string {
	return wordlist
}
