// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Destination interface {
	IsDestination()
	GetLatitude() float64
	GetLongitude() float64
	GetName() string
	GetEstimatedTravelTime() string
	GetWazeDeeplink() string
}

type DestinationFilters struct {
	Type DestinationType `json:"type"`
}

type Length struct {
	Value float64    `json:"value"`
	Unit  LengthUnit `json:"unit"`
}

type Point struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type RandomDestinationWithinRing struct {
	Filters *DestinationFilters `json:"filters,omitempty"`
	Ring    *Ring               `json:"ring"`
}

type RandomDestinationsWithinRing struct {
	Filters *DestinationFilters `json:"filters,omitempty"`
	Ring    *Ring               `json:"ring"`
}

type Restaurant struct {
	Latitude            float64  `json:"latitude"`
	Longitude           float64  `json:"longitude"`
	Name                string   `json:"name"`
	EstimatedTravelTime string   `json:"estimatedTravelTime"`
	WazeDeeplink        string   `json:"wazeDeeplink"`
	Rating              *float64 `json:"rating,omitempty"`
	NumRatings          *int     `json:"numRatings,omitempty"`
	Types               []string `json:"types"`
	Hours               []string `json:"hours"`
	IconURL             *string  `json:"iconURL,omitempty"`
}

func (Restaurant) IsDestination()                      {}
func (this Restaurant) GetLatitude() float64           { return this.Latitude }
func (this Restaurant) GetLongitude() float64          { return this.Longitude }
func (this Restaurant) GetName() string                { return this.Name }
func (this Restaurant) GetEstimatedTravelTime() string { return this.EstimatedTravelTime }
func (this Restaurant) GetWazeDeeplink() string        { return this.WazeDeeplink }

type Ring struct {
	Center      *Point  `json:"center"`
	InnerRadius *Length `json:"innerRadius"`
	OuterRadius *Length `json:"outerRadius"`
}

type DestinationType string

const (
	DestinationTypeRestaurant DestinationType = "Restaurant"
)

var AllDestinationType = []DestinationType{
	DestinationTypeRestaurant,
}

func (e DestinationType) IsValid() bool {
	switch e {
	case DestinationTypeRestaurant:
		return true
	}
	return false
}

func (e DestinationType) String() string {
	return string(e)
}

func (e *DestinationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DestinationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DestinationType", str)
	}
	return nil
}

func (e DestinationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type LengthUnit string

const (
	LengthUnitMiles      LengthUnit = "Miles"
	LengthUnitKilometers LengthUnit = "Kilometers"
)

var AllLengthUnit = []LengthUnit{
	LengthUnitMiles,
	LengthUnitKilometers,
}

func (e LengthUnit) IsValid() bool {
	switch e {
	case LengthUnitMiles, LengthUnitKilometers:
		return true
	}
	return false
}

func (e LengthUnit) String() string {
	return string(e)
}

func (e *LengthUnit) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LengthUnit(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LengthUnit", str)
	}
	return nil
}

func (e LengthUnit) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
