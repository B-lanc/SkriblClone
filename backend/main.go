package main

import "fmt"

// import "github.com/gorilla/websocket"
import "net/http"
import "html"
import "log"

func main() {
	fmt.Println("I'll do it later :p")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
