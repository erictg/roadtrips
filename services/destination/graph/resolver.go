package graph

import "github.com/erictg/roadtrips/services/destination/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client *service.DestinationService
}

func NewResolver(client *service.DestinationService) *Resolver {
	return &Resolver{
		client: client,
	}
}
