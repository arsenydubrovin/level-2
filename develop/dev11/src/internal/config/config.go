package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	hostEnv   = "HOST"
	portEnv   = "PORT"
	dbPathEnv = "DB_PATH"
)

type Config struct {
	host   string
	port   string
	dbPath string
}

func MustLoad(path string) *Config {
	var cfg Config

	err := godotenv.Load(path)
	if err != nil {
		log.Fatal(err)
	}

	host, err := getEnvVar(hostEnv)
	if err != nil {
		log.Fatal(err)
	}
	cfg.host = host

	port, err := getEnvVar(portEnv)
	if err != nil {
		log.Fatal(err)
	}
	cfg.port = port

	dbPath, err := getEnvVar(dbPathEnv)
	if err != nil {
		log.Fatal(err)
	}
	cfg.dbPath = dbPath

	return &cfg
}

func getEnvVar(envVar string) (string, error) {
	if value, exists := os.LookupEnv(envVar); exists {
		return value, nil
	}

	return "", fmt.Errorf("env variable %s not found", envVar)
}

func (c *Config) Address() string {
	return fmt.Sprintf("%s:%s", c.host, c.port)
}

func (c *Config) DBPath() string {
	return c.dbPath
}
