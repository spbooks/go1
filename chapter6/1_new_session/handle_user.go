package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HandleUserNew(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	RenderTemplate(w, r, "users/new", nil)
}

func HandleUserCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user, err := NewUser(
		r.FormValue("username"),
		r.FormValue("email"),
		r.FormValue("password"),
	)

	if err != nil {
		if IsValidationError(err) {
			RenderTemplate(w, r, "users/new", map[string]interface{}{
				"Error": err.Error(),
				"User":  user,
			})
			return
		}
		panic(err)
		return
	}

	err = globalUserStore.Save(user)
	if err != nil {
		panic(err)
		return
	}

	// Create a new session
	session := NewSession(w)
	session.UserID = user.ID

	http.Redirect(w, r, "/?flash=User+created", http.StatusFound)
}
