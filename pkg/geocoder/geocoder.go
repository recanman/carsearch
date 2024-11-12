// made by recanman
// The stupidest abstraction I've ever made. Too tired to refactor
package geocoder

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/openstreetmap"
)

var g geo.Geocoder

func Initialize() {
	g = openstreetmap.Geocoder()
}

func Geocode(address string) (*geo.Location, error) {
	return g.Geocode(address)
}
