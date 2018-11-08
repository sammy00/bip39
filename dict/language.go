package dict

type Language int

const (
	Reserved = Language(iota)
	ChineseSimpilified
	ChineseTraditional
	English
	French
	Italian
	Japanese
	Korean
	Spanish
)

var languageDescriptions = map[Language]string{}

func (lang Language) String() string {
	return languageDescriptions[lang]
}
