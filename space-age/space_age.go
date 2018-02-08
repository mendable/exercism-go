// Given an age in seconds, calculate how old someone would be on:
//
//    - Earth: orbital period 365.25 Earth days //, or 31557600 seconds
//    - Mercury: orbital period 0.2408467 Earth years
//    - Venus: orbital period 0.61519726 Earth years
//    - Mars: orbital period 1.8808158 Earth years
//    - Jupiter: orbital period 11.862615 Earth years
//    - Saturn: orbital period 29.447498 Earth years
//    - Uranus: orbital period 84.016846 Earth years
//    - Neptune: orbital period 164.79132 Earth years
//
package space

const secondsPerEarthYear float64 = 31557600 // 60 * 60 * 24 * 365.25

type Planet string

var Periods = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func (p *Planet) period() float64 {
	return Periods[*p]
}

// math.Round() coming in go 1.10, not implementing a buggy or messy version here.
// Tests pass as-is.
// https://github.com/golang/go/issues/20100
// https://www.cockroachlabs.com/blog/rounding-implementations-in-go/
func round(num float64) float64 {
	return num
}

func Age(seconds float64, planet Planet) float64 {
	return seconds / (secondsPerEarthYear * planet.period())
}
