package isogram

import "unicode"

func IsIsogram(word string) bool {
	characters := map[rune]bool{}
	for _, char := range word {
		char := unicode.ToLower(char)
		if unicode.IsLetter(char) && characters[char] {
			return false
		}
		characters[char] = true
	}
	return true
}
