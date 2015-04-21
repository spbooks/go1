package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	unauthenticatedRouter := NewRouter()
	unauthenticatedRouter.GET("/", HandleHome)
	unauthenticatedRouter.GET("/register", HandleUserNew)

	authenticatedRouter := NewRouter()
	authenticatedRouter.GET("/images/new", HandleImageNew)

	middleware := Middleware{}
	middleware.Add(unauthenticatedRouter)
	middleware.Add(http.HandlerFunc(AuthenticateRequest))
	middleware.Add(authenticatedRouter)

	http.ListenAndServe(":3000", middleware)
}

// Creates a new router
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = func(http.ResponseWriter, *http.Request) {}
	return router
}
