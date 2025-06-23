package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "ApiKey NOTREAL")
	apiKey, err := GetAPIKey(header)
	if err != nil {
		t.Fail()
	}
	if apiKey != "NOTREAL" {
		t.Fail()
	}
}

func TestGetAPIKey_NoAuthHeader(t *testing.T) {
	header := http.Header{}
	_, err := GetAPIKey(header)
	if err == nil {
		t.Log("An error should have occurred")
		t.Fail()
	}
	if err != ErrNoAuthHeaderIncluded {
		t.Log("Incorrect error message")
		t.Fail()
	}
}

func TestGetAPIKey_MalformedAuthHeader(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "SHABLAM")
	_, err := GetAPIKey(header)
	if err == nil {
		t.Log("An error should have occurred")
		t.Fail()
	}
	if err.Error() != "malformed authorization header" {
		t.Log("Incorrect error message")
		t.Fail()
	}
}
