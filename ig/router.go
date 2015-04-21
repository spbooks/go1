package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = func(http.ResponseWriter, *http.Request) {}
	return router
}
