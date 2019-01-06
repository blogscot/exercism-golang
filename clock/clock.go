package clock

import (
	"fmt"
)

// Clock holds both hours and minutes.
type Clock struct {
	minutes int
}

// New creates a new clock.
func New(hours, minutes int) Clock {
	return Clock{hours*60 + minutes}.normalise()
}

// Add adds minutes to the clock time.
func (clock Clock) Add(minutes int) Clock {
	return Clock{clock.minutes + minutes}.normalise()
}

// Subtract subtracts minutes to the clock time.
func (clock Clock) Subtract(minutes int) Clock {
	return Clock{clock.minutes - minutes}.normalise()
}

// String converts a clock to a string.
func (clock Clock) String() string {
	hours := (clock.minutes / 60) % 24
	minutes := clock.minutes % 60
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

func (clock Clock) normalise() Clock {
	minutes := clock.minutes
	for minutes <= 0 {
		minutes += 24 * 60
	}
	return Clock{minutes % (24 * 60)}
}
