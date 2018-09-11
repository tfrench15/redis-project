package proxy

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseKey(t *testing.T) {
	tests := []struct {
		path string
		key  string
	}{
		{"/sanfrancisco", "sanfrancisco"},
		{"/newyork", "newyork"},
		{"/chicago", "chicago"},
		{"/miami", "miami"},
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, world!")
	}))
	defer srv.Close()

	for _, test := range tests {
		req := httptest.NewRequest("GET", srv.URL+test.path, nil)
		rec := httptest.NewRecorder()
		key := ParseKey(rec, req)
		if expected := test.key; expected != key {
			t.Errorf("Error: expected key %v, got %v", expected, key)
		}
	}
}
