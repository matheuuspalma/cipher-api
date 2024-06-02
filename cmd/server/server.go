package server

import (
	"cipher/application"
	"cipher/cmd/server/adapters"
	"fmt"
	"html"
	"log"
	"net/http"
)

/*Function that will start the server http*/
func Server() {

	application.Init()

	http.HandleFunc("/cipher", adapters.Encode)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	fmt.Println("Listen on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
