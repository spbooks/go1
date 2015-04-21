package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", "Go Server")
		fmt.Fprintf(w, `
		<html>
			<body>
				Hello Gopher
			</body>
		</html>`)
	})

	http.ListenAndServe(":3000", nil)
}
