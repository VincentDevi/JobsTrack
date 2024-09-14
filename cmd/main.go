package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/VincentDevi/JobsTrack/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(handler.NewHomeHandler().ServeHTTP))
	mux.Handle("/open", http.HandlerFunc(handler.NewHomeHandler().Open))
	mux.Handle("/close", http.HandlerFunc(handler.NewHomeHandler().Close))

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	fmt.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
