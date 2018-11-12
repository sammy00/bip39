package dict

const (
	IdeographicSpaces = "\u3000"
	ASCIISpace        = " "
)

func Whitespace(lang Language) string {
	if Japanese == lang {
		return IdeographicSpaces
	}

	return ASCIISpace
}
