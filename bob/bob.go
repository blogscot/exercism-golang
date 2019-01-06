package bob

import (
	"regexp"
	"strings"
	"unicode"
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
		return "Calm down, I know what I'm doing!"
	case isShouting:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	case isSilent:
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}

}
