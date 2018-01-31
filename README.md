# geodistance

[![Build Status](https://travis-ci.org/berrak/geodistance.svg?branch=master)](https://travis-ci.org/berrak/geodistance)

The geodistance package calculates various distances measured along the surface of the earth.

All formulaes in this package calculate distances between points which are defined by geographical coordinates in terms of latitude and longitude.

The usecase for this is knowing it's own latitude and longitude and asking for distance and magnetic compass heading to a given target latitude and longitude.

## Usage

Install the library first.

	go get github.com/berrak/geodistance

In your program, all your need to do is to import the package.

Give input origin latitude and longitude, in decimal degrees, and target latitude and longitude with sign to *Geostruct*.

Call the *SphericalEarthProjection* with *Geostruct*.

Calculated distance and compass heading to target is returned in the *HeadingStruct*.

~~~ go
import "github.com/berrak/geodistance"
import "fmt"

// ...

  // Input struct
  var gs geodistance.GeoStruct
  gs.SrcLatitude = 60.567
  gs.SrcLongitude = 19.011
  gs.TgtLatitude = 61.123
  gs.TgtLongitude = 20.789
  
  // Output struct
  var hs geodistance.HeadingStruct
  var err error
  if hs, err = geodistance.SphericalEarthProjection(gs); err != nil {
    fmt.Println("SphericalEarthProjection failed:", err)
  } else {
    fmt.Println("\nCalculated distance and heading structure:")
    fmt.Printf("%+v\n", hs)
  }
  
~~~

There is also a working example under [example](https://github.com/berrak/geodistance/tree/master/example) directory.
