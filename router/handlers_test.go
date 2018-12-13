package router

import (
	"net/http"
	"testing"
)

func TestHealth(t *testing.T) {
	r := NewRouter()

	req, err := http.NewRequest("GET", , nil)
	w := httptest.NewRecorder()
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("HTTP status expected: 200, recieved: %d", res.StatusCode)
	}

}
