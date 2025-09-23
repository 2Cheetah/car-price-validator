package main

import (
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/2Cheetah/car-price-validator/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	MustLoadEnv()

	SetupLogger()

	s := server.NewServer()
	s.RegisterHandlers()
	s.MustRun()
}

func MustLoadEnv() {
	requiredEnvs := []string{"BASE_URL", "URL_PATH", "ORIGIN", "REFERER"}

	if hasAllRequiredEnvs(requiredEnvs) {
		return
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("env variables not found and couldn't be loaded from .env file: %v", err)
	}

	if !hasAllRequiredEnvs(requiredEnvs) {
		log.Fatal("some of the required envs are still missing even after loading .env file")
	}
}

func hasAllRequiredEnvs(requiredEnvs []string) bool {
	for _, env := range requiredEnvs {
		if _, exists := os.LookupEnv(env); !exists {
			return false
		}
	}
	return true
}

func SetupLogger() {
	logLevel := GetLogLevelFromEnv()
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))

	slog.SetDefault(logger)
	slog.Info("logger initialized", "level", logLevel)
}

func GetLogLevelFromEnv() slog.Level {
	logLevelEnv, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		return slog.LevelInfo
	}
	logLevelEnv = strings.TrimSpace(logLevelEnv)
	logLevelEnv = strings.ToUpper(logLevelEnv)

	switch logLevelEnv {
	case "DEBUG":
		return slog.LevelDebug
	case "WARN", "WARNING":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
