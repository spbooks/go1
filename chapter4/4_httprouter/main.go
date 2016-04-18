package main

import (
	"fmt"
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

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", middleware)
}

type NotFound struct{}

func (n *NotFound) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

// Creates a new router
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	notFound := new(NotFound)
	router.NotFound = notFound
	return router
}
