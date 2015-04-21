package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HandleImageNew(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Display New Image Form
	panic(fmt.Errorf("This should not be reached!"))
}
