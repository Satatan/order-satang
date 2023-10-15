package configs

import (
	"time"
)

type PostgresConfig struct {
	Port        int           `envconfig:"POSTGRES_PORT"`
	Host        string        `envconfig:"POSTGRES_HOST"`
	User        string        `envconfig:"POSTGRES_USER"`
	Pass        string        `envconfig:"POSTGRES_PASS"`
	Database    string        `envconfig:"POSTGRES_DATABASE"`
	Debug       bool          `envconfig:"POSTGRES_DEBUG"`
	Maxidle     int           `envconfig:"POSTGRES_MAXIDLE"`
	Maxopen     int           `envconfig:"POSTGRES_MAXOPEN"`
	Maxlifetime time.Duration `envconfig:"POSTGRES_MAXLIFETIME"`
}
