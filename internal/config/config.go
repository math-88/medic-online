package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App    `yaml:"app"`
		MySQL  `yaml:"mysql"`
		Openai `yaml:"openai"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
		IsDebug *bool  `env-required:"true" yaml:"is_debug" env:"IS_DEBUG"`
	}

	// MySQL -.
	MySQL struct {
		Host     string `env-required:"true" yaml:"host" env:"HOST"`
		Username string `env-required:"true" yaml:"username" env:"USERNAME"`
		Password string `env-required:"true" yaml:"password" env:"PASSWORD"`
		Database string `env-required:"true" yaml:"database" env:"DATABASE"`
	}

	// HTTP -.
	Openai struct {
		Token string `env-required:"true" yaml:"token" env:"OPENAI_TOKEN"`
	}
)

// New returns app config.
func New() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./internal/config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
