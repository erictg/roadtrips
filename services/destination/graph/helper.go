package graph

import (
	"github.com/erictg/roadtrips/services/destination/graph/model"
	"googlemaps.github.io/maps"
)

func adaptPointToLatLon(pt model.Point) *maps.LatLng {
	return &maps.LatLng{
		Lat: pt.Latitude,
		Lng: pt.Longitude,
	}
}

func adaptDestinationTypeToPlaceType(dt model.DestinationType) maps.PlaceType {
	switch dt {
	case model.DestinationTypeRestaurant:
		return maps.PlaceTypeRestaurant
	default:
		return ""
	}
}

func ptrOf[T any](v T) *T {
	return &v
}

func lengthToMeters(l model.Length) float64 {
	switch l.Unit {
	case model.LengthUnitKilometers:
		return l.Value * 1000
	case model.LengthUnitMiles:
		return l.Value * 1609.344
	default:
		// we shouldn't be able to hit this anyways
		return l.Value
	}
}
