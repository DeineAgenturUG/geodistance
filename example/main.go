package main

import "github.com/berrak/geodistance"

import "fmt"

func main() {

	var gs geodistance.GeoStruct
	gs.SrcLatitude = 60.567
	gs.SrcLongitude = 19.011
	gs.TgtLatitude = 61.123
	gs.TgtLongitude = 20.789

	fmt.Println("\nGiven latitude and longitude positions:")
	fmt.Printf("%+v\n", gs)

	var hs geodistance.HeadingStruct
	var err error
	if hs, err = geodistance.SphericalEarthProjection(gs); err != nil {
		fmt.Println("SphericalEarthProjection failed:", err)
	} else {
		fmt.Println("\nCalculated distance and heading structure, round result to default (0) decimal precision:")
		fmt.Printf("%+v\n", hs)
	}

	fmt.Println("\nConvert from degrees, minutes and seconds to decimal degrees. Then get distance and heading")

	var ps geodistance.PointStruct
	ps.Position = "North"
	ps.Degrees = 59.0
	ps.Minutes = 18.1
	ps.Seconds = 0

	fmt.Println("Given latitude position:")
	fmt.Printf("%+v\n", ps)
	gs.TgtLatitude, _ = geodistance.PointsToDegrees(ps)

	ps.Position = "East"
	ps.Degrees = 18.0
	ps.Minutes = 36.9
	ps.Seconds = 0

	fmt.Println("Given longitude position:")
	fmt.Printf("%+v\n", ps)
	gs.TgtLongitude, _ = geodistance.PointsToDegrees(ps)

	gs.Precision = 2
	if hs, err = geodistance.SphericalEarthProjection(gs); err != nil {
		fmt.Println("SphericalEarthProjection failed:", err)
	} else {
		fmt.Println("\nCalculated (rounded result to two decimals) distance and heading:")
		fmt.Printf("%+v\n", hs)
	}

}
