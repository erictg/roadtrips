package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/erictg/roadtrips/services/destination/service/helper"
	"googlemaps.github.io/maps"
)

// const (
// 	metersPerMile            float64 = 1609.34
// 	closeEnoughMetersPerMile         = 1609
// 	randLocationMaxTries             = 100
// )

type DestinationService struct {
	client *maps.Client
}

func NewDestinationService(mapsApiKey string) (*DestinationService, error) {
	c, err := maps.NewClient(maps.WithAPIKey(mapsApiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to init maps client, %w", err)
	}

	return &DestinationService{
		client: c,
	}, nil
}

func (d *DestinationService) GetRandomDestinations(ctx context.Context, center *maps.LatLng, innerRadius, outerRadius float64, destType maps.PlaceType, keywords ...string) ([]maps.PlacesSearchResult, error) {
	randomLocation := helper.GenerateRandomLocation(center, innerRadius, outerRadius)

	keyword := ""
	if len(keyword) != 0 {
		keyword = strings.Join(keywords, " ")
	}

	resp, err := d.client.NearbySearch(ctx, &maps.NearbySearchRequest{
		Location: randomLocation,
		RankBy: maps.RankByDistance,
		Type: destType,
		Keyword: keyword,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to do a nearby search, %w", err)
	}

	return resp.Results, nil
}

func (d *DestinationService) GetRandomDestination(ctx context.Context, center *maps.LatLng, innerRadius, outerRadius float64, destType maps.PlaceType, keywords ...string) (maps.PlacesSearchResult, error) {
	options, err := d.GetRandomDestinations(ctx, center, innerRadius, outerRadius, destType, keywords...)
	if err != nil {
		return maps.PlacesSearchResult{}, err
	}

	switch len(options) {
	case 0:
		return maps.PlacesSearchResult{}, errors.New("no destinations found")
	case 1:
		return options[0], nil
	default:
		return options[rand.Intn(len(options)-1)], nil
	}
}

func (d *DestinationService) GetTravelTime(ctx context.Context, start, target *maps.LatLng) (time.Duration, error) {
	routes, _, err := d.client.Directions(ctx, &maps.DirectionsRequest{
		Origin: start.String(),
		Destination: target.String(),
		Mode: maps.TravelModeDriving,
		Optimize: true,
	})
	if err != nil {
		return -1, fmt.Errorf("failed to get directions, %w", err)		
	}

	if len(routes) == 0 {
		return -1, fmt.Errorf("unable to compute duration of drive, %w", err)
	}

	// aggregate time of each leg
	var travelTime time.Duration
	for _, leg := range routes[0].Legs {
		if leg.DurationInTraffic != 0 {
			travelTime += leg.DurationInTraffic
		} else {
			travelTime += leg.Duration
		}
	}

	return travelTime, nil
}

func (d *DestinationService) GetLatLon(ctx context.Context, address string) (*maps.LatLng, error) {
	resp, err := d.client.Geocode(ctx, &maps.GeocodingRequest{
		Address: address,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to geocode address, %w", err)
	}

	if len(resp) == 0 {
		return nil, errors.New("no results found")
	}

	return &resp[0].Geometry.Location, nil
}