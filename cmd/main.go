package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rifkiystark/portfolios-api/cmd/ipr"
	"github.com/rifkiystark/portfolios-api/cmd/project"
	"github.com/rifkiystark/portfolios-api/config"
	"github.com/rifkiystark/portfolios-api/graph"
	"github.com/rifkiystark/portfolios-api/internal/database"
	"github.com/rifkiystark/portfolios-api/internal/imagekit"
)

func main() {
	var mb int64 = 1 << 20
	var cfg = config.Get()
	db := database.Connect(cfg.MongoDBUser, cfg.MongoDBPassword, cfg.MongoDBHost, cfg.MongoDBName)
	ik := imagekit.InitImageKit(cfg.ImageKitPublicKey, cfg.ImageKitPrivateKey, cfg.ImageKitURL)

	projectRepository := project.NewRepository(db)
	projectService := project.NewService(projectRepository, ik)

	iprRepository := ipr.NewRepository(db)
	iprService := ipr.NewService(iprRepository, ik)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{ProjectService: projectService, IPRService: iprService}}))
	srv.AddTransport(transport.MultipartForm{
		MaxMemory:     32 * mb,
		MaxUploadSize: 50 * mb,
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.HTTPPort)
	log.Fatal(http.ListenAndServe(":"+cfg.HTTPPort, nil))
}
