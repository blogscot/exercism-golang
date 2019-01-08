package piglatin

import (
	"regexp"
	"strings"
)

// Sentence translates the given English sentence into Pig Latin
func Sentence(sentence string) string {
	piglatin := []string{}
	for _, word := range strings.Split(sentence, " ") {
		piglatin = append(piglatin, translate(word))
	}
	return strings.Join(piglatin, " ")
}

func translate(word string) string {
	if match, _ := regexp.MatchString("^(yt|xr|[aeiou]|[aeiou]qu).*", word); match {
		return word + "ay"
	} else if match, _ := regexp.MatchString("^(thr|sch|[^aeiou]qu).*", word); match {
		return word[3:] + word[:3] + "ay"
	} else if match, _ := regexp.MatchString("^(ch|qu|th|rh).*", word); match {
		return word[2:] + word[:2] + "ay"
	}
	return word[1:] + word[:1] + "ay"
}
