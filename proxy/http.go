package proxy

import (
	"net/http"
	"path"
)

// ParseKey parses the key to lookup from the request URL,
// which is the base of the URL's path.  Returns the key if
// found, and the empty string otherwise.
func ParseKey(w http.ResponseWriter, r *http.Request) string {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		key := path.Base(r.URL.Path)
		return key
	default:
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Error: please issue a GET request", http.StatusBadRequest)
		return ""
	}
}

// ServeHTTP satisfies ensures RedisProxy satisfies the http.Handler
// interface.
/*
func (rp *RedisProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := ParseKey(w, r)

}
*/
