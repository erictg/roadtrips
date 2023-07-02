package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"

	"github.com/joho/godotenv"
	"googlemaps.github.io/maps"
)

const (
	metersPerMile            float64 = 1609.34
	closeEnoughMetersPerMile         = 1609
	randLocationMaxTries             = 100
	distFromRand                     = uint(10 * closeEnoughMetersPerMile)
)

/*
	Calculations are in meters
*/

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	mapsApiKey := os.Getenv("MAPS_KEY")
	if mapsApiKey == "" {
		panic("api key not found")
	}

	fromAddy := os.Getenv("ADDRESS")

	c, err := maps.NewClient(maps.WithAPIKey(mapsApiKey))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	fromLL, err := getFromLatLon(ctx, c, fromAddy)
	if err != nil {
		panic(err)
	}

	log.Printf("lat=%v lon=%v", fromLL.Lat, fromLL.Lng)

	maxMiles := float64(100)
	minMiles := float64(50)

	maxTries := 5

	var results []maps.PlacesSearchResult
	for i := 0; i < maxTries; i++ {
		fmt.Printf("lookup attempt %v\n", i)
		randLoc := generateRandomLocation(fromLL, minMiles*metersPerMile, maxMiles*metersPerMile)

		dist := distance(fromLL, randLoc)
		fmt.Println(randLoc.String())
		fmt.Printf("miles=%v\n", dist/metersPerMile)

		resp, err := c.TextSearch(ctx, &maps.TextSearchRequest{
			Location: randLoc,
			Radius:   distFromRand,
			Query:    "restaurant",
		})
		if err != nil {
			panic(err)
		}

		if len(resp.Results) == 0 {
			continue
		} else {
			results = resp.Results
			break
		}
	}

	// print results
	for _, res := range results {
		distFromMe := distance(fromLL, &res.Geometry.Location)
		log.Printf("name=%s|addy=%s|miles=%v", res.Name, res.FormattedAddress, distFromMe/metersPerMile)
	}

	// get random restautrant from list
	target := results[rand.Intn(len(results)-1)]
	distToTarget := distance(fromLL, &target.Geometry.Location) / metersPerMile
	log.Println("----------")
	log.Println("RANDOM CHOICE:")
	log.Printf("Name: %s", target.Name)
	log.Printf("Address: %s", target.FormattedAddress)
	log.Printf("Distance(miles): %v", distToTarget)
	log.Printf("Cost: %s", costToString(target.PriceLevel))
	log.Printf("Rating: %v", target.Rating)
	log.Printf("TotalRatings: %v", target.UserRatingsTotal)
	log.Printf("Type: %v", target.Types)
	log.Println("----------")
}

func costToString(i int) string {
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

func getFromLatLon(ctx context.Context, c *maps.Client, fromAddy string) (*maps.LatLng, error) {
	fromLocs, err := c.Geocode(ctx, &maps.GeocodingRequest{
		Address: fromAddy,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to geocode address, %w", err)
	}

	if len(fromLocs) == 0 {
		return nil, errors.New("no locations found")
	}

	// assume the first one is right
	return &fromLocs[0].Geometry.Location, nil
}

func pointAtDistance(location *maps.LatLng, radius float64) *maps.LatLng {
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

func generateRandomLocation(location *maps.LatLng, min, max float64) *maps.LatLng {
	for i := 0; i < randLocationMaxTries; i++ {
		fmt.Printf("gen rand loc try %v\n", i)
		rnd := rand.Float64()
		ranDist := math.Sqrt(rnd) * max
		randLoc := pointAtDistance(location, ranDist)

		// hack until i spend the time to figure out the math to do it between to radii
		dist := distance(location, randLoc)
		if dist > min {
			return randLoc
		}
	}

	return nil
}

func distance(ll1 *maps.LatLng, ll2 *maps.LatLng) float64 {
	radlat1 := float64(math.Pi * ll1.Lat / 180)
	radlat2 := float64(math.Pi * ll2.Lat / 180)

	theta := float64(ll1.Lng - ll2.Lng)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	// convert to meters
	return dist * metersPerMile
}
