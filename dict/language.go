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

var languageDescriptions = make(map[Language]string)

func (lang Language) String() string {
	return languageDescriptions[lang]
}

/*
func init() {
	languageDescriptions = map[Language]string{
		ChineseSimpilified: "Chinese-Simplified",
		ChineseTraditional: "Chinese-Traditional",
		English:            "English",
		French:             "French",
		Italian:            "Italian",
		Japanese:           "Japanese",
		Korean:             "Korean",
		Spanish:            "Spanish",
	}
}
*/
