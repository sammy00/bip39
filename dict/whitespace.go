package dict

// spaces enumerations
const (
	// IdeographicSpaces is the whitespace used specifically for Japanese
	IdeographicSpaces = "\u3000"
	// ASCIISpace is the default whitespace employed for languages except above
	ASCIISpace = " "
)

// Whitespace finds the whitespace for the given language
func Whitespace(lang Language) string {
	if Japanese == lang {
		return IdeographicSpaces
	}

	return ASCIISpace
}
