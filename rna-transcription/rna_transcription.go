package strand

// ToRNA creates the RNA complment foa a given DNA strand
func ToRNA(dna string) string {
	var complement string

	for _, rna := range dna {

		switch rna {
		case 'C':
			complement += string('G')
		case 'G':
			complement += string('C')
		case 'T':
			complement += string('A')
		case 'A':
			complement += string('U')
		}
	}
	return complement
}
