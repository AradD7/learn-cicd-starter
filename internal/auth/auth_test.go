package auth

import (
	"net/http"
	"testing"
)


func TestGetApiKey(t *testing.T) {
	req1, _ := http.NewRequest("GET", "http://example.com", nil)
	req2, _ := http.NewRequest("GET", "http://example.com", nil)
	req3, _ := http.NewRequest("GET", "http://example.com", nil)
	req4, _ := http.NewRequest("GET", "http://example.com", nil)

	req1.Header.Add("Authorization", "ApiKey 123456abc")
	req2.Header.Add("Authorization", "Token meow")
	req3.Header.Add("Authorization", "ApiKey")

	if apiKey, err := GetAPIKey(req1.Header); err != nil || apiKey != "123456abc" {
		t.Fatalf("want 123456abc, got: %s and error: %v", apiKey, err)
	}

	if _, err := GetAPIKey(req2.Header); err == nil {
		t.Fatalf("wanted an error, got nothing")
	}

	if _, err := GetAPIKey(req3.Header); err == nil {
		t.Fatalf("wanted an error, got nothing")
	}

	if _, err := GetAPIKey(req4.Header); err == nil {
		t.Fatalf("wanted an error, got nothing")
	}
}
