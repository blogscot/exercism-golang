package bob

import (
	"regexp"
	"strings"
	"unicode"
)

func containsLetter(letters string) bool {
	for _, letter := range letters {
		if unicode.IsLetter(letter) {
			return true
		}
	}
	return false
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
	isBlank := len(remark) == 0

	switch {
	case isShouting && isQuestion:
		return "Calm down, I know what I'm doing!"
	case isShouting:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	case isBlank:
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}

}
