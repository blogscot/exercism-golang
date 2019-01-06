package bob

import (
	"regexp"
	"strings"
	"unicode"
)

const (
	calm     = "Calm down, I know what I'm doing!"
	whoa     = "Whoa, chill out!"
	sure     = "Sure."
	fine     = "Fine. Be that way!"
	whatever = "Whatever."
)

func containsLetter(letters string) bool {
	return strings.IndexFunc(letters, unicode.IsLetter) >= 0
}

func isShouting(remark string) bool {
	return strings.ToUpper(remark) == remark && containsLetter(remark)
}

// Hey provides a snappy reply to your overgrown child
func Hey(remark string) string {
	re, _ := regexp.Compile("([^\\w?])")
	remark = re.ReplaceAllString(remark, "")

	isShouting := isShouting(remark)
	isQuestion := strings.HasSuffix(remark, "?")
	isSilent := len(remark) == 0

	switch {
	case isShouting && isQuestion:
		return calm
	case isShouting:
		return whoa
	case isQuestion:
		return sure
	case isSilent:
		return fine
	default:
		return whatever
	}

}
