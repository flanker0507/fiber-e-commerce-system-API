package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const (
	defaultAppPort     = "8080"
	defaultAppEnv      = "development"
	defaultMidtransEnv = "sandbox"
)

// Config contains all runtime configuration. Secret values are loaded once at
// startup and passed only to the components that require them.
type Config struct {
	AppPort string
	AppEnv  string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	JWTSecret string

	MidtransServerKey string
	MidtransClientKey string
	MidtransEnv       string
}

// Load reads a local .env file when present, then validates the environment.
// A .env file is optional so deployments can provide variables directly.
func Load() (Config, error) {
	if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return Config{}, fmt.Errorf("load .env: %w", err)
	}

	cfg := Config{
		AppPort:           valueOrDefault("APP_PORT", defaultAppPort),
		AppEnv:            valueOrDefault("APP_ENV", defaultAppEnv),
		DBHost:            strings.TrimSpace(os.Getenv("DB_HOST")),
		DBPort:            strings.TrimSpace(os.Getenv("DB_PORT")),
		DBUser:            strings.TrimSpace(os.Getenv("DB_USER")),
		DBPassword:        os.Getenv("DB_PASSWORD"),
		DBName:            strings.TrimSpace(os.Getenv("DB_NAME")),
		JWTSecret:         os.Getenv("JWT_SECRET"),
		MidtransServerKey: os.Getenv("MIDTRANS_SERVER_KEY"),
		MidtransClientKey: os.Getenv("MIDTRANS_CLIENT_KEY"),
		MidtransEnv:       strings.ToLower(valueOrDefault("MIDTRANS_ENV", defaultMidtransEnv)),
	}

	required := []struct {
		name  string
		value string
	}{
		{"DB_HOST", cfg.DBHost},
		{"DB_PORT", cfg.DBPort},
		{"DB_USER", cfg.DBUser},
		{"DB_PASSWORD", cfg.DBPassword},
		{"DB_NAME", cfg.DBName},
		{"JWT_SECRET", cfg.JWTSecret},
		{"MIDTRANS_SERVER_KEY", cfg.MidtransServerKey},
		{"MIDTRANS_CLIENT_KEY", cfg.MidtransClientKey},
	}

	for _, variable := range required {
		if strings.TrimSpace(variable.value) == "" {
			return Config{}, fmt.Errorf("missing required environment variable %s", variable.name)
		}
	}

	if cfg.MidtransEnv != "sandbox" && cfg.MidtransEnv != "production" {
		return Config{}, fmt.Errorf("invalid environment variable MIDTRANS_ENV: must be sandbox or production")
	}

	return cfg, nil
}

func valueOrDefault(name, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(name)); value != "" {
		return value
	}
	return fallback
}
