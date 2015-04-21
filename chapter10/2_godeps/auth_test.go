package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRequestNewImageUnauthenticated(t *testing.T) {
	request, _ := http.NewRequest("GET", "/images/new", nil)
	recorder := httptest.NewRecorder()

	app := NewApp()
	app.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusFound {
		t.Error("Expected a redirect code, but got", recorder.Code)
	}

	loc := recorder.HeaderMap.Get("Location")
	if loc != "/login?next=%252Fimages%252Fnew" {
		t.Error("Expected Location to redirect to sign in, but got", loc)
	}
}

func BenchmarkRequestNewImageUnauthenticated(b *testing.B) {
	request, _ := http.NewRequest("GET", "/images/new", nil)

	recorder := httptest.NewRecorder()
	app := NewApp()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(recorder, request)
	}
}

func TestRequestNewImageUnauthenticatedPerformance(t *testing.T) {
	if testing.Short() {
		return
	}

	result := testing.Benchmark(func(b *testing.B) {
		request, _ := http.NewRequest("GET", "/images/new", nil)

		recorder := httptest.NewRecorder()
		app := NewApp()

		for i := 0; i < b.N; i++ {
			app.ServeHTTP(recorder, request)
		}
	})

	speed := result.NsPerOp()
	if speed > 4000 {
		t.Error("Expected response to take less than 4000 nanoseconds, but took", speed)
	}
}

type MockSessionStore struct {
	Session *Session
}

func (store MockSessionStore) Find(string) (*Session, error) {
	return store.Session, nil
}

func (store MockSessionStore) Save(*Session) error {
	return nil
}

func (store MockSessionStore) Delete(*Session) error {
	return nil
}

func TestRequestNewImageAuthenticated(t *testing.T) {
	// Replace the user store temporarily
	oldUserStore := globalUserStore
	defer func() {
		globalUserStore = oldUserStore
	}()
	globalUserStore = &MockUserStore{
		findUser: &User{},
	}

	expiry := time.Now().Add(time.Hour)

	// Replace the session store temporarily
	oldSessionStore := globalSessionStore
	defer func() {
		globalSessionStore = oldSessionStore
	}()
	globalSessionStore = &MockSessionStore{
		Session: &Session{
			ID:     "session_123",
			UserID: "user_123",
			Expiry: expiry,
		},
	}

	// Create a cookie for the
	authCookie := &http.Cookie{
		Name:    sessionCookieName,
		Value:   "session_123",
		Expires: expiry,
	}

	request, _ := http.NewRequest("GET", "/images/new", nil)
	request.AddCookie(authCookie)

	recorder := httptest.NewRecorder()

	app := NewApp()
	app.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Error("Expected a redirect code, but got", recorder.Code)
	}
}
