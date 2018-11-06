package bip39

type Locale int

var wordList []string

func GetLocale() Locale                      { return 0 }
func Locales() []Locale                      { return nil }
func SetLocale(locale Locale)                {}
func UseWordList(words []string)             {}
func WordList4Locale(locale Locale) []string { return nil }
