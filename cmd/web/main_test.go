package main

import (
	"net/http"
	"testing"
	"time"
)

func TestStartServer(t *testing.T) {
	port := ":8080"

	go func() {
		startServer(port)
	}()

	time.Sleep(100 * time.Millisecond)

	resp, err := http.Get("http://localhost" + port + "/")
	if err != nil {
		t.Fatalf("Failed to send request to startServer: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotFound {
		t.Errorf("Unexpected status code from startServer: %d", resp.StatusCode)
	}
}
