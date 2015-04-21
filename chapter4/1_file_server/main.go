package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))

	log.Fatal(http.ListenAndServe(":3000", mux))
}
