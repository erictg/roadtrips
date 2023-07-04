package helper

import (
	"fmt"
	"math"
	"math/rand"

	"googlemaps.github.io/maps"
)

const (
	metersPerMile float64 = 1609.344
)

func PointAtDistance(location *maps.LatLng, radius float64) *maps.LatLng {
	y0 := location.Lat
	x0 := location.Lng
	rd := radius / 111300
	u := rand.Float64()
	v := rand.Float64()
	w := rd * math.Sqrt(u)
	t := 2 * math.Pi * v
	x := w * math.Cos(t)
	y := w * math.Sin(t)

	return &maps.LatLng{
		Lat: y + y0,
		Lng: x + x0,
	}
}

func Distance(ll1 *maps.LatLng, ll2 *maps.LatLng) float64 {
	radlat1 := float64(math.Pi * ll1.Lat / 180)
	radlat2 := float64(math.Pi * ll2.Lat / 180)

	theta := float64(ll1.Lng - ll2.Lng)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	// return (math.Acos(dist) * 180 / math.Pi) * 60 * metersPerMile * 1.1515
	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	// convert to meters
	return dist * metersPerMile
}

func GenerateRandomLocation(location *maps.LatLng, min, max float64) *maps.LatLng {
	for {
		rnd := rand.Float64()
		ranDist := math.Sqrt(rnd) * max
		randLoc := PointAtDistance(location, ranDist)

		// hack until i spend the time to figure out the math to do it between to radii
		dist := Distance(location, randLoc)
		if dist > min {
			return randLoc
		}
	}
}

func CostToString(i int) string {
	switch maps.PriceLevel(fmt.Sprintf("%v", i)) {
	case maps.PriceLevelFree:
		return "Free"
	case maps.PriceLevelInexpensive:
		return "Inexpensive"
	case maps.PriceLevelModerate:
		return "Moderate"
	case maps.PriceLevelExpensive:
		return "Expensive"
	case maps.PriceLevelVeryExpensive:
		return "Very Expensive"
	default:
		return "UNKNOWN"
	}
}
