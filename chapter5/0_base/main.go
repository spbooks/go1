package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
}

// Creates a new router
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = func(http.ResponseWriter, *http.Request) {}
	return router
}
