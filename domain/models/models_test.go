package models

import (
	"encoding/json"
	"testing"
)

func TestUserJSONOmitsSensitiveFields(t *testing.T) {
	user := User{
		ID:       1,
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password-hash-marker",
		Token:    "token-marker",
	}

	payload, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("json.Marshal() returned an error: %v", err)
	}

	var decoded map[string]any
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("json.Unmarshal() returned an error: %v", err)
	}

	for _, field := range []string{"Password", "password", "Token", "token"} {
		if _, exists := decoded[field]; exists {
			t.Fatalf("serialized user contains sensitive field %q", field)
		}
	}
}

func TestCartJSONOmitsRawUserModel(t *testing.T) {
	payload, err := json.Marshal(Cart{
		ID:     1,
		UserID: 2,
		User: User{
			Password: "password-hash-marker",
			Token:    "token-marker",
		},
	})
	if err != nil {
		t.Fatalf("json.Marshal() returned an error: %v", err)
	}

	var decoded map[string]any
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("json.Unmarshal() returned an error: %v", err)
	}
	if _, exists := decoded["User"]; exists {
		t.Fatal("serialized cart contains its raw user model")
	}
}
