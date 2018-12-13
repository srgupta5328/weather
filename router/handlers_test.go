package router

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	r := NewRouter()
	r.HandleFunc("/GET", HealthCheck).Methods("GET")
	server := httptest.NewServer(r)
	defer server.Close()
	healthURL := fmt.Sprintf("%s/health", server.URL)
	req, err := http.NewRequest("GET", healthURL, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("HTTP status expected: 200, recieved: %d", res.StatusCode)
	}

}
