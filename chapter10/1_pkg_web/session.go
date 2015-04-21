package main

import (
	"net/http"
	"net/url"
	"time"
)

const (
	// Keep users logged in for 3 days
	sessionLength     = 24 * 3 * time.Hour
	sessionCookieName = "GophrSession"
	sessionIDLength   = 20
)

type Session struct {
	ID     string
	UserID string
	Expiry time.Time
}

func (session *Session) Expired() bool {
	return session.Expiry.Before(time.Now())
}

func NewSession(w http.ResponseWriter) *Session {
	expiry := time.Now().Add(sessionLength)

	session := &Session{
		ID:     GenerateID("sess", sessionIDLength),
		Expiry: expiry,
	}

	cookie := http.Cookie{
		Name:    sessionCookieName,
		Value:   session.ID,
		Expires: session.Expiry,
	}

	http.SetCookie(w, &cookie)
	return session
}

func RequestSession(r *http.Request) *Session {
	cookie, err := r.Cookie(sessionCookieName)
	if err != nil {
		return nil
	}

	session, err := globalSessionStore.Find(cookie.Value)
	if err != nil {
		panic(err)
	}

	if session == nil {
		return nil
	}

	if session.Expired() {
		globalSessionStore.Delete(session)
		return nil
	}
	return session
}

func RequestUser(r *http.Request) *User {
	session := RequestSession(r)
	if session == nil || session.UserID == "" {
		return nil
	}

	user, err := globalUserStore.Find(session.UserID)
	if err != nil {
		panic(err)
	}
	return user
}

func RequireLogin(w http.ResponseWriter, r *http.Request) {
	// Let the request pass if we've got a user
	if RequestUser(r) != nil {
		return
	}

	query := url.Values{}
	query.Add("next", url.QueryEscape(r.URL.String()))

	http.Redirect(w, r, "/login?"+query.Encode(), http.StatusFound)
}

func FindOrCreateSession(w http.ResponseWriter, r *http.Request) *Session {
	session := RequestSession(r)
	if session == nil {
		session = NewSession(w)
	}

	return session
}
