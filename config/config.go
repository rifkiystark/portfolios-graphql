package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// HTTP Server
	HTTPPort string `envconfig:"HTTP_PORT" default:"8080"`

	// DB is the database configuration
	MongoDBHost     string `envconfig:"MONGODB_HOST" default:"localhost"`
	MongoDBUser     string `envconfig:"MONGODB_USER" default:""`
	MongoDBPassword string `envconfig:"MONGODB_PASSWORD" default:""`
	MongoDBName     string `envconfig:"MONGODB_NAME" default:"portfolios"`

	// ImageKit is the imagekit configuration
	ImageKitPublicKey  string `envconfig:"IMAGEKIT_PUBLIC_KEY" default:""`
	ImageKitPrivateKey string `envconfig:"IMAGEKIT_PRIVATE_KEY" default:""`
	ImageKitURL        string `envconfig:"IMAGEKIT_URL" default:""`
}

func Get() *Config {
	var C Config
	envconfig.MustProcess("", &C)
	return &C
}
