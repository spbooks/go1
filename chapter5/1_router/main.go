package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := NewRouter()

	router.Handle("GET", "/", HandleHome)

	router.ServeFiles(
		"/assets/*filepath",
		http.Dir("assets/"),
	)

	middleware := Middleware{}
	middleware.Add(router)
	log.Fatal(http.ListenAndServe(":3000", middleware))
}

// Creates a new router
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = func(http.ResponseWriter, *http.Request) {}
	return router
}
