package bip39

type Locale int

func GetLocale() Locale          { return 0 }
func SetLocale(locale Locale)    {}
func UseWordList(words []string) {}
