// Implements a clock that handles times without dates. You can add and
// subtract minutes to and from the clock. Two clocks that represent the same
// time will be equal to each other.
package clock

import "fmt"

type Clock struct {
	hour   int // hours
	minute int // minutes
}

// New returns a build instance of Clock with the time set correctly.
func New(hour, minute int) Clock {
	c := Clock{hour: 0, minute: 0}

	if hour > 0 {
		c = c.Add(60 * hour)
	} else if hour < 0 {
		c = c.Subtract(60 * hour * -1)
	}

	if minute > 0 {
		c = c.Add(minute)
	} else if minute < 0 {
		c = c.Subtract(minute * -1)
	}

	return c
}

// Adds minutes to Clock, and returns amended clock instance.
func (c Clock) Add(minutes int) Clock {
	totalMinutes := (c.hour * 60) + c.minute + minutes

	c.minute = totalMinutes % 60
	c.hour = totalMinutes / 60

	if c.hour >= 24 {
		c.hour = c.hour % 24
	}

	return c
}

// Subtracts minutes from Clock, and returns and amended clock instance.
func (c Clock) Subtract(minutes int) Clock {
	totalMinutes := ((c.hour * 60) + c.minute) - minutes

	if totalMinutes < 0 { // TODO: Is there a more elegant way to write this?
		totalMinutes = totalMinutes * -1                             // -1500 => 1500
		totalMinutes = totalMinutes - (1440 * (totalMinutes / 1440)) // 1500 - (1500 / 1440) = 60
		totalMinutes = (totalMinutes * -1) + 1440                    // (60 * -1) + 1400 = 1380 (23:00)
	}

	c.minute = totalMinutes % 60
	c.hour = totalMinutes / 60

	return c
}

// String returns a string representation for clock, in HH:MM format.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
