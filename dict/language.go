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

/*
var languageDescriptions = map[Language]string{
	Reserved:           "Reserved",
	ChineseSimpilified: "ChineseSimpilified",
	ChineseTraditional: "ChineseTraditional",
	English:            "English",
	French:             "French",
	Italian:            "Italian",
	Japanese:           "Japanese",
	Korean:             "Korean",
	Spanish:            "Spanish",
}
*/

func (lang Language) String() string {
	return languageDescriptions[lang]
}
