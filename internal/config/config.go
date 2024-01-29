package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default: "localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default: "4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default: "60s"`
}

func MustLoad() *Config {
	//configPath := os.Getenv("CONFIG_PATH") Разобраться почему там дофига путей и можно ли их снести
	configPath := "/Users/samoylowgennadiy/Desktop/GitHub/url-shortener/config/local.yaml"

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", configPath)
	}

	return &cfg
}
