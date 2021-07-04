package configs

import "github.com/kelseyhightower/envconfig"

type Config struct {
	// db configuration
	DBHost                  string `envconfig:"DB_HOST" default:""`
	DBPort                  string `envconfig:"DB_PORT" default:"3306"`
	DBUserName              string `envconfig:"DB_USERNAME" default:""`
	DBName                  string `envconfig:"DB_NAME" default:""`
	DBPass                  string `envconfig:"DB_PASS" default:""`
	DBLogMode               bool   `envconfig:"DB_LOG_MODE" default:"true"`
	DBMaxIdleConnection     int    `envconfig:"DB_MAX_IDLE_CONNECTION" default:"8"`
	DBMaxOpenConnection     int    `envconfig:"DB_MAX_OPEN_CONNECTION" default:"11"`
	DBMaxLifetimeConnection int    `envconfig:"DB_MAX_LIFETIME_CONNECTION" default:"10"` // in minutes
}

// Get to get defined configuration
func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
