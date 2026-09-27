package config

import (
	"fmt"
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
)

func Load[T any](path string, logger *slog.Logger) (*T, error) {
	// env.ReadConfig reads environment variables into the Config struct
	// env.With
	var cfg T
	var err error
	if path == "" {
		err = cleanenv.ReadEnv(&cfg)
		logger.Info("config loaded from environment")
	} else {
		err = cleanenv.ReadConfig(path, &cfg)
		logger.Info("config loaded from file", "path", path)
	}
	if err != nil {
		return nil, fmt.Errorf("load config error: %w", err)
	}
	return &cfg, nil
}
