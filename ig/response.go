package main

import "net/http"

func Handle404(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func RenderError(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	RenderTemplate(w, r, "errors/500", map[string]interface{}{
		"Error": err.Error(),
	})
}
