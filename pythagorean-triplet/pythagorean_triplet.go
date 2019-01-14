package pythagorean

// Triplet contains three ints
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) (triplets []Triplet) {

	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := min; c <= max; c++ {
				if a*a+b*b == c*c {
					triplets = append(triplets, Triplet{a, b, c})
				}
			}
		}
	}
	return
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) (triplets []Triplet) {
	tripletsInRange := Range(1, p)
	for _, t := range tripletsInRange {
		if t[0]+t[1]+t[2] == p {
			triplets = append(triplets, t)
		}
	}
	return
}
