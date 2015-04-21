package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.New("t").ParseGlob("templates/**/*.html"))

var errorTemplate = `
<html>
	<body>
		<h1>Error rendering template %s</h1>
		<p>%s</p>
	</body>
</html>
`

func RenderTemplate(w http.ResponseWriter, request *http.Request, name string, data interface{}) {
	err := templates.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf(errorTemplate, name, err),
			http.StatusInternalServerError,
		)
	}
}
