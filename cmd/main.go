package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rifkiystark/portfolios-api/cmd/certificate"
	"github.com/rifkiystark/portfolios-api/cmd/ipr"
	"github.com/rifkiystark/portfolios-api/cmd/project"
	"github.com/rifkiystark/portfolios-api/config"
	"github.com/rifkiystark/portfolios-api/graph"
	"github.com/rifkiystark/portfolios-api/internal/database"
	"github.com/rifkiystark/portfolios-api/internal/imagekit"
	"github.com/rs/cors"
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

	certificateRepository := certificate.NewRepository(db)
	certificateService := certificate.NewService(certificateRepository, ik)

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{
					ProjectService:     projectService,
					IPRService:         iprService,
					CertificateService: certificateService,
				},
			},
		),
	)

	srv.AddTransport(transport.MultipartForm{
		MaxMemory:     32 * mb,
		MaxUploadSize: 50 * mb,
	})

	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.HTTPPort)
	log.Fatal(http.ListenAndServe(":"+cfg.HTTPPort, nil))
}
