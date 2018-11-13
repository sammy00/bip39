package dict

// Language denotes the type of languages
type Language int

// enumerations of different languages as named
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

// languageDescriptions maintains description for different languages
var languageDescriptions = make(map[Language]string)

// String returns the description bound to the language if registered,
// and empty string otherwise.
func (lang Language) String() string {
	return languageDescriptions[lang]
}
