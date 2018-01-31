package geodistance

import (
	"math"
)

/*
////////////////////////////////////////
// Internal package  utility functions
////////////////////////////////////////
*/

func degreesToRadians(degree float64) float64 {
	return degree * math.Pi / 180
}

func radiansToDegree(radians float64) float64 {
	return radians * 180 / math.Pi
}

func compassHeading(degrees float64) float64 {
	if degrees < 0 {
		return degrees + 360
	}
	return degrees
}

// https://gist.github.com/DavidVaini/1030
func round(f float64) float64 {
	return math.Floor(f + .5)
}

// https://gist.github.com/DavidVaini/1030
func roundPlus(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return round(f*shift) / shift
}
