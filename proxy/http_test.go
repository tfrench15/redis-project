package proxy

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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
		key, err := ParseKey(rec, req)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if expected := test.key; expected != key {
			t.Errorf("Error: expected key %v, got %v", expected, key)
		}
	}
}

func TestServeHTTP(t *testing.T) {
	cache := NewCache(5, 10*time.Second)
	client, err := NewRedisClient("localhost:6379")
	if err != nil {
		t.Errorf("Error connecting to Redis: %v", err)
	}
	rp := NewRedisProxy(cache, client)
	srv := httptest.NewServer(rp)

	tests := []struct {
		path       string
		header     string
		statuscode int
	}{
		{"/sf", "MISS", http.StatusOK},
		{"/sf", "HIT", http.StatusOK},
		{"/ny", "MISS", http.StatusOK},
		{"/ny", "HIT", http.StatusOK},
		{"/hello", "MISS", http.StatusNotFound},
	}
	for _, test := range tests {
		req := httptest.NewRequest("GET", srv.URL+test.path, nil)
		rec := httptest.NewRecorder()
		rp.ServeHTTP(rec, req)
		res := rec.Result()
		head := res.Header.Get("X-Cache")
		if head != test.header {
			fmt.Println(test.path)
			t.Errorf("Error: expected header %v, got header %v", test.header, head)
		}
		if res.StatusCode != test.statuscode {
			fmt.Println(test.path)
			t.Errorf("Error: expected status %v, got %v", test.statuscode, res.StatusCode)
		}
		res.Body.Close()
	}
}
