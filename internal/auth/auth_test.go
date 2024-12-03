package auth

import (
	"net/http"
	"testing"
)

func TestMissingAuthorizationHeaderShouldReturnAnError(t *testing.T) {
	headers := make(http.Header)

	key, err := GetAPIKey(headers)

	if err == nil {
		t.Fatalf("Expected error, got nil")
	}

	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("Expected error to be %q, got %q", ErrNoAuthHeaderIncluded.Error(), err.Error())
	}

	if key != "" {
		t.Fatalf("Expected key to be empty got %q", key)
	}
}

func TestMalformedApiKeyShouldReturnAnError(t *testing.T) {
	headers := make(http.Header)
	headers.Set("Authorization", "ApiKex asdfasdf")

	key, err := GetAPIKey(headers)

	if err == nil {
		t.Fatalf("Expected error, got nil")
	}

	if key != "" {
		t.Fatalf("Expected key to be empty got %q", key)
	}
}

func TestGetCorrectAuthKey(t *testing.T) {
	headers := make(http.Header)
	headers.Set("Authorization", "ApiKey asdfasdf")

	key, err := GetAPIKey(headers)

	if err != nil {
		t.Fatalf("Unexpected error, got %q", err.Error())
	}

	if key != "asdfasdf" {
		t.Fatalf("Expected key to be %q got %q", "asdfasdf", key)
	}
}
