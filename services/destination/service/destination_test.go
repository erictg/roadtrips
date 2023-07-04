package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/erictg/roadtrips/services/destination/service/helper"
	"github.com/stretchr/testify/require"
	"googlemaps.github.io/maps"
)

const (
	metersPerMile float64 = 1609.34
)

func TestService(t *testing.T) {
	req := require.New(t)

	svc, err := NewDestinationService("CHANGE_ME")
	req.NoError(err)
	req.NotNil(svc)

	latLon, err := svc.GetLatLon(context.Background(), "1920 Georgetown Dr. Sewickley PA, 15143")
	req.NoError(err)
	req.NotNil(latLon)

	options, err := svc.GetRandomDestinations(context.Background(), latLon, 50*metersPerMile, 100*metersPerMile, maps.PlaceTypeRestaurant)
	req.NoError(err)
	req.NotEmpty(options)

	for _, opt := range options {
		loc := opt.Vicinity
		if opt.FormattedAddress != "" {
			loc = opt.FormattedAddress
		}

		dist := helper.Distance(latLon, &opt.Geometry.Location)
		fmt.Printf("address=%s | dis=%v\n", loc, dist/metersPerMile)
	}

	fmt.Printf("TRAVEL TIME TO %s (%s)", options[0].Name, options[0].FormattedAddress)

	dur, err := svc.GetTravelTime(context.Background(), latLon, &options[0].Geometry.Location)
	req.NoError(err)
	fmt.Printf("TRAVEL TIME: %s", dur.String())
}
