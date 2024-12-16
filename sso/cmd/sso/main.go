package main

import (
	"os"
	"sso/sso/internal/config"
	"sso/sso/internal/logger"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	logger := setupLogger(envLocal)
	logger.PrintInfo("Server started: 	"+cfg.Env, nil)

}

func setupLogger(env string) *logger.Logger {
	switch env {
	case envLocal:
		return logger.New(os.Stdout, logger.LevelInfo)
	}
	return nil
}
