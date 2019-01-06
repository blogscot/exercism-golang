package clock

import (
	"fmt"
)

// Clock holds both hours and minutes.
type Clock struct {
	hours   int
	minutes int
}

// New creates a new clock.
func New(hours, minutes int) Clock {
	return Clock{hours, minutes}.normalise()
}

// Add adds minutes to the clock time.
func (clock Clock) Add(minutes int) Clock {
	hours := minutes / 60
	mins := minutes % 60
	return Clock{clock.hours + hours, clock.minutes + mins}.normalise()
}

// Subtract subtracts minutes to the clock time.
func (clock Clock) Subtract(minutes int) Clock {
	hours := minutes / 60
	mins := minutes % 60
	return Clock{clock.hours - hours, clock.minutes - mins}.normalise()
}

// String converts a clock to a string.
func (clock Clock) String() string {
	c := clock.normalise()
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}

func (clock Clock) normalise() Clock {
	newHours := clock.minutes / 60
	minutes := clock.minutes % 60
	if minutes < 0 {
		minutes += 60
		newHours--
	}
	hours := (clock.hours + newHours) % 24
	if hours < 0 {
		hours += 24
	}
	return Clock{hours, minutes}
}
