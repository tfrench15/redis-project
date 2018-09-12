package proxy

import (
	"errors"
	"io"
	"net/http"
	"path"
)

// ParseKey parses the key to lookup from the request URL,
// which is the base of the URL's path.  Returns the key if
// found, and the empty string otherwise.
func ParseKey(w http.ResponseWriter, r *http.Request) (string, error) {
	switch r.Method {
	case "GET":
		key := path.Base(r.URL.Path)
		return key, nil
	default:
		http.Error(w, "Error: please issue a GET request", http.StatusBadRequest)
		return "", errors.New("Bad Request")
	}
}

// ServeHTTP satisfies the http.Handler interface and implements the core
// Redis proxy service.
func (rp *RedisProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key, err := ParseKey(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if key == "" {
		http.Error(w, "Error: no key provided in request", http.StatusBadRequest)
	}
	value, ok := rp.cache.Lookup(key)
	if ok {
		w.Header().Set("X-Cache", "HIT")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, value)
		return
	}
	value, ok = rp.FetchFromRedis(key)
	if ok {
		w.Header().Set("X-Cache", "MISS")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, value)
		return
	}
	w.Header().Set("X-Cache", "MISS")
	http.Error(w, "Error: key not found", http.StatusNotFound)
}
