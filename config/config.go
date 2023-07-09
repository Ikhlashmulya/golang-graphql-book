package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
}

func NewConfig(filenames ...string) *Config {
	err := godotenv.Load(filenames...)
	if err != nil {
		panic(err)
	}
	
	return &Config{}
}

func (config *Config) Get(key string) string {
	return os.Getenv(key)
}