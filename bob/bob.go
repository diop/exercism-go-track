// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	sure := "Sure."
	chill := "Whoa, chill out!"
	fine := "Fine. Be that way!"
	whatever := "Whatever."
	remark = strings.Trim(remark, " \t\n\r")
	switch {
	case strings.ToUpper(remark) == remark && hasLetters(&remark):
		return chill
	case remark == "":
		return fine
	case strings.HasSuffix(remark, "?"):
		return sure
	default:
		return whatever
	}
}

func hasLetters(s *string) bool {
	minLetter := 'a'
	maxLetter := 'z'
	toLower := strings.ToLower(*s)
	for _, l := range toLower {
		if l <= maxLetter && l >= minLetter {
			return true
		}
	}
	return false
}
