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
	if err := godotenv.Load(); err != nil {
		log.Fatal("couldn't load env variables from .env file")
	}
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
