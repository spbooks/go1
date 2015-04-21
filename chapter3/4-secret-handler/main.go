package main

import (
	"fmt"
	"net/http"
	"time"
)

// UptimeHandler writes the number of seconds since starting to the response.
type UptimeHandler struct {
	Started time.Time
}

func NewUptimeHandler() UptimeHandler {
	return UptimeHandler{Started: time.Now()}
}

func (h UptimeHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(
		w,
		fmt.Sprintf("Current Uptime: %s", time.Since(h.Started)),
	)
}

// SecretTokenHandler secures a request with a secret token.
type SecretTokenHandler struct {
	next   http.Handler
	secret string
}

// ServeHTTP makes SecretTokenHandler implement the http.Handler interface.
func (h SecretTokenHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Check the query string for the secret token
	if req.URL.Query().Get("secret_token") == h.secret {
		// The secret token matched, call the next handler
		h.next.ServeHTTP(w, req)
	} else {
		// No match, return a 404 Not Found response
		http.NotFound(w, req)
	}
}

func main() {
	http.Handle("/", SecretTokenHandler{
		next:   NewUptimeHandler(),
		secret: "MySecret",
	})

	http.ListenAndServe(":3000", nil)
}
