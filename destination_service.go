package destination

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/erictg/roadtrips/services/destination/graph"
	"github.com/erictg/roadtrips/services/destination/service"
)

func init() {
	hndlr, err := NewGQLHandler()
	if err != nil {
		log.Fatalf("failed to int gql handler, %v", err.Error())
	}

	functions.HTTP("DestinationGQL", hndlr.serveGQL)
}

type GQLHandler struct {
	server *handler.Server
}

func NewGQLHandler() (*GQLHandler, error) {
	mapsKey := os.Getenv("MAPS_API_KEY")
	if mapsKey == "" {
		return nil, errors.New("maps key not specified")
	}

	svc, err := service.NewDestinationService(mapsKey)
	if err != nil {
		return nil, fmt.Errorf("failed to init destination svc, %w", err)
	}

	res := graph.NewResolver(svc)

	return &GQLHandler{
		server: handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: res})),
	}, nil
}

func (g *GQLHandler) serveGQL(w http.ResponseWriter, r *http.Request) {
	log.Println("handler")
	g.server.ServeHTTP(w, r)
}
