package geodistance

import (
	"fmt"
	"testing"
)

func TestSphericalEarthProjection(t *testing.T) {
	var gs GeoStruct
	var want, got HeadingStruct
	var err error

	want.DistanceMeter = 123945
	want.DistanceKiloMeter = 124
	want.HeadingDegrees = 45
	want.UncorrectedDistanceMeter = 157254
	want.DistanceNauticalMile = 67
	want.DistanceMile = 77
	want.DistanceFoot = 406644
	want.DistanceYard = 135548
	want.NorthBound = true
	want.SouthBound = false
	want.EastBound = true
	want.WestBound = false

	gs.SrcLatitude = 60.0
	gs.SrcLongitude = 19.0
	gs.TgtLatitude = 61.0
	gs.TgtLongitude = 20.0

	if got, err = SphericalEarthProjection(gs); err != nil {
		fmt.Println("SphericalEarthProjection failed:", err)
	} else {
		if got != want {
			t.Errorf("Heading data differs: want %+v\n but got %+v\n degreees", want, got)
		}
	}

	gs.SrcLatitude = 1.0
	gs.SrcLongitude = 1.0
	gs.TgtLatitude = -1.0
	gs.TgtLongitude = -1.0

	want.DistanceMeter = 314507
	want.DistanceKiloMeter = 315
	want.HeadingDegrees = 225
	want.UncorrectedDistanceMeter = 314507
	want.DistanceNauticalMile = 170
	want.DistanceMile = 195
	want.DistanceFoot = 1.031847e+06
	want.DistanceYard = 343949
	want.NorthBound = false
	want.SouthBound = true
	want.EastBound = false
	want.WestBound = true

	if got, err = SphericalEarthProjection(gs); err != nil {
		fmt.Println("SphericalEarthProjection failed:", err)
	} else {
		if got != want {
			t.Errorf("Heading data differs: want %+v\n but got %+v\n degreees", want, got)
		}
	}

	gs.SrcLatitude = 10.0
	gs.SrcLongitude = 15.0
	gs.TgtLatitude = 10.0
	gs.TgtLongitude = 15.0

	want.DistanceMeter = 0
	want.DistanceKiloMeter = 0
	want.HeadingDegrees = 0
	want.UncorrectedDistanceMeter = 0
	want.DistanceNauticalMile = 0
	want.DistanceMile = 0
	want.DistanceFoot = 0
	want.DistanceYard = 0
	want.NorthBound = false
	want.SouthBound = false
	want.EastBound = false
	want.WestBound = false

	got, _ = SphericalEarthProjection(gs)
	if got != want {
		t.Errorf("Heading data differs: want %+v\n but got %+v\n degreees", want, got)
	}

}
