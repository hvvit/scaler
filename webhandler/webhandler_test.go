package webhandler

import "testing"

func TestUrlFormatter(t *testing.T) {
	// test the urlFormatter function
	// it should return the url in the format http://host_addr:port/path
	url, _ := urlFormat("http://localhost:8080", "/app/status")
	if url != "http://localhost:8080/app/status" {
		t.Errorf("expected http://localhost:8080/app/status, got %v", url)
	}
}
