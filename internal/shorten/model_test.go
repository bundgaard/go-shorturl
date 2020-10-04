package shorten

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP(t *testing.T) {

	api, err := NewShorten()
	if err != nil {
		t.Error(err)
	}
	server := httptest.NewServer(api)
	defer server.Close()

	var (
		location string
	)
	t.Run("POST request", func(t *testing.T) {
		var buf bytes.Buffer
		req1 := ShortURLModel{Location: "https://www.google.com"}
		if err = json.NewEncoder(&buf).Encode(req1); err != nil {
			t.Error("failed to encode JSON")
		}
		resp, err := http.Post(server.URL, "application/json", &buf)
		if err != nil {
			t.Error(err)
		}
		locationURL, err := resp.Location()
		if err != nil {
			t.Error(err)
		}
		location = locationURL.String()
		log.Println("Received ", location)
	})

	t.Run("GET request", func(t *testing.T) {

		resp, err := http.Get(location)
		if err != nil {
			t.Error(err)
		}
		if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
			t.Fail()
		}

	})

}
