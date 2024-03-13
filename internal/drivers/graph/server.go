package graph

import (
	"devices_crud/internal/devices"
	"devices_crud/internal/drivers/graph/generated"
	"devices_crud/internal/drivers/graph/resolver"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

type GraphQLServer struct {
	server *handler.Server
	port   string
}

func RunServer(deviceDeps *devices.DependencyTree) *GraphQLServer {
	port := "8080"
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{
		DeviceService: deviceDeps.DeviceSerivce,
		Logger:        deviceDeps.Logger,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Fatal(http.ListenAndServe(":"+port, nil))
	return &GraphQLServer{
		server: srv,
		port:   port,
	}
}

func (gqls *GraphQLServer) Run() {
	log.Fatal(http.ListenAndServe(":"+gqls.port, nil))
}
