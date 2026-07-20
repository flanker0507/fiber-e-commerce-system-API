package config

import (
	"strings"
	"testing"
)

func TestLoadUsesDefaultsAndReadsRequiredVariables(t *testing.T) {
	setRequiredEnvironment(t)
	t.Setenv("APP_PORT", "")
	t.Setenv("APP_ENV", "")
	t.Setenv("MIDTRANS_ENV", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() returned an error: %v", err)
	}

	if cfg.AppPort != "8080" {
		t.Fatalf("AppPort = %q; want %q", cfg.AppPort, "8080")
	}
	if cfg.AppEnv != "development" {
		t.Fatalf("AppEnv = %q; want %q", cfg.AppEnv, "development")
	}
	if cfg.MidtransEnv != "sandbox" {
		t.Fatalf("MidtransEnv = %q; want %q", cfg.MidtransEnv, "sandbox")
	}
}

func TestLoadReportsMissingVariableNameWithoutValues(t *testing.T) {
	setRequiredEnvironment(t)
	t.Setenv("JWT_SECRET", "")

	_, err := Load()
	if err == nil {
		t.Fatal("Load() returned nil error; want a missing-variable error")
	}
	if !strings.Contains(err.Error(), "JWT_SECRET") {
		t.Fatalf("Load() error = %q; want it to identify JWT_SECRET", err)
	}
}

func TestLoadRejectsInvalidMidtransEnvironment(t *testing.T) {
	setRequiredEnvironment(t)
	t.Setenv("MIDTRANS_ENV", "invalid")

	_, err := Load()
	if err == nil {
		t.Fatal("Load() returned nil error; want an invalid-environment error")
	}
	if !strings.Contains(err.Error(), "MIDTRANS_ENV") {
		t.Fatalf("Load() error = %q; want it to identify MIDTRANS_ENV", err)
	}
}

func setRequiredEnvironment(t *testing.T) {
	t.Helper()

	values := map[string]string{
		"DB_HOST":             "127.0.0.1",
		"DB_PORT":             "3306",
		"DB_USER":             "test-user",
		"DB_PASSWORD":         "test-password",
		"DB_NAME":             "test-database",
		"JWT_SECRET":          "test-jwt-secret",
		"MIDTRANS_SERVER_KEY": "test-server-key",
		"MIDTRANS_CLIENT_KEY": "test-client-key",
	}
	for name, value := range values {
		t.Setenv(name, value)
	}
}
