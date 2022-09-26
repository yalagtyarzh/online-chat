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
