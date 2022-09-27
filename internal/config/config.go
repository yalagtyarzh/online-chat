package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"os"
	"sync"
)

// Config represents the config of the launched http server
type Config struct {
	LogLevel  string `env:"LOG_LEVEL" envDefault:"debug"`
	PrettyLog bool   `env:"PRETTY_LOG" envDefault:"false"`
	Bind      string `env:"BIND" envDefault:":8080"`
	ProbeBind string `env:"PROBE_BIND" envDefault:":9091"`
	Postgres  struct {
		Username string `env:"POSTGRES_USERNAME" envDefault:"default"`
		Password string `env:"POSTGRES_PASSWORD" envDefault:""`
		Database string `env:"POSTGRES_DATABASE" envDefault:"chat"`
		Host     string `env:"POSTGRES_PORT" envDefault:"localhost:5432"`
	}
	Redis struct {
		Addr     string `env:"REDIS_ADDRESS" envDefault:"localhost:6379"`
		Password string `env:"REDIS_PASSWORD" envDefault:""`
	}
	JWT struct {
		Issuer string `env:"JWT_ISSUER" envDefault:"go-chat"`
		Secret string `env:"JWT_SECRET" envDefault:"secret"`
	}
}

var cfg *Config
var once sync.Once

// InitConfig gets the config from the environment variables and writes the values to the Config structure, returning it
func InitConfig() *Config {
	once.Do(
		func() {
			cfg = &Config{}

			if err := env.Parse(cfg); err != nil {
				fmt.Printf("environment is not OK: %s\n", err)
				os.Exit(1)
			}
		},
	)

	return cfg
}
