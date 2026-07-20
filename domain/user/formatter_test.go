package user

import (
	"encoding/json"
	"fiber-e-commerce-system-API/domain/models"
	"testing"
	"time"
)

func TestFormatUsersContainsOnlyPublicFields(t *testing.T) {
	createdAt := time.Date(2026, time.July, 20, 1, 2, 3, 0, time.UTC)
	updatedAt := createdAt.Add(time.Hour)
	responses := FormatUsers([]models.User{{
		ID:        7,
		Name:      "Safe User",
		Email:     "safe@example.com",
		Password:  "password-hash-marker",
		Role:      "admin",
		Token:     "token-marker",
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}})

	payload, err := json.Marshal(responses)
	if err != nil {
		t.Fatalf("json.Marshal() returned an error: %v", err)
	}

	var decoded []map[string]any
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("json.Unmarshal() returned an error: %v", err)
	}
	if len(decoded) != 1 {
		t.Fatalf("serialized response contains %d users; want 1", len(decoded))
	}

	wantFields := map[string]bool{
		"id":         true,
		"name":       true,
		"email":      true,
		"created_at": true,
		"updated_at": true,
	}
	for field := range decoded[0] {
		if !wantFields[field] {
			t.Fatalf("serialized user response contains unexpected field %q", field)
		}
	}
	for field := range wantFields {
		if _, exists := decoded[0][field]; !exists {
			t.Fatalf("serialized user response is missing field %q", field)
		}
	}
}

func TestFormatAuthResponseIncludesTokenWithoutPassword(t *testing.T) {
	response := FormatAuthResponse(models.User{
		ID:       7,
		Name:     "Safe User",
		Email:    "safe@example.com",
		Password: "password-hash-marker",
		Role:     "user",
	}, "issued-token")

	payload, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("json.Marshal() returned an error: %v", err)
	}

	var decoded map[string]any
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("json.Unmarshal() returned an error: %v", err)
	}
	if decoded["token"] != "issued-token" {
		t.Fatal("authentication response does not contain the issued token")
	}
	if _, exists := decoded["password"]; exists {
		t.Fatal("authentication response contains a password field")
	}
}
