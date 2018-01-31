package geodistance

import (
	"math"
	"testing"
)

// Simple singel test
func TestDegreesToRadians(t *testing.T) {
	degree := 180.0
	want := math.Pi
	got := degreesToRadians(degree)
	if got != want {
		t.Errorf("Degree; %f want %f but got %f radians", degree, want, got)
	}

}

// Table driven test
func TestDegreesToRadiansTable(t *testing.T) {
	var tests = []struct {
		degree float64
		want   float64
	}{
		{270.0, 3 * math.Pi / 2},
		{0, 0},
		{360.0, 2 * math.Pi},
		{90.0, math.Pi / 2},
	}
	for _, test := range tests {
		got := degreesToRadians(test.degree)
		if got != test.want {
			t.Errorf("Degree: %f want %f but got %f radians", test.degree, test.want, got)
		}
	}
}

// Simple singel test
func TestRadiansToDegree(t *testing.T) {
	radians := math.Pi
	want := 180.0
	got := radiansToDegree(radians)
	if got != want {
		t.Errorf("Radians; %f want %f but got %f degreees", radians, want, got)
	}

}

// Table driven test
func TestRadiansToDegreeTable(t *testing.T) {
	var tests = []struct {
		radians float64
		want    float64
	}{
		{3 * math.Pi / 2, 270.0},
		{0, 0},
		{2 * math.Pi, 360.0},
		{math.Pi / 2, 90.0},
	}
	for _, test := range tests {
		got := radiansToDegree(test.radians)
		if got != test.want {
			t.Errorf("Radians: %f want %f but got %f degrees", test.radians, test.want, got)
		}
	}
}

// Table driven test
func TestCompassHeadingTable(t *testing.T) {
	var tests = []struct {
		degrees float64
		want    float64
	}{
		{270.0, 270.0},
		{0, 0},
		{-90.0, 270.0},
		{90.0, 90.0},
		{-180.0, 180.0},
	}
	for _, test := range tests {
		got := compassHeading(test.degrees)
		if got != test.want {
			t.Errorf("Degrees: %f want %f but got %f degrees", test.degrees, test.want, got)
		}
	}
}

// Simple singel test
func TestPointsToDegrees(t *testing.T) {
	var p PointStruct
	var err error

	p.Degrees = 117.0
	p.Minutes = 29.0
	p.Seconds = 50.5
	p.Position = "North"

	want := 117.497361
	var got float64
	if got, err = PointsToDegrees(p); err != nil {
		gotRounded := roundPlus(got, 6)
		if gotRounded != want {
			t.Errorf("Degrees: want %f but got %f degreees", want, gotRounded)
		}
	}

	p.Degrees = 117.0
	p.Minutes = 29.0
	p.Seconds = 50.5
	p.Position = "South"

	want = -117.497361

	if got, err := PointsToDegrees(p); err != nil {
		gotRounded := roundPlus(got, 6)
		if gotRounded != want {
			t.Errorf("Degrees: want %f but got %f degreees", want, gotRounded)
		}
	}

}
