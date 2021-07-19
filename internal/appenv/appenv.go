package appenv

import (
	"github.com/adalekin/otusk8s3/internal/repository/reporegistry"
	"github.com/unrolled/render"
)

type Config struct {
	Port       int    `env:"PORT" envDefault:"80"`
	DbHost     string `env:"DB_HOST"`
	DbPort     int    `env:"DB_PORT" envDefault:"5432"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbName     string `env:"DB_NAME"`
}

type AppEnv struct {
	Config       Config
	RepoRegistry reporegistry.IRepoRegistry
	Render       *render.Render
}
