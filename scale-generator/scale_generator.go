package scale

import (
	"strings"
)

var scaleWithSharps = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
var scaleWithFlats = []string{"F", "Gb", "G", "Ab", "A", "Bb", "B", "C", "Db", "D", "Eb", "E"}
var useSharps = []string{
	"C", "G", "D", "A", "E", "B", "F# major", "a", "e", "b", "f#", "c#", "g#", "d# minor",
}

type myScale []string

// Scale generates a musical scale starting with the tonic
// and following the specified interval pattern
func Scale(tonic, interval string) []string {
	return myScale(buildScale(tonic)).applyInterval(interval)
}

func buildScale(tonic string) []string {
	newScale := make([]string, 0)

	scale := getScale(tonic)
	start := myScale(scale).findTonic(toUpper(tonic))
	for i := range scale {
		newScale = append(newScale, scale[(start+i)%len(scale)])
	}
	return newScale
}

func (scale myScale) applyInterval(interval string) []string {
	position := 1
	for _, step := range interval {
		switch step {
		case 'M':
			scale = myScale(scale).remove(position, 1)
		case 'A':
			scale = myScale(scale).remove(position, 2)
		}
		position++
	}
	return scale
}

func getScale(tonic string) []string {
	for _, v := range useSharps {
		if v == tonic {
			return scaleWithSharps
		}
	}
	return scaleWithFlats
}

func (scale myScale) findTonic(tonic string) int {
	for i := range scale {
		if scale[i] == tonic {
			return i
		}
	}
	return -1
}

func (scale myScale) remove(pos int, amount int) []string {
	return append(scale[:pos], scale[pos+amount:]...)
}

func toUpper(tonic string) string {
	newTonic := strings.ToUpper(string(tonic[0]))
	if len(tonic) > 1 {
		newTonic += string(tonic[1])
	}
	return newTonic
}
