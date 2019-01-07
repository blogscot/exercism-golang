package space

// Planet represents planets in the solar system
type Planet string

// EarthPeriod the time taken to orbit Sol in seconds
const EarthPeriod = 31557600

// Age calculates a person's age for the given planet
func Age(seconds float64, planet Planet) float64 {
	var ratio float64

	switch planet {
	case "Earth":
		ratio = 1.0
	case "Mercury":
		ratio = 0.2408467
	case "Venus":
		ratio = 0.61519726
	case "Mars":
		ratio = 1.8808158
	case "Jupiter":
		ratio = 11.862615
	case "Saturn":
		ratio = 29.447498
	case "Uranus":
		ratio = 84.016846
	case "Neptune":
		ratio = 164.79132
	}
	return seconds / EarthPeriod / ratio
}
