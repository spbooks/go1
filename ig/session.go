package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	sessionLength      = 24 * 3 * time.Hour
	sessionCookieName = "GophrSession"
	sessionIDLength = 20
)

type Session struct {
	ID     string
	UserID string
	Expiry time.Time
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

func (session Session) Expired() bool {
	return session.Expiry.Before(time.Now())
}

func (session *Session) Destroy() error {
	return globalSessionStore.Delete(session)
}

func RequireLogin(w http.ResponseWriter, r *http.Request) {
	// Let the request pass if we've got a user
	fmt.Println("Logged in?")
	if RequestUser(r) != nil {
		fmt.Println("Logged in")
		return
	}

	fmt.Println("Not Logged in")
	query := url.Values{}
	query.Add("next", url.QueryEscape(r.URL.String()))

	http.Redirect(w, r, "/login?"+query.Encode(), http.StatusFound)
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
		err = session.Destroy()
		if err != nil {
			panic(err)
		}
		return nil
	}
	return session
}

func MustRequestSession(w http.ResponseWriter, r *http.Request) *Session {
	session := RequestSession(r)
	if session == nil {
		session = NewSession(w)
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
