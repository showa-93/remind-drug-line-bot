package bot

import (
	"log"

	"github.com/caarlos0/env/v9"
)

type Config struct {
	Port     string `env:"PORT" envDefault:"8080"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"DEBUG"`

	LineSecret string `env:"LINE_SECRET"`
}

func NewConfig() Config {
	var c Config
	if err := env.Parse(&c); err != nil {
		log.Fatal("環境変数の読み込みに失敗")
	}

	return c
}
