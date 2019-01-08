package strand

import "strings"

// ToRNA creates the RNA complment for a given DNA strand
func ToRNA(dna string) string {
	return strings.Map(transcribe, dna)
}

func transcribe(r rune) rune {
	switch r {
	case 'G':
		return 'C'
	case 'C':
		return 'G'
	case 'T':
		return 'A'
	case 'A':
		return 'U'
	}
	return r
}
