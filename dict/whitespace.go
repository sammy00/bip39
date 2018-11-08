package dict

const IdeographicSpaces = "\u3000"

func Whitespace(lang Language) string {
	if Japanese == lang {
		return IdeographicSpaces
	}

	return " "
}
