package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/VincentDevi/JobsTrack/internal/handler"
	"github.com/VincentDevi/JobsTrack/internal/view"
	"github.com/a-h/templ"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(handler.NewHomeHandler().ServeHTTP))
	mux.Handle("/home", templ.Handler(view.Hello()))

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
