package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Env        string `env:"TODO_ENV" envDefault:"dev"`
	Port       int    `env:"PORT" envDefault:"80"`
	DBHost     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DBPort     int    `env:"DB_PORT" envDefault:"33306"`
	DBUser     string `env:"DBUser" envDefault:"todo"`
	DBPassword string `env:"DBPassword" envDefault:"todo"`
	DBName     string `env:"DBName" envDefault:"todo"`
	RedisHost  string `env:"REDIS_HOST" envDefault:"127.0.0.1"`
	RedisPort  int    `env:"REDIS_PORT" envDegault:"36379"`
}

func New() (*Config, error) {

	c := &Config{}
	if err := env.Parse(c); err != nil {
		return nil, err
	}

	return c, nil
}
