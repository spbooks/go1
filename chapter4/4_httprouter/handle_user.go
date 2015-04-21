package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HandleUserNew(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Display Home Page
	RenderTemplate(w, r, "users/new", nil)
}
