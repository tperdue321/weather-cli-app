package Flags

import (
	"flag"
)

type Location struct {
	City        string
	State       string
	Zipcode     string
	CountryCode string
	Latitude    string
	Longitude   string
}

func SetFlags() *Location {
	location := new(Location)
	flag.StringVar(&location.City, "city", "", "A city")
	flag.StringVar(&location.State, "state", "", "A state")
	flag.StringVar(&location.Zipcode, "zipcode", "", "A zipcode")
	flag.StringVar(&location.CountryCode, "c", "us", "A Country Code")
	flag.StringVar(&location.Latitude, "lat", "", "Latitude")
	flag.StringVar(&location.Longitude, "long", "", "Longitude")
	flag.Parse()
	return location
}
