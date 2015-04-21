package main

import (
	"log"
	"net/http"
)

func main() {
	setupImageStore()
	router := NewRouter()
	//mux.HandlerFunc("GET", "/login", HandleSessionNew)
	//mux.HandlerFunc("POST", "/login", HandleSessionCreate)

	router.HandlerFunc("GET", "/", HandleFrontpage)
	router.HandlerFunc("GET", "/register", HandleUserNew)
	router.HandlerFunc("POST", "/register", HandleUserCreate)
	router.HandlerFunc("GET", "/login", HandleSessionNew)
	router.HandlerFunc("POST", "/login", HandleSessionCreate)
	router.Handle("GET", "/image/:imageID", HandleImageShow)

	router.ServeFiles(
		"/assets/*filepath",
		http.Dir("assets/"),
	)
	router.ServeFiles(
		"/im/*filepath",
		http.Dir("data/images/"),
	)

	mux2 := NewRouter()
	mux2.HandlerFunc("GET", "/account", HandleUserEdit)
	mux2.HandlerFunc("POST", "/account", HandleUserUpdate)
	mux2.HandlerFunc("GET", "/sign-out", HandleSessionDestroy)

	mux2.Handle("GET", "/images/new", HandleImageNew)
	mux2.Handle("POST", "/images/new", HandleImageCreate)
	mux2.Handle("GET", "/image/:imageID/edit", HandleImageEdit)
	mux2.Handle("POST", "/image/:imageID/edit", HandleImageUpdate)

	mux2.Handle("GET", "/user/:userID", HandleUserShow)

	mw := Middleware{}
	mw.Add(router)
	mw.Add(http.HandlerFunc(RequireLogin))
	mw.Add(mux2)

	log.Fatal(http.ListenAndServe(":3000", mw))
}
