package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HandleImageNew(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	RenderTemplate(w, r, "images/new", nil)
}

func HandleImageCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.FormValue("url") != "" {
		HandleImageCreateFromURL(w, r)
		return
	}

	HandleImageCreateFromFile(w, r)
}

func HandleImageCreateFromURL(w http.ResponseWriter, r *http.Request) {
	user := RequestUser(r)

	image := NewImage(user)
	image.Description = r.FormValue("description")

	err := image.CreateFromURL(r.FormValue("url"))

	if err != nil {
		if IsValidationError(err) {
			RenderTemplate(w, r, "images/new", map[string]interface{}{
				"Error":    err,
				"ImageURL": r.FormValue("url"),
				"Image":    image,
			})
			return
		}
		panic(err)
	}

	http.Redirect(w, r, "/?flash=Image+Uploaded+Successfully", http.StatusFound)
}

func HandleImageCreateFromFile(w http.ResponseWriter, r *http.Request) {

	user := RequestUser(r)
	image := NewImage(user)
	image.Description = r.FormValue("description")

	file, headers, err := r.FormFile("file")

	// No file was uploaded.
	if file == nil {
		RenderTemplate(w, r, "images/new", map[string]interface{}{
			"Error": errNoImage,
			"Image": image,
		})
		return
	}

	// A file was uploaded, but an error occurredd
	if err != nil {
		panic(err)
	}

	defer file.Close()

	err = image.CreateFromFile(file, headers)
	if err != nil {
		RenderTemplate(w, r, "images/new", map[string]interface{}{
			"Error": err,
			"Image": image,
		})
		return
	}

	http.Redirect(w, r, "/?flash=Image+Uploaded+Successfully", http.StatusFound)
}
