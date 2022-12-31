package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/imagekit-developer/imagekit-go"
	"github.com/rifkiystark/portfolios-api/config"
	"github.com/rifkiystark/portfolios-api/database"
	"github.com/rifkiystark/portfolios-api/graph"
)

func main() {
	var mb int64 = 1 << 20
	var cfg = config.Get()
	db := connectToDatabase(cfg)
	ik := initImageKit(cfg)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: db, IK: ik}}))
	srv.AddTransport(transport.MultipartForm{
		MaxMemory:     32 * mb,
		MaxUploadSize: 50 * mb,
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.HTTPPort)
	log.Fatal(http.ListenAndServe(":"+cfg.HTTPPort, nil))
}

func connectToDatabase(cfg *config.Config) *database.DB {
	url := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", cfg.MongoDBUser, cfg.MongoDBPassword, cfg.MongoDBHost)
	return database.Connect(url, cfg.MongoDBName)
}

func initImageKit(cfg *config.Config) *imagekit.ImageKit {
	ik := imagekit.NewFromParams(imagekit.NewParams{
		PublicKey:   cfg.ImageKitPublicKey,
		PrivateKey:  cfg.ImageKitPrivateKey,
		UrlEndpoint: cfg.ImageKitURL,
	})

	return ik
}
