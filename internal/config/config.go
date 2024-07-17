package config

import (
	"github.com/caarlos0/env/v10"
	"log"
	"os"
	"time"
)

type Config struct {
	Env     string `env:"env" envDefault:"local" env-required:"true"`
	AppName string `envDefault:"scheduler"`
	HTTPServer
}

type HTTPServer struct {
	Address string `env:"SERVER_ADDRESS" envDefault:"localhost:8085"`
	//
	Timeout time.Duration `env:"SERVER_TIMEOUT" envDefault:"4s"`
	//
	IdleTimeout time.Duration `env:"SERVER_IDLE_TIMEOUT" envDefault:"60s"`
}

func GetConfig() *Config {
	var cfg Config

	//читаем конфиг файл и заполняем структуру
	if err := env.Parse(&cfg); err != nil {
		log.Printf("cannot read config: %s \n", err)
		os.Exit(1)
	}

	return &cfg
}
