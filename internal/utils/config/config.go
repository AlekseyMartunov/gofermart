package config

import (
	"flag"
	"os"
)

type Config struct {
	runAddress           string `env:"RUN_ADDRESS"`
	dataBaseDSN          string `env:"DATABASE_URI"`
	accrualSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
}

func New() *Config {
	return &Config{}
}

func (c *Config) ParseFlags() {
	flag.StringVar(&c.runAddress, "a", "",
		"address of server")

	flag.StringVar(&c.dataBaseDSN, "d", "",
		"data base dsn")

	flag.StringVar(&c.accrualSystemAddress, "r", "",
		"address of bonuses accrual system")

	if envRunAddr := os.Getenv("RUN_ADDRESS"); envRunAddr != "" {
		c.runAddress = envRunAddr
	}

	if envDSN := os.Getenv("DATABASE_URI"); envDSN != "" {
		c.dataBaseDSN = envDSN
	}

	if accSystem := os.Getenv("RUN_ADDRESS"); accSystem != "" {
		c.accrualSystemAddress = accSystem
	}
}

func (c *Config) RunAddr() string {
	return c.runAddress
}

func (c *Config) DSN() string {
	return c.dataBaseDSN
}

func (c *Config) AccSystemAddr() string {
	return c.accrualSystemAddress
}