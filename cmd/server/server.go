package server

import (
	"cipher/cmd/server/adapters"
	"fmt"
	"html"
	"log"
	"net/http"
)

/*Function that will start the server http*/
func Server() {

	http.HandleFunc("/cipher", adapters.Encode)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

/* mock testes */
