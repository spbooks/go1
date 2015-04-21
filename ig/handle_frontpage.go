package main

import "net/http"

func HandleFrontpage(w http.ResponseWriter, r *http.Request) {
	images, err := globalImageStore.FindAll(0)
	if err != nil {
		panic(err)
	}
	RenderTemplate(w, r, "index/home", map[string]interface{}{
		"Images": images,
	})
}
