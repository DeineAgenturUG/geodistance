/*
Package geodistance calculates various distances measured along the surface of the earth.
All formulaes in this package calculate distances between points which are
defined by geographical coordinates in terms of latitude and longitude.
The usecase for this is knowing it's own latitude and longitude and asking for distance and
magnetic compass heading to a given target latitude and longitude.

Note that the longitude is singular at the Poles. Calculations that are sufficiently
accurate for other positions, may be inaccurate or completely wrong at or near the Poles.
Also the discontinuity at the 180° meridian is not taken into calculations.

*/
package geodistance

import (
	"errors"
	"math"
)

// GeoStruct inputs use lattitide and longitude in decimal form only-
type GeoStruct struct {
	SrcLatitude, SrcLongitude, TgtLatitude, TgtLongitude float64
	Precision                                            int
}

// HeadingStruct returns calcuated heading information
type HeadingStruct struct {
	DistanceMeter                                float64
	DistanceKiloMeter                            float64
	HeadingDegrees                               float64
	UncorrectedDistanceMeter                     float64
	DistanceNauticalMile                         float64
	DistanceMile                                 float64
	DistanceFoot                                 float64
	DistanceYard                                 float64
	NorthBound, SouthBound, EastBound, WestBound bool
}

const earthRadiusInMeter float64 = 6371009
const internationalNauticalMileInMeter float64 = 1852
const footInMeter float64 = 0.3048
const yardInMeter float64 = 0.9144
const mileInMeter float64 = 1609.344
const defaultPrecision int = 0

// PointStruct takes input positions given in degrees, minutes and seconds
// and a Position string that is one of {"North"|"East"|"South"|"West""}.
// If quiteErrorOutput is set to true error messages is suppressed.
type PointStruct struct {
	Position string
	Degrees  float64
	Minutes  float64
	Seconds  float64
}

// SphericalEarthProjection use a spherical earth to project distance to a flat plan
func SphericalEarthProjection(geo GeoStruct) (hs HeadingStruct, err error) {

	if geo.Precision == 0 {
		geo.Precision = defaultPrecision
	}

	var srcLat = degreesToRadians(geo.SrcLatitude)
	var tgtLat = degreesToRadians(geo.TgtLatitude)
	var adlat = math.Abs(tgtLat - srcLat)

	// Ignore special case moving over ±90 degrees (poles)
	if tgtLat > srcLat {
		hs.NorthBound = true
	}
	if tgtLat < srcLat {
		hs.SouthBound = true
	}

	var srcLong = degreesToRadians(geo.SrcLongitude)
	var tgtLong = degreesToRadians(geo.TgtLongitude)
	var adlong = math.Abs(tgtLong - srcLong)

	// Ignore special case moving over 180th meridian (International Date Line)
	if tgtLong > srcLong {
		hs.EastBound = true // The preferred convention, East longitudes is positive
	}
	if tgtLong < srcLong {
		hs.WestBound = true // The preferred convention, West longitudes is negative
	}

	// Target position must be different from current position
	if srcLat == tgtLat && srcLong == tgtLong {
		return hs, errors.New("target position is equal with start position")
	}

	// Correct distance for variation in latitude with lateral position on earth
	hs.UncorrectedDistanceMeter = roundPlus(earthRadiusInMeter*math.Sqrt(adlat*adlat+adlong*adlong), geo.Precision)
	meanLattitude := (srcLat + tgtLat) / 2
	lattitudeCorrectionFactor := math.Cos(meanLattitude)

	// Distance in different units
	hs.DistanceMeter = roundPlus(earthRadiusInMeter*math.Sqrt(adlat*adlat+(lattitudeCorrectionFactor*adlong)*(lattitudeCorrectionFactor*adlong)), geo.Precision)
	hs.DistanceKiloMeter = roundPlus(hs.DistanceMeter/1000.0, geo.Precision)
	hs.DistanceNauticalMile = roundPlus(hs.DistanceMeter/internationalNauticalMileInMeter, geo.Precision)
	hs.DistanceMile = roundPlus(hs.DistanceMeter/mileInMeter, geo.Precision)
	hs.DistanceFoot = roundPlus(hs.DistanceMeter/footInMeter, geo.Precision)
	hs.DistanceYard = roundPlus(hs.DistanceMeter/yardInMeter, geo.Precision)

	// Heading to target position
	dlat := tgtLat - srcLat
	dlong := tgtLong - srcLong
	uncorrectedHeading := math.Atan2(dlong, dlat)
	hs.HeadingDegrees = roundPlus(compassHeading(radiansToDegree(uncorrectedHeading)), geo.Precision)

	return hs, nil
}

// PointsToDegrees converts position from degrees, minutes and seconds to decimal degrees
func PointsToDegrees(ps PointStruct) (position float64, err error) {

	position = ps.Degrees + ps.Minutes/60 + ps.Seconds/3600

	if ps.Position == "South" || ps.Position == "West" {
		position = -1 * position
	}
	return position, nil
}
