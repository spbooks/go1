package main

import (
	"fmt"
	"log"
	"net/http"

	"middleware"

	"github.com/julienschmidt/httprouter"
)

func init() {
	// Assign a user store
	store, err := NewFileUserStore("./data/users.json")
	if err != nil {
		panic(fmt.Errorf("Error creating user store: %s", err))
	}
	globalUserStore = store

	// Assign a session store
	sessionStore, err := NewFileSessionStore("./data/sessions.json")
	if err != nil {
		panic(fmt.Errorf("Error creating session store: %s", err))
	}
	globalSessionStore = sessionStore

	// Assign a sql database
	db, err := NewMySQLDB("root:root@tcp(127.0.0.1:3306)/gophr")
	if err != nil {
		panic(err)
	}
	globalMySQLDB = db

	// Assign an image store
	globalImageStore = NewDBImageStore()
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", NewApp()))
}

func NewApp() middleware.Middleware {
	router := NewRouter()

	router.Handle("GET", "/", HandleHome)
	router.Handle("GET", "/register", HandleUserNew)
	router.Handle("POST", "/register", HandleUserCreate)
	router.Handle("GET", "/login", HandleSessionNew)
	router.Handle("POST", "/login", HandleSessionCreate)
	router.Handle("GET", "/image/:imageID", HandleImageShow)
	router.Handle("GET", "/user/:userID", HandleUserShow)

	router.ServeFiles(
		"/assets/*filepath",
		http.Dir("assets/"),
	)
	router.ServeFiles(
		"/im/*filepath",
		http.Dir("data/images/"),
	)

	secureRouter := NewRouter()
	secureRouter.Handle("GET", "/sign-out", HandleSessionDestroy)
	secureRouter.Handle("GET", "/account", HandleUserEdit)
	secureRouter.Handle("POST", "/account", HandleUserUpdate)
	secureRouter.Handle("GET", "/images/new", HandleImageNew)
	secureRouter.Handle("POST", "/images/new", HandleImageCreate)

	middleware := middleware.Middleware{}
	middleware.Add(router)
	middleware.Add(http.HandlerFunc(RequireLogin))
	middleware.Add(secureRouter)

	return middleware
}

// Creates a new router
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = func(http.ResponseWriter, *http.Request) {}
	return router
}
