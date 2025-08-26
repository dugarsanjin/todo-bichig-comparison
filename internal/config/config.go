package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env        string `yaml:"env" env:"ENV" env-required:"true"`
	ModelPath  string `yaml:"model_path" env:"MODEL_PATH" env-required:"true"`
	Server     `yaml:"server" env-required:"true"`
	Image      `yaml:"image" env-required:"true"`
	Processing `yaml:"processing" env-required:"true"`
}

type Server struct {
	Host         string        `yaml:"host" env:"SERVER_HOST" env-required:"true"`
	Port         string        `yaml:"port" env:"SERVER_PORT" env-required:"true"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

type Image struct {
	Width  int `yaml:"width" env:"WIDTH" env-required:"true"`
	Height int `yaml:"height" env:"HEIGHT" env-required:"true"`
}

type Processing struct {
	Timeout time.Duration `yaml:"timeout" env:"PROCESSING_TIMEOUT" env-required:"true"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist at path: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	return &cfg
}
