package main

import (
	"net/http"
	"net/url"
)

func HandleSessionDestroy(w http.ResponseWriter, r *http.Request) {
	session := RequestSession(r)
	if session != nil {
		err := globalSessionStore.Delete(session)
		if err != nil {
			panic(err)
		}

	}

	RenderTemplate(w, r, "sessions/destroy", nil)
}

func HandleSessionNew(w http.ResponseWriter, r *http.Request) {
	next := r.URL.Query().Get("next")
	RenderTemplate(w, r, "sessions/new", map[string]interface{}{
		"Next": next,
	})
}

func HandleSessionCreate(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	next := r.FormValue("next")

	user, err := FindUser(username, password)
	if err != nil {
		RenderTemplate(w, r, "sessions/new", map[string]interface{}{
			"Error": err,
			"User":  user,
			"Next":  next,
		})
	}

	session := MustRequestSession(w, r)
	session.UserID = user.ID
	err = globalSessionStore.Save(session)
	if err != nil {
		panic(err)
	}

	if next == "" {
		next = "/"
	}

	nextUnescaped, err := url.QueryUnescape(next)

	http.Redirect(w, r, nextUnescaped+"?flash=Signed+in", http.StatusFound)
}
